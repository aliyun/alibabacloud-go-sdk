// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVendorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListVendorsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListVendorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVendorsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVendorsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListVendorsRequest
	GetRoleFor() *int64
	SetVendorIds(v []*string) *ListVendorsRequest
	GetVendorIds() []*string
	SetVendorName(v string) *ListVendorsRequest
	GetVendorName() *string
	SetVendorType(v string) *ListVendorsRequest
	GetVendorType() *string
}

type ListVendorsRequest struct {
	// example:
	//
	// en。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor   *int64    `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	VendorIds []*string `json:"VendorIds,omitempty" xml:"VendorIds,omitempty" type:"Repeated"`
	// example:
	//
	// 111。
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
	// example:
	//
	// preset。
	VendorType *string `json:"VendorType,omitempty" xml:"VendorType,omitempty"`
}

func (s ListVendorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVendorsRequest) GoString() string {
	return s.String()
}

func (s *ListVendorsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListVendorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVendorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVendorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVendorsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListVendorsRequest) GetVendorIds() []*string {
	return s.VendorIds
}

func (s *ListVendorsRequest) GetVendorName() *string {
	return s.VendorName
}

func (s *ListVendorsRequest) GetVendorType() *string {
	return s.VendorType
}

func (s *ListVendorsRequest) SetLang(v string) *ListVendorsRequest {
	s.Lang = &v
	return s
}

func (s *ListVendorsRequest) SetMaxResults(v int32) *ListVendorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVendorsRequest) SetNextToken(v string) *ListVendorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVendorsRequest) SetRegionId(v string) *ListVendorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVendorsRequest) SetRoleFor(v int64) *ListVendorsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListVendorsRequest) SetVendorIds(v []*string) *ListVendorsRequest {
	s.VendorIds = v
	return s
}

func (s *ListVendorsRequest) SetVendorName(v string) *ListVendorsRequest {
	s.VendorName = &v
	return s
}

func (s *ListVendorsRequest) SetVendorType(v string) *ListVendorsRequest {
	s.VendorType = &v
	return s
}

func (s *ListVendorsRequest) Validate() error {
	return dara.Validate(s)
}
