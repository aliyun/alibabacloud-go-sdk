// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeywords(v string) *AddKeywordLibRequest
	GetKeywords() *string
	SetKeywordsObject(v string) *AddKeywordLibRequest
	GetKeywordsObject() *string
	SetLibName(v string) *AddKeywordLibRequest
	GetLibName() *string
	SetProperties(v string) *AddKeywordLibRequest
	GetProperties() *string
	SetRegionId(v string) *AddKeywordLibRequest
	GetRegionId() *string
	SetTenantCode(v string) *AddKeywordLibRequest
	GetTenantCode() *string
}

type AddKeywordLibRequest struct {
	// The keywords. Separate multiple keywords with
	//
	// .
	//
	// example:
	//
	// keyword1\\nkeyword2
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The name of the keyword file.
	//
	// example:
	//
	// upload/1e5353c0-0d91-40ba-9d41-ae7abd3fe561.txt
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
	// The name of the keyword library.
	//
	// example:
	//
	// TestLibrary.
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// The properties.
	//
	// example:
	//
	// {"attribute":"xx"}
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The library code.
	//
	// - desensitize: desensitization library
	//
	// example:
	//
	// desensitize
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s AddKeywordLibRequest) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *AddKeywordLibRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *AddKeywordLibRequest) GetKeywordsObject() *string {
	return s.KeywordsObject
}

func (s *AddKeywordLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *AddKeywordLibRequest) GetProperties() *string {
	return s.Properties
}

func (s *AddKeywordLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddKeywordLibRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *AddKeywordLibRequest) SetKeywords(v string) *AddKeywordLibRequest {
	s.Keywords = &v
	return s
}

func (s *AddKeywordLibRequest) SetKeywordsObject(v string) *AddKeywordLibRequest {
	s.KeywordsObject = &v
	return s
}

func (s *AddKeywordLibRequest) SetLibName(v string) *AddKeywordLibRequest {
	s.LibName = &v
	return s
}

func (s *AddKeywordLibRequest) SetProperties(v string) *AddKeywordLibRequest {
	s.Properties = &v
	return s
}

func (s *AddKeywordLibRequest) SetRegionId(v string) *AddKeywordLibRequest {
	s.RegionId = &v
	return s
}

func (s *AddKeywordLibRequest) SetTenantCode(v string) *AddKeywordLibRequest {
	s.TenantCode = &v
	return s
}

func (s *AddKeywordLibRequest) Validate() error {
	return dara.Validate(s)
}
