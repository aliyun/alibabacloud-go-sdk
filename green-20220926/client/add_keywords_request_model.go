// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeywords(v string) *AddKeywordsRequest
	GetKeywords() *string
	SetKeywordsObject(v string) *AddKeywordsRequest
	GetKeywordsObject() *string
	SetLibId(v string) *AddKeywordsRequest
	GetLibId() *string
	SetRegionId(v string) *AddKeywordsRequest
	GetRegionId() *string
	SetTenantCode(v string) *AddKeywordsRequest
	GetTenantCode() *string
}

type AddKeywordsRequest struct {
	// The keywords. Separate multiple keywords with
	//
	// .
	//
	// example:
	//
	// Keyword1\\nKeyword2
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The keyword file name.
	//
	// example:
	//
	// upload/1e5353c0-0d91-40ba-9d41-ae7abd3fe561.txt
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
	// The keyword library ID.
	//
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The library code.
	//
	// - desensitize: masking library
	//
	// example:
	//
	// desensitize
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s AddKeywordsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsRequest) GoString() string {
	return s.String()
}

func (s *AddKeywordsRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *AddKeywordsRequest) GetKeywordsObject() *string {
	return s.KeywordsObject
}

func (s *AddKeywordsRequest) GetLibId() *string {
	return s.LibId
}

func (s *AddKeywordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddKeywordsRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *AddKeywordsRequest) SetKeywords(v string) *AddKeywordsRequest {
	s.Keywords = &v
	return s
}

func (s *AddKeywordsRequest) SetKeywordsObject(v string) *AddKeywordsRequest {
	s.KeywordsObject = &v
	return s
}

func (s *AddKeywordsRequest) SetLibId(v string) *AddKeywordsRequest {
	s.LibId = &v
	return s
}

func (s *AddKeywordsRequest) SetRegionId(v string) *AddKeywordsRequest {
	s.RegionId = &v
	return s
}

func (s *AddKeywordsRequest) SetTenantCode(v string) *AddKeywordsRequest {
	s.TenantCode = &v
	return s
}

func (s *AddKeywordsRequest) Validate() error {
	return dara.Validate(s)
}
