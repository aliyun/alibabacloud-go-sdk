// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGitRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetGitRepositoryResponseBody
	GetDescription() *string
	SetHtmlUrl(v string) *GetGitRepositoryResponseBody
	GetHtmlUrl() *string
	SetIsPrivate(v bool) *GetGitRepositoryResponseBody
	GetIsPrivate() *bool
	SetRequestId(v string) *GetGitRepositoryResponseBody
	GetRequestId() *string
}

type GetGitRepositoryResponseBody struct {
	// example:
	//
	// FASTJSON 2.0.x has been released, faster and more secure, recommend you upgrade.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://github.com/alibaba/fastjson
	HtmlUrl *string `json:"HtmlUrl,omitempty" xml:"HtmlUrl,omitempty"`
	// example:
	//
	// False
	IsPrivate *bool `json:"IsPrivate,omitempty" xml:"IsPrivate,omitempty"`
	// example:
	//
	// 13B71887-D215-53B5-8818-4D3190604B26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGitRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGitRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetGitRepositoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetGitRepositoryResponseBody) GetHtmlUrl() *string {
	return s.HtmlUrl
}

func (s *GetGitRepositoryResponseBody) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *GetGitRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGitRepositoryResponseBody) SetDescription(v string) *GetGitRepositoryResponseBody {
	s.Description = &v
	return s
}

func (s *GetGitRepositoryResponseBody) SetHtmlUrl(v string) *GetGitRepositoryResponseBody {
	s.HtmlUrl = &v
	return s
}

func (s *GetGitRepositoryResponseBody) SetIsPrivate(v bool) *GetGitRepositoryResponseBody {
	s.IsPrivate = &v
	return s
}

func (s *GetGitRepositoryResponseBody) SetRequestId(v string) *GetGitRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGitRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
