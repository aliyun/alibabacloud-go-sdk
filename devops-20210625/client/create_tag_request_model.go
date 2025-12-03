// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateTagRequest
	GetAccessToken() *string
	SetMessage(v string) *CreateTagRequest
	GetMessage() *string
	SetRef(v string) *CreateTagRequest
	GetRef() *string
	SetTagName(v string) *CreateTagRequest
	GetTagName() *string
	SetOrganizationId(v string) *CreateTagRequest
	GetOrganizationId() *string
}

type CreateTagRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Message     *string `json:"message,omitempty" xml:"message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// master
	Ref *string `json:"ref,omitempty" xml:"ref,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1.0
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateTagRequest) GetMessage() *string {
	return s.Message
}

func (s *CreateTagRequest) GetRef() *string {
	return s.Ref
}

func (s *CreateTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *CreateTagRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateTagRequest) SetAccessToken(v string) *CreateTagRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateTagRequest) SetMessage(v string) *CreateTagRequest {
	s.Message = &v
	return s
}

func (s *CreateTagRequest) SetRef(v string) *CreateTagRequest {
	s.Ref = &v
	return s
}

func (s *CreateTagRequest) SetTagName(v string) *CreateTagRequest {
	s.TagName = &v
	return s
}

func (s *CreateTagRequest) SetOrganizationId(v string) *CreateTagRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateTagRequest) Validate() error {
	return dara.Validate(s)
}
