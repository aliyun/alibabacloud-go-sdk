// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVendorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListVendorsShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListVendorsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVendorsShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVendorsShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListVendorsShrinkRequest
	GetRoleFor() *int64
	SetVendorIdsShrink(v string) *ListVendorsShrinkRequest
	GetVendorIdsShrink() *string
	SetVendorName(v string) *ListVendorsShrinkRequest
	GetVendorName() *string
	SetVendorType(v string) *ListVendorsShrinkRequest
	GetVendorType() *string
}

type ListVendorsShrinkRequest struct {
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return for this request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. If this is your first query or if no next page exists, you do not need to specify this parameter. If a next page exists, set the value to the NextToken value that is returned in the last response.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the Data Management center for threat analysis is located. Select a region for the Management Hub based on the region where your asset resides. Valid values:
	//
	// - cn-hangzhou: Your asset is in the Chinese mainland.
	//
	// - ap-southeast-1: Your asset is outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can use this ID to switch to the member\\"s perspective.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// A list of vendors.
	VendorIdsShrink *string `json:"VendorIds,omitempty" xml:"VendorIds,omitempty"`
	// The vendor name.
	//
	// example:
	//
	// 111
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
	// The vendor type. Valid values:
	//
	// - preset
	//
	// - custom
	//
	// example:
	//
	// preset
	VendorType *string `json:"VendorType,omitempty" xml:"VendorType,omitempty"`
}

func (s ListVendorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVendorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListVendorsShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListVendorsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVendorsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVendorsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVendorsShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListVendorsShrinkRequest) GetVendorIdsShrink() *string {
	return s.VendorIdsShrink
}

func (s *ListVendorsShrinkRequest) GetVendorName() *string {
	return s.VendorName
}

func (s *ListVendorsShrinkRequest) GetVendorType() *string {
	return s.VendorType
}

func (s *ListVendorsShrinkRequest) SetLang(v string) *ListVendorsShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListVendorsShrinkRequest) SetMaxResults(v int32) *ListVendorsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVendorsShrinkRequest) SetNextToken(v string) *ListVendorsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListVendorsShrinkRequest) SetRegionId(v string) *ListVendorsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListVendorsShrinkRequest) SetRoleFor(v int64) *ListVendorsShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListVendorsShrinkRequest) SetVendorIdsShrink(v string) *ListVendorsShrinkRequest {
	s.VendorIdsShrink = &v
	return s
}

func (s *ListVendorsShrinkRequest) SetVendorName(v string) *ListVendorsShrinkRequest {
	s.VendorName = &v
	return s
}

func (s *ListVendorsShrinkRequest) SetVendorType(v string) *ListVendorsShrinkRequest {
	s.VendorType = &v
	return s
}

func (s *ListVendorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
