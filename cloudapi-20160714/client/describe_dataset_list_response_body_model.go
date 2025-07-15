// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetInfoList(v []*DescribeDatasetListResponseBodyDatasetInfoList) *DescribeDatasetListResponseBody
	GetDatasetInfoList() []*DescribeDatasetListResponseBodyDatasetInfoList
	SetPageNumber(v int32) *DescribeDatasetListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatasetListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDatasetListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDatasetListResponseBody
	GetTotalCount() *int32
}

type DescribeDatasetListResponseBody struct {
	// The returned dataset information. It is an array consisting of datasetinfo.
	DatasetInfoList []*DescribeDatasetListResponseBodyDatasetInfoList `json:"DatasetInfoList,omitempty" xml:"DatasetInfoList,omitempty" type:"Repeated"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
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
	// The ID of the request.
	//
	// example:
	//
	// D403C6E6-21B3-5B78-82DA-E3B6********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatasetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatasetListResponseBody) GetDatasetInfoList() []*DescribeDatasetListResponseBodyDatasetInfoList {
	return s.DatasetInfoList
}

func (s *DescribeDatasetListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatasetListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatasetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatasetListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDatasetListResponseBody) SetDatasetInfoList(v []*DescribeDatasetListResponseBodyDatasetInfoList) *DescribeDatasetListResponseBody {
	s.DatasetInfoList = v
	return s
}

func (s *DescribeDatasetListResponseBody) SetPageNumber(v int32) *DescribeDatasetListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatasetListResponseBody) SetPageSize(v int32) *DescribeDatasetListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatasetListResponseBody) SetRequestId(v string) *DescribeDatasetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatasetListResponseBody) SetTotalCount(v int32) *DescribeDatasetListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDatasetListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDatasetListResponseBodyDatasetInfoList struct {
	// The time when the dataset was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-21T12:58:43Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// 6304ce6b4ae6453f********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The dataset type. Valid values:
	//
	// 	- JWT_BLOCKING : a JSON Web Token (JWT) blacklist
	//
	// 	- IP_WHITELIST_CIDR : an IP address whitelist
	//
	// 	- PARAMETER_ACCESS: a list of parameters for parameter-based access control
	//
	// example:
	//
	// IP_WHITELIST_CIDR
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the dataset was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-21T12:58:43Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The tags.
	Tags []*DescribeDatasetListResponseBodyDatasetInfoListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDatasetListResponseBodyDatasetInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetListResponseBodyDatasetInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) GetDatasetType() *string {
	return s.DatasetType
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) GetDescription() *string {
	return s.Description
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) GetTags() []*DescribeDatasetListResponseBodyDatasetInfoListTags {
	return s.Tags
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) SetCreatedTime(v string) *DescribeDatasetListResponseBodyDatasetInfoList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) SetDatasetId(v string) *DescribeDatasetListResponseBodyDatasetInfoList {
	s.DatasetId = &v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) SetDatasetName(v string) *DescribeDatasetListResponseBodyDatasetInfoList {
	s.DatasetName = &v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) SetDatasetType(v string) *DescribeDatasetListResponseBodyDatasetInfoList {
	s.DatasetType = &v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) SetDescription(v string) *DescribeDatasetListResponseBodyDatasetInfoList {
	s.Description = &v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) SetModifiedTime(v string) *DescribeDatasetListResponseBodyDatasetInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) SetTags(v []*DescribeDatasetListResponseBodyDatasetInfoListTags) *DescribeDatasetListResponseBodyDatasetInfoList {
	s.Tags = v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeDatasetListResponseBodyDatasetInfoListTags struct {
	// The tag key.
	//
	// example:
	//
	// ENV
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDatasetListResponseBodyDatasetInfoListTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetListResponseBodyDatasetInfoListTags) GoString() string {
	return s.String()
}

func (s *DescribeDatasetListResponseBodyDatasetInfoListTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDatasetListResponseBodyDatasetInfoListTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDatasetListResponseBodyDatasetInfoListTags) SetKey(v string) *DescribeDatasetListResponseBodyDatasetInfoListTags {
	s.Key = &v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoListTags) SetValue(v string) *DescribeDatasetListResponseBodyDatasetInfoListTags {
	s.Value = &v
	return s
}

func (s *DescribeDatasetListResponseBodyDatasetInfoListTags) Validate() error {
	return dara.Validate(s)
}
