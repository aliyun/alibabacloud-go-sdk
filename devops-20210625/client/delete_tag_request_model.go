// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteTagRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *DeleteTagRequest
	GetOrganizationId() *string
	SetTagName(v string) *DeleteTagRequest
	GetTagName() *string
}

type DeleteTagRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 609633ffd40eb063bac8165a
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1.0
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteTagRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *DeleteTagRequest) SetAccessToken(v string) *DeleteTagRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteTagRequest) SetOrganizationId(v string) *DeleteTagRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteTagRequest) SetTagName(v string) *DeleteTagRequest {
	s.TagName = &v
	return s
}

func (s *DeleteTagRequest) Validate() error {
	return dara.Validate(s)
}
