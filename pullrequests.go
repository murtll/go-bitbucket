package bitbucket

import (
	"encoding/json"
	"log"
	"net/url"
)

type PullRequests struct {
	c *Client
}

func (p *PullRequests) Create(po *PullRequestsOptions) (interface{}, error) {
	data, err := p.buildPullRequestBody(po)
	if err != nil {
		return nil, err
	}
	urlStr := p.c.requestUrl("/projects/%s/repos/%s/pull-requests/", po.Project, po.RepoSlug)
  log.Printf("sending '%s' to '%s'", data, urlStr)
	return p.c.executeWithContext("POST", urlStr, data, po.ctx)
}

func (p *PullRequests) Update(po *PullRequestsOptions) (interface{}, error) {
	data, err := p.buildPullRequestBody(po)
	if err != nil {
		return nil, err
	}
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID
	return p.c.execute("PUT", urlStr, data)
}

func (p *PullRequests) Gets(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/"

	if po.States != nil && len(po.States) != 0 {
		parsed, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}
		query := parsed.Query()
		for _, state := range po.States {
			query.Set("state", state)
		}
		parsed.RawQuery = query.Encode()
		urlStr = parsed.String()
	}

	if po.Query != "" {
		parsed, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}
		query := parsed.Query()
		query.Set("q", po.Query)
		parsed.RawQuery = query.Encode()
		urlStr = parsed.String()
	}

	if po.Sort != "" {
		parsed, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}
		query := parsed.Query()
		query.Set("sort", po.Sort)
		parsed.RawQuery = query.Encode()
		urlStr = parsed.String()
	}

	return p.c.executePaginated("GET", urlStr, "", nil)
}

func (p *PullRequests) Get(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/projects/" + po.Project + "/" + po.RepoSlug + "/pull-requests/" + po.ID
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Activities(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/activity"
	return p.c.executePaginated("GET", urlStr, "", nil)
}

func (p *PullRequests) Activity(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/activity"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Commits(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/commits"
	return p.c.executePaginated("GET", urlStr, "", nil)
}

func (p *PullRequests) Patch(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/patch"
	return p.c.executeRaw("GET", urlStr, "")
}

func (p *PullRequests) Diff(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/diff"
	return p.c.executeRaw("GET", urlStr, "")
}

func (p *PullRequests) Merge(po *PullRequestsOptions) (interface{}, error) {
	data, err := p.buildPullRequestBody(po)
	if err != nil {
		return nil, err
	}
	urlStr := p.c.GetApiBaseURL() + "/projects/" + po.Project + "/repos/" + po.RepoSlug + "/pull-requests/" + po.ID + "/merge"
  log.Printf("sending '%s' to '%s'", data, urlStr)
	return p.c.executeWithContext("POST", urlStr, data, po.ctx)
}

func (p *PullRequests) Decline(po *PullRequestsOptions) (interface{}, error) {
	data, err := p.buildPullRequestBody(po)
	if err != nil {
		return nil, err
	}
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/decline"
	return p.c.executeWithContext("POST", urlStr, data, po.ctx)
}

func (p *PullRequests) Approve(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/approve"
	return p.c.executeWithContext("POST", urlStr, "", po.ctx)
}

func (p *PullRequests) UnApprove(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/approve"
	return p.c.execute("DELETE", urlStr, "")
}

func (p *PullRequests) RequestChanges(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/request-changes"
	return p.c.executeWithContext("POST", urlStr, "", po.ctx)
}

func (p *PullRequests) UnRequestChanges(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/request-changes"
	return p.c.execute("DELETE", urlStr, "")
}

func (p *PullRequests) AddComment(co *PullRequestCommentOptions) (interface{}, error) {
	data, err := p.buildPullRequestCommentBody(co)
	if err != nil {
		return nil, err
	}

	urlStr := p.c.requestUrl("/repositories/%s/%s/pullrequests/%s/comments", co.Owner, co.RepoSlug, co.PullRequestID)
	return p.c.executeWithContext("POST", urlStr, data, co.ctx)
}

func (p *PullRequests) UpdateComment(co *PullRequestCommentOptions) (interface{}, error) {
	data, err := p.buildPullRequestCommentBody(co)
	if err != nil {
		return nil, err
	}

	urlStr := p.c.requestUrl("/repositories/%s/%s/pullrequests/%s/comments/%s", co.Owner, co.RepoSlug, co.PullRequestID, co.CommentId)
	return p.c.execute("PUT", urlStr, data)
}

func (p *PullRequests) DeleteComment(co *PullRequestCommentOptions) (interface{}, error) {
	urlStr := p.c.requestUrl("/repositories/%s/%s/pullrequests/%s/comments/%s", co.Owner, co.RepoSlug, co.PullRequestID, co.CommentId)
	return p.c.execute("DELETE", urlStr, "")
}

func (p *PullRequests) GetComments(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/comments/"
	return p.c.executePaginated("GET", urlStr, "", nil)
}

func (p *PullRequests) GetComment(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/comments/" + po.CommentID
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Statuses(po *PullRequestsOptions) (interface{}, error) {
	urlStr := p.c.GetApiBaseURL() + "/repositories/" + po.Project + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/statuses"
	if po.Query != "" {
		parsed, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}
		query := parsed.Query()
		query.Set("q", po.Query)
		parsed.RawQuery = query.Encode()
		urlStr = parsed.String()
	}

	if po.Sort != "" {
		parsed, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}
		query := parsed.Query()
		query.Set("sort", po.Sort)
		parsed.RawQuery = query.Encode()
		urlStr = parsed.String()
	}
	return p.c.executePaginated("GET", urlStr, "", nil)
}

