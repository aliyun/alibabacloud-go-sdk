// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordsToLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeywords(v string) *AddKeywordsToLibRequest
	GetKeywords() *string
	SetKeywordsObject(v string) *AddKeywordsToLibRequest
	GetKeywordsObject() *string
	SetLibId(v string) *AddKeywordsToLibRequest
	GetLibId() *string
	SetProperties(v string) *AddKeywordsToLibRequest
	GetProperties() *string
	SetRegionId(v string) *AddKeywordsToLibRequest
	GetRegionId() *string
	SetTenantCode(v string) *AddKeywordsToLibRequest
	GetTenantCode() *string
}

type AddKeywordsToLibRequest struct {
	// The keyword to be added.
	//
	// example:
	//
	// keyword
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The name of the keyword file.
	//
	// example:
	//
	// upload/1e5353c0-0d91-40ba-9d41-ae7abd3fe561.txt
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
	// The id of the keyword library.
	//
	// example:
	//
	// customxx_xxxx
	LibId      *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s AddKeywordsToLibRequest) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsToLibRequest) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *AddKeywordsToLibRequest) GetKeywordsObject() *string {
	return s.KeywordsObject
}

func (s *AddKeywordsToLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *AddKeywordsToLibRequest) GetProperties() *string {
	return s.Properties
}

func (s *AddKeywordsToLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddKeywordsToLibRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *AddKeywordsToLibRequest) SetKeywords(v string) *AddKeywordsToLibRequest {
	s.Keywords = &v
	return s
}

func (s *AddKeywordsToLibRequest) SetKeywordsObject(v string) *AddKeywordsToLibRequest {
	s.KeywordsObject = &v
	return s
}

func (s *AddKeywordsToLibRequest) SetLibId(v string) *AddKeywordsToLibRequest {
	s.LibId = &v
	return s
}

func (s *AddKeywordsToLibRequest) SetProperties(v string) *AddKeywordsToLibRequest {
	s.Properties = &v
	return s
}

func (s *AddKeywordsToLibRequest) SetRegionId(v string) *AddKeywordsToLibRequest {
	s.RegionId = &v
	return s
}

func (s *AddKeywordsToLibRequest) SetTenantCode(v string) *AddKeywordsToLibRequest {
	s.TenantCode = &v
	return s
}

func (s *AddKeywordsToLibRequest) Validate() error {
	return dara.Validate(s)
}
