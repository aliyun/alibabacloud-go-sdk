// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSearchCodePreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *GetSearchCodePreviewRequest
	GetDocId() *string
	SetIsDsl(v bool) *GetSearchCodePreviewRequest
	GetIsDsl() *bool
	SetKeyword(v string) *GetSearchCodePreviewRequest
	GetKeyword() *string
	SetOrganizationId(v string) *GetSearchCodePreviewRequest
	GetOrganizationId() *string
}

type GetSearchCodePreviewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 60d54f3daccf2bbd6659f3ad/gitlabhq/master/config/environments/test.rb
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// false
	IsDsl *bool `json:"isDsl,omitempty" xml:"isDsl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 627475075b46541dd2ff01bc
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetSearchCodePreviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSearchCodePreviewRequest) GoString() string {
	return s.String()
}

func (s *GetSearchCodePreviewRequest) GetDocId() *string {
	return s.DocId
}

func (s *GetSearchCodePreviewRequest) GetIsDsl() *bool {
	return s.IsDsl
}

func (s *GetSearchCodePreviewRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetSearchCodePreviewRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetSearchCodePreviewRequest) SetDocId(v string) *GetSearchCodePreviewRequest {
	s.DocId = &v
	return s
}

func (s *GetSearchCodePreviewRequest) SetIsDsl(v bool) *GetSearchCodePreviewRequest {
	s.IsDsl = &v
	return s
}

func (s *GetSearchCodePreviewRequest) SetKeyword(v string) *GetSearchCodePreviewRequest {
	s.Keyword = &v
	return s
}

func (s *GetSearchCodePreviewRequest) SetOrganizationId(v string) *GetSearchCodePreviewRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetSearchCodePreviewRequest) Validate() error {
	return dara.Validate(s)
}