func (p *PullRequests) buildPullRequestBody(po *PullRequestsOptions) (string, error) {
	body := map[string]interface{}{}
	body["fromRef"] = map[string]interface{}{}
	body["toRef"] = map[string]interface{}{}
	body["reviewers"] = []map[string]string{}
	body["title"] = ""
	body["description"] = ""

	if n := len(po.Reviewers); n > 0 {
		body["reviewers"] = make([]map[string]string, n)
		for i, uuid := range po.Reviewers {
			body["reviewers"].([]map[string]string)[i] = map[string]string{"uuid": uuid}
		}
	}

	if po.SourceBranch != "" {
		body["fromRef"].(map[string]interface{})["displayId"] = po.SourceBranch
		body["fromRef"].(map[string]interface{})["id"] = "refs/heads/" + po.SourceBranch
	}

	if po.SourceRepository != "" {
		body["fromRef"].(map[string]interface{})["repository"] = map[string]interface{}{
      "name": po.SourceRepository,
      "project": map[string]interface{}{
        "key": po.Project,
      },
      "slug": po.SourceRepository,
    }
	}

	if po.DestinationBranch != "" {
		body["toRef"].(map[string]interface{})["displayId"] = po.DestinationBranch
		body["toRef"].(map[string]interface{})["id"] = "refs/heads/" + po.DestinationBranch
	}

	if po.DestinationRepository != "" {
		body["toRef"].(map[string]interface{})["repository"] = map[string]interface{}{
      "name": po.DestinationRepository,
      "project": map[string]interface{}{
        "key": po.Project,
      },
      "slug": po.DestinationRepository,
    }
	}

	if po.Title != "" {
		body["title"] = po.Title
	}

	if po.Description != "" {
		body["description"] = po.Description
	}

	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (p *PullRequests) buildPullRequestCommentBody(co *PullRequestCommentOptions) (string, error) {
	body := map[string]interface{}{}
	body["content"] = map[string]interface{}{
		"raw": co.Content,
	}

	if co.Parent != nil {
		body["parent"] = map[string]interface{}{
			"id": co.Parent,
		}
	}

	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
