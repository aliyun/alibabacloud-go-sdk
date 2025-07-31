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
}

type AddKeywordsRequest struct {
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// upload/1e5353c0-0d91-40ba-9d41-ae7abd3fe561.txt
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *AddKeywordsRequest) Validate() error {
	return dara.Validate(s)
}
