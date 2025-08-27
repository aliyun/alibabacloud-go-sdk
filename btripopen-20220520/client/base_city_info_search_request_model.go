// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseCityInfoSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *BaseCityInfoSearchRequest
	GetKeyword() *string
	SetRegion(v string) *BaseCityInfoSearchRequest
	GetRegion() *string
}

type BaseCityInfoSearchRequest struct {
	// This parameter is required.
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 0
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s BaseCityInfoSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s BaseCityInfoSearchRequest) GoString() string {
	return s.String()
}

func (s *BaseCityInfoSearchRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *BaseCityInfoSearchRequest) GetRegion() *string {
	return s.Region
}

func (s *BaseCityInfoSearchRequest) SetKeyword(v string) *BaseCityInfoSearchRequest {
	s.Keyword = &v
	return s
}

func (s *BaseCityInfoSearchRequest) SetRegion(v string) *BaseCityInfoSearchRequest {
	s.Region = &v
	return s
}

func (s *BaseCityInfoSearchRequest) Validate() error {
	return dara.Validate(s)
}
