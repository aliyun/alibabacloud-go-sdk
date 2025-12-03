// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPrivateDNSResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPrivateDNSResponseBody
	GetPageSize() *int32
	SetPrivateDNSList(v []*ListPrivateDNSResponseBodyPrivateDNSList) *ListPrivateDNSResponseBody
	GetPrivateDNSList() []*ListPrivateDNSResponseBodyPrivateDNSList
	SetRequestId(v string) *ListPrivateDNSResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPrivateDNSResponseBody
	GetTotalCount() *int32
}

type ListPrivateDNSResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The internal domain name resolutions.
	PrivateDNSList []*ListPrivateDNSResponseBodyPrivateDNSList `json:"PrivateDNSList,omitempty" xml:"PrivateDNSList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPrivateDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateDNSResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateDNSResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPrivateDNSResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPrivateDNSResponseBody) GetPrivateDNSList() []*ListPrivateDNSResponseBodyPrivateDNSList {
	return s.PrivateDNSList
}

func (s *ListPrivateDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateDNSResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPrivateDNSResponseBody) SetPageNumber(v int32) *ListPrivateDNSResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPrivateDNSResponseBody) SetPageSize(v int32) *ListPrivateDNSResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPrivateDNSResponseBody) SetPrivateDNSList(v []*ListPrivateDNSResponseBodyPrivateDNSList) *ListPrivateDNSResponseBody {
	s.PrivateDNSList = v
	return s
}

func (s *ListPrivateDNSResponseBody) SetRequestId(v string) *ListPrivateDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateDNSResponseBody) SetTotalCount(v int32) *ListPrivateDNSResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPrivateDNSResponseBody) Validate() error {
	if s.PrivateDNSList != nil {
		for _, item := range s.PrivateDNSList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivateDNSResponseBodyPrivateDNSList struct {
	// The instances that are associated with the resolution.
	BindInstances []*string `json:"BindInstances,omitempty" xml:"BindInstances,omitempty" type:"Repeated"`
	// The time when the resolution was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-10T08:17:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The internal domain name.
	//
	// example:
	//
	// api.demo.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// The resolution records.
	Records []*ListPrivateDNSResponseBodyPrivateDNSListRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The internal domain name resolution type. Valid values:
	//
	// 	- VPC: resolution for VPC access authorizations. A resolution of this type can be bound only to traditional dedicated instances.
	//
	// 	- A: resolution that supports A records. A resolution of this type can be bound only to VPC integration dedicated instances.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrivateDNSResponseBodyPrivateDNSList) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateDNSResponseBodyPrivateDNSList) GoString() string {
	return s.String()
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) GetBindInstances() []*string {
	return s.BindInstances
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) GetRecords() []*ListPrivateDNSResponseBodyPrivateDNSListRecords {
	return s.Records
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) GetType() *string {
	return s.Type
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) SetBindInstances(v []*string) *ListPrivateDNSResponseBodyPrivateDNSList {
	s.BindInstances = v
	return s
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) SetCreatedTime(v string) *ListPrivateDNSResponseBodyPrivateDNSList {
	s.CreatedTime = &v
	return s
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) SetIntranetDomain(v string) *ListPrivateDNSResponseBodyPrivateDNSList {
	s.IntranetDomain = &v
	return s
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) SetRecords(v []*ListPrivateDNSResponseBodyPrivateDNSListRecords) *ListPrivateDNSResponseBodyPrivateDNSList {
	s.Records = v
	return s
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) SetType(v string) *ListPrivateDNSResponseBodyPrivateDNSList {
	s.Type = &v
	return s
}

func (s *ListPrivateDNSResponseBodyPrivateDNSList) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivateDNSResponseBodyPrivateDNSListRecords struct {
	// The resolution record.
	//
	// example:
	//
	// 192.168.0.1
	Record *string `json:"Record,omitempty" xml:"Record,omitempty"`
	// The weight of the record.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListPrivateDNSResponseBodyPrivateDNSListRecords) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateDNSResponseBodyPrivateDNSListRecords) GoString() string {
	return s.String()
}

func (s *ListPrivateDNSResponseBodyPrivateDNSListRecords) GetRecord() *string {
	return s.Record
}

func (s *ListPrivateDNSResponseBodyPrivateDNSListRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *ListPrivateDNSResponseBodyPrivateDNSListRecords) SetRecord(v string) *ListPrivateDNSResponseBodyPrivateDNSListRecords {
	s.Record = &v
	return s
}

func (s *ListPrivateDNSResponseBodyPrivateDNSListRecords) SetWeight(v int32) *ListPrivateDNSResponseBodyPrivateDNSListRecords {
	s.Weight = &v
	return s
}

func (s *ListPrivateDNSResponseBodyPrivateDNSListRecords) Validate() error {
	return dara.Validate(s)
}
