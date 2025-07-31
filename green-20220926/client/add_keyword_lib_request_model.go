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
	SetRegionId(v string) *AddKeywordLibRequest
	GetRegionId() *string
}

type AddKeywordLibRequest struct {
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// upload/1e5353c0-0d91-40ba-9d41-ae7abd3fe561.txt
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
	LibName        *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *AddKeywordLibRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *AddKeywordLibRequest) SetRegionId(v string) *AddKeywordLibRequest {
	s.RegionId = &v
	return s
}

func (s *AddKeywordLibRequest) Validate() error {
	return dara.Validate(s)
}
