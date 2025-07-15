// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamLocationBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeStreamLocationBlockResponseBody
	GetCount() *int32
	SetPageNum(v int32) *DescribeStreamLocationBlockResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeStreamLocationBlockResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeStreamLocationBlockResponseBody
	GetRequestId() *string
	SetStreamBlockList(v *DescribeStreamLocationBlockResponseBodyStreamBlockList) *DescribeStreamLocationBlockResponseBody
	GetStreamBlockList() *DescribeStreamLocationBlockResponseBodyStreamBlockList
	SetTotalPage(v int32) *DescribeStreamLocationBlockResponseBody
	GetTotalPage() *int32
}

type DescribeStreamLocationBlockResponseBody struct {
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7C69682-7F88-40DD-A198-10D0309E439B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations.
	StreamBlockList *DescribeStreamLocationBlockResponseBodyStreamBlockList `json:"StreamBlockList,omitempty" xml:"StreamBlockList,omitempty" type:"Struct"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeStreamLocationBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamLocationBlockResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamLocationBlockResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeStreamLocationBlockResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeStreamLocationBlockResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStreamLocationBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamLocationBlockResponseBody) GetStreamBlockList() *DescribeStreamLocationBlockResponseBodyStreamBlockList {
	return s.StreamBlockList
}

func (s *DescribeStreamLocationBlockResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeStreamLocationBlockResponseBody) SetCount(v int32) *DescribeStreamLocationBlockResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBody) SetPageNum(v int32) *DescribeStreamLocationBlockResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBody) SetPageSize(v int32) *DescribeStreamLocationBlockResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBody) SetRequestId(v string) *DescribeStreamLocationBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBody) SetStreamBlockList(v *DescribeStreamLocationBlockResponseBodyStreamBlockList) *DescribeStreamLocationBlockResponseBody {
	s.StreamBlockList = v
	return s
}

func (s *DescribeStreamLocationBlockResponseBody) SetTotalPage(v int32) *DescribeStreamLocationBlockResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStreamLocationBlockResponseBodyStreamBlockList struct {
	StreamBlock []*DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock `json:"StreamBlock,omitempty" xml:"StreamBlock,omitempty" type:"Repeated"`
}

func (s DescribeStreamLocationBlockResponseBodyStreamBlockList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamLocationBlockResponseBodyStreamBlockList) GoString() string {
	return s.String()
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockList) GetStreamBlock() []*DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	return s.StreamBlock
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockList) SetStreamBlock(v []*DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) *DescribeStreamLocationBlockResponseBodyStreamBlockList {
	s.StreamBlock = v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockList) Validate() error {
	return dara.Validate(s)
}

type DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock struct {
	// The name of the application in which the blocking applies.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The blocking type. Valid values:
	//
	// 	- blacklist
	//
	// 	- whitelist
	//
	// example:
	//
	// blacklist
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The blocked region. If multiple regions are specified, such as CN and AS, they are separated by commas (,).
	//
	// example:
	//
	// CN
	LocationList *string `json:"LocationList,omitempty" xml:"LocationList,omitempty"`
	// The time when the blocking ends. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-08-13T09:27Z
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The blocking status. Valid values: 0 (not started), 1 (blocking), 2 (successful), 3 (failed), 4 (expired), and 5 (deleting).
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the stream.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The time when the configuration was last modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-08-14T09:27Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GoString() string {
	return s.String()
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GetAppName() *string {
	return s.AppName
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GetBlockType() *string {
	return s.BlockType
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GetLocationList() *string {
	return s.LocationList
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) SetAppName(v string) *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	s.AppName = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) SetBlockType(v string) *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	s.BlockType = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) SetDomainName(v string) *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	s.DomainName = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) SetLocationList(v string) *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	s.LocationList = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) SetReleaseTime(v string) *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) SetStatus(v int32) *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	s.Status = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) SetStreamName(v string) *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	s.StreamName = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) SetUpdateTime(v string) *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock {
	s.UpdateTime = &v
	return s
}

func (s *DescribeStreamLocationBlockResponseBodyStreamBlockListStreamBlock) Validate() error {
	return dara.Validate(s)
}
