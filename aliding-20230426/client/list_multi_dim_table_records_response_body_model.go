// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiDimTableRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *ListMultiDimTableRecordsResponseBody
	GetHasMore() *bool
	SetNextToken(v string) *ListMultiDimTableRecordsResponseBody
	GetNextToken() *string
	SetRecords(v []*ListMultiDimTableRecordsResponseBodyRecords) *ListMultiDimTableRecordsResponseBody
	GetRecords() []*ListMultiDimTableRecordsResponseBodyRecords
	SetRequestId(v string) *ListMultiDimTableRecordsResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *ListMultiDimTableRecordsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListMultiDimTableRecordsResponseBody
	GetVendorType() *string
}

type ListMultiDimTableRecordsResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 1234567890
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// []
	Records []*ListMultiDimTableRecordsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListMultiDimTableRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMultiDimTableRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiDimTableRecordsResponseBody) GetRecords() []*ListMultiDimTableRecordsResponseBodyRecords {
	return s.Records
}

func (s *ListMultiDimTableRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultiDimTableRecordsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListMultiDimTableRecordsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListMultiDimTableRecordsResponseBody) SetHasMore(v bool) *ListMultiDimTableRecordsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBody) SetNextToken(v string) *ListMultiDimTableRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBody) SetRecords(v []*ListMultiDimTableRecordsResponseBodyRecords) *ListMultiDimTableRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListMultiDimTableRecordsResponseBody) SetRequestId(v string) *ListMultiDimTableRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBody) SetVendorRequestId(v string) *ListMultiDimTableRecordsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBody) SetVendorType(v string) *ListMultiDimTableRecordsResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBody) Validate() error {
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

type ListMultiDimTableRecordsResponseBodyRecords struct {
	CreatedBy *ListMultiDimTableRecordsResponseBodyRecordsCreatedBy `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// title,shortTitle
	Fields map[string]interface{} `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// example:
	//
	// 123
	Id             *string                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	LastModifiedBy *ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy `json:"LastModifiedBy,omitempty" xml:"LastModifiedBy,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
}

func (s ListMultiDimTableRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) GetCreatedBy() *ListMultiDimTableRecordsResponseBodyRecordsCreatedBy {
	return s.CreatedBy
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) GetFields() map[string]interface{} {
	return s.Fields
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) GetId() *string {
	return s.Id
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) GetLastModifiedBy() *ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy {
	return s.LastModifiedBy
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) SetCreatedBy(v *ListMultiDimTableRecordsResponseBodyRecordsCreatedBy) *ListMultiDimTableRecordsResponseBodyRecords {
	s.CreatedBy = v
	return s
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) SetCreatedTime(v int64) *ListMultiDimTableRecordsResponseBodyRecords {
	s.CreatedTime = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) SetFields(v map[string]interface{}) *ListMultiDimTableRecordsResponseBodyRecords {
	s.Fields = v
	return s
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) SetId(v string) *ListMultiDimTableRecordsResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) SetLastModifiedBy(v *ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy) *ListMultiDimTableRecordsResponseBodyRecords {
	s.LastModifiedBy = v
	return s
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) SetLastModifiedTime(v int64) *ListMultiDimTableRecordsResponseBodyRecords {
	s.LastModifiedTime = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBodyRecords) Validate() error {
	if s.CreatedBy != nil {
		if err := s.CreatedBy.Validate(); err != nil {
			return err
		}
	}
	if s.LastModifiedBy != nil {
		if err := s.LastModifiedBy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMultiDimTableRecordsResponseBodyRecordsCreatedBy struct {
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMultiDimTableRecordsResponseBodyRecordsCreatedBy) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsResponseBodyRecordsCreatedBy) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsResponseBodyRecordsCreatedBy) GetUserId() *string {
	return s.UserId
}

func (s *ListMultiDimTableRecordsResponseBodyRecordsCreatedBy) SetUserId(v string) *ListMultiDimTableRecordsResponseBodyRecordsCreatedBy {
	s.UserId = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBodyRecordsCreatedBy) Validate() error {
	return dara.Validate(s)
}

type ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy struct {
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy) GetUserId() *string {
	return s.UserId
}

func (s *ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy) SetUserId(v string) *ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy {
	s.UserId = &v
	return s
}

func (s *ListMultiDimTableRecordsResponseBodyRecordsLastModifiedBy) Validate() error {
	return dara.Validate(s)
}
