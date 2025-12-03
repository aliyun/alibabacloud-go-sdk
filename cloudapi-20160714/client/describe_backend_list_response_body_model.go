// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendInfoList(v []*DescribeBackendListResponseBodyBackendInfoList) *DescribeBackendListResponseBody
	GetBackendInfoList() []*DescribeBackendListResponseBodyBackendInfoList
	SetPageNumber(v int32) *DescribeBackendListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackendListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackendListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBackendListResponseBody
	GetTotalCount() *int32
}

type DescribeBackendListResponseBody struct {
	// The backend services.
	BackendInfoList []*DescribeBackendListResponseBodyBackendInfoList `json:"BackendInfoList,omitempty" xml:"BackendInfoList,omitempty" type:"Repeated"`
	// The number of the current page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 090A0DF9-9144-5236-8CBA-E18DE317722D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackendListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackendListResponseBody) GetBackendInfoList() []*DescribeBackendListResponseBodyBackendInfoList {
	return s.BackendInfoList
}

func (s *DescribeBackendListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackendListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackendListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackendListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBackendListResponseBody) SetBackendInfoList(v []*DescribeBackendListResponseBodyBackendInfoList) *DescribeBackendListResponseBody {
	s.BackendInfoList = v
	return s
}

func (s *DescribeBackendListResponseBody) SetPageNumber(v int32) *DescribeBackendListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackendListResponseBody) SetPageSize(v int32) *DescribeBackendListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackendListResponseBody) SetRequestId(v string) *DescribeBackendListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackendListResponseBody) SetTotalCount(v int32) *DescribeBackendListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackendListResponseBody) Validate() error {
	if s.BackendInfoList != nil {
		for _, item := range s.BackendInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackendListResponseBodyBackendInfoList struct {
	// The ID of the backend service.
	//
	// example:
	//
	// 35bd31d32c9c425ebbe9330db9f8c375
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// The name of the backend service.
	//
	// example:
	//
	// test
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// The type of the backend service.
	//
	// example:
	//
	// HTTP
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
	// The time when the backend service was created.
	//
	// example:
	//
	// 2022-01-25T11:22:29Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the backend service.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the backend service was modified.
	//
	// example:
	//
	// 2022-01-25T11:22:29Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The list of tags.
	Tags []*DescribeBackendListResponseBodyBackendInfoListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeBackendListResponseBodyBackendInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendListResponseBodyBackendInfoList) GoString() string {
	return s.String()
}

func (s *DescribeBackendListResponseBodyBackendInfoList) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeBackendListResponseBodyBackendInfoList) GetBackendName() *string {
	return s.BackendName
}

func (s *DescribeBackendListResponseBodyBackendInfoList) GetBackendType() *string {
	return s.BackendType
}

func (s *DescribeBackendListResponseBodyBackendInfoList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeBackendListResponseBodyBackendInfoList) GetDescription() *string {
	return s.Description
}

func (s *DescribeBackendListResponseBodyBackendInfoList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeBackendListResponseBodyBackendInfoList) GetTags() []*DescribeBackendListResponseBodyBackendInfoListTags {
	return s.Tags
}

func (s *DescribeBackendListResponseBodyBackendInfoList) SetBackendId(v string) *DescribeBackendListResponseBodyBackendInfoList {
	s.BackendId = &v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoList) SetBackendName(v string) *DescribeBackendListResponseBodyBackendInfoList {
	s.BackendName = &v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoList) SetBackendType(v string) *DescribeBackendListResponseBodyBackendInfoList {
	s.BackendType = &v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoList) SetCreatedTime(v string) *DescribeBackendListResponseBodyBackendInfoList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoList) SetDescription(v string) *DescribeBackendListResponseBodyBackendInfoList {
	s.Description = &v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoList) SetModifiedTime(v string) *DescribeBackendListResponseBodyBackendInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoList) SetTags(v []*DescribeBackendListResponseBodyBackendInfoListTags) *DescribeBackendListResponseBodyBackendInfoList {
	s.Tags = v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoList) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackendListResponseBodyBackendInfoListTags struct {
	// The name of the tag.
	//
	// example:
	//
	// groupName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// e3b881d0-e2d0-4dfb-b1fb-a2a3d1e534b7
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackendListResponseBodyBackendInfoListTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendListResponseBodyBackendInfoListTags) GoString() string {
	return s.String()
}

func (s *DescribeBackendListResponseBodyBackendInfoListTags) GetKey() *string {
	return s.Key
}

func (s *DescribeBackendListResponseBodyBackendInfoListTags) GetValue() *string {
	return s.Value
}

func (s *DescribeBackendListResponseBodyBackendInfoListTags) SetKey(v string) *DescribeBackendListResponseBodyBackendInfoListTags {
	s.Key = &v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoListTags) SetValue(v string) *DescribeBackendListResponseBodyBackendInfoListTags {
	s.Value = &v
	return s
}

func (s *DescribeBackendListResponseBodyBackendInfoListTags) Validate() error {
	return dara.Validate(s)
}
