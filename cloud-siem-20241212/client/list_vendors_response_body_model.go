// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVendorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVendorsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVendorsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVendorsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVendorsResponseBody
	GetTotalCount() *int32
	SetVendors(v []*ListVendorsResponseBodyVendors) *ListVendorsResponseBody
	GetVendors() []*ListVendorsResponseBodyVendors
}

type ListVendorsResponseBody struct {
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
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57。
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vendors    []*ListVendorsResponseBodyVendors `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s ListVendorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVendorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVendorsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVendorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVendorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVendorsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVendorsResponseBody) GetVendors() []*ListVendorsResponseBodyVendors {
	return s.Vendors
}

func (s *ListVendorsResponseBody) SetMaxResults(v int32) *ListVendorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVendorsResponseBody) SetNextToken(v string) *ListVendorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVendorsResponseBody) SetRequestId(v string) *ListVendorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVendorsResponseBody) SetTotalCount(v int32) *ListVendorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVendorsResponseBody) SetVendors(v []*ListVendorsResponseBodyVendors) *ListVendorsResponseBody {
	s.Vendors = v
	return s
}

func (s *ListVendorsResponseBody) Validate() error {
	if s.Vendors != nil {
		for _, item := range s.Vendors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVendorsResponseBodyVendors struct {
	// example:
	//
	// 1733269771123。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1733269771123。
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// vd-qlsw5eocx94w9。
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
	// example:
	//
	// 111。
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
	// example:
	//
	// preset。
	VendorType *string `json:"VendorType,omitempty" xml:"VendorType,omitempty"`
}

func (s ListVendorsResponseBodyVendors) String() string {
	return dara.Prettify(s)
}

func (s ListVendorsResponseBodyVendors) GoString() string {
	return s.String()
}

func (s *ListVendorsResponseBodyVendors) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListVendorsResponseBodyVendors) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListVendorsResponseBodyVendors) GetVendorId() *string {
	return s.VendorId
}

func (s *ListVendorsResponseBodyVendors) GetVendorName() *string {
	return s.VendorName
}

func (s *ListVendorsResponseBodyVendors) GetVendorType() *string {
	return s.VendorType
}

func (s *ListVendorsResponseBodyVendors) SetCreateTime(v int64) *ListVendorsResponseBodyVendors {
	s.CreateTime = &v
	return s
}

func (s *ListVendorsResponseBodyVendors) SetUpdateTime(v int64) *ListVendorsResponseBodyVendors {
	s.UpdateTime = &v
	return s
}

func (s *ListVendorsResponseBodyVendors) SetVendorId(v string) *ListVendorsResponseBodyVendors {
	s.VendorId = &v
	return s
}

func (s *ListVendorsResponseBodyVendors) SetVendorName(v string) *ListVendorsResponseBodyVendors {
	s.VendorName = &v
	return s
}

func (s *ListVendorsResponseBodyVendors) SetVendorType(v string) *ListVendorsResponseBodyVendors {
	s.VendorType = &v
	return s
}

func (s *ListVendorsResponseBodyVendors) Validate() error {
	return dara.Validate(s)
}
