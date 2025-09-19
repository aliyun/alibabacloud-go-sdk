// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBlockRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeCustomBlockRecordsResponseBodyPageInfo) *DescribeCustomBlockRecordsResponseBody
	GetPageInfo() *DescribeCustomBlockRecordsResponseBodyPageInfo
	SetRecordList(v []*DescribeCustomBlockRecordsResponseBodyRecordList) *DescribeCustomBlockRecordsResponseBody
	GetRecordList() []*DescribeCustomBlockRecordsResponseBodyRecordList
	SetRequestId(v string) *DescribeCustomBlockRecordsResponseBody
	GetRequestId() *string
}

type DescribeCustomBlockRecordsResponseBody struct {
	// The pagination information.
	PageInfo *DescribeCustomBlockRecordsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the defense rules.
	RecordList []*DescribeCustomBlockRecordsResponseBodyRecordList `json:"RecordList,omitempty" xml:"RecordList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6B7T8Y
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomBlockRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockRecordsResponseBody) GetPageInfo() *DescribeCustomBlockRecordsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeCustomBlockRecordsResponseBody) GetRecordList() []*DescribeCustomBlockRecordsResponseBodyRecordList {
	return s.RecordList
}

func (s *DescribeCustomBlockRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomBlockRecordsResponseBody) SetPageInfo(v *DescribeCustomBlockRecordsResponseBodyPageInfo) *DescribeCustomBlockRecordsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBody) SetRecordList(v []*DescribeCustomBlockRecordsResponseBodyRecordList) *DescribeCustomBlockRecordsResponseBody {
	s.RecordList = v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBody) SetRequestId(v string) *DescribeCustomBlockRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomBlockRecordsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 708
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCustomBlockRecordsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockRecordsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) SetCount(v int32) *DescribeCustomBlockRecordsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeCustomBlockRecordsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) SetPageSize(v int32) *DescribeCustomBlockRecordsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) SetTotalCount(v int32) *DescribeCustomBlockRecordsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomBlockRecordsResponseBodyRecordList struct {
	// The timestamp generated when the block action on the IP address becomes invalid.
	//
	// example:
	//
	// 1671506882063
	BlockExpireDate *int64 `json:"BlockExpireDate,omitempty" xml:"BlockExpireDate,omitempty"`
	// The blocked IP address.
	//
	// example:
	//
	// 45.227.XX.XX
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// The direction of the traffic that is sent by the blocked IP address. Valid values:
	//
	// 	- **in**
	//
	// 	- **out**
	//
	// example:
	//
	// in
	Bound *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The number of servers for which the defense rule is enabled.
	//
	// example:
	//
	// 4
	EnableCount *int32 `json:"EnableCount,omitempty" xml:"EnableCount,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 353376
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The total number of servers on which the IP address is blocked.
	//
	// example:
	//
	// 6
	ServerCount *int32 `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
	// The source of the defense rule.
	//
	// example:
	//
	// UserRule
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the defense rule against brute-force attacks. Valid values:
	//
	// 	- **0**: invalid.
	//
	// 	- **1**: enabled.
	//
	// 	- **2**: failed.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The servers for which the defense rule is enabled.
	TargetList []*DescribeCustomBlockRecordsResponseBodyRecordListTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
}

func (s DescribeCustomBlockRecordsResponseBodyRecordList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockRecordsResponseBodyRecordList) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetBlockExpireDate() *int64 {
	return s.BlockExpireDate
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetBlockIp() *string {
	return s.BlockIp
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetBound() *string {
	return s.Bound
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetEnableCount() *int32 {
	return s.EnableCount
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetId() *int64 {
	return s.Id
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetServerCount() *int32 {
	return s.ServerCount
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetSource() *string {
	return s.Source
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) GetTargetList() []*DescribeCustomBlockRecordsResponseBodyRecordListTargetList {
	return s.TargetList
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetBlockExpireDate(v int64) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.BlockExpireDate = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetBlockIp(v string) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.BlockIp = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetBound(v string) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.Bound = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetEnableCount(v int32) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.EnableCount = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetId(v int64) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.Id = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetServerCount(v int32) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.ServerCount = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetSource(v string) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.Source = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetStatus(v int32) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.Status = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) SetTargetList(v []*DescribeCustomBlockRecordsResponseBodyRecordListTargetList) *DescribeCustomBlockRecordsResponseBodyRecordList {
	s.TargetList = v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordList) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomBlockRecordsResponseBodyRecordListTargetList struct {
	// The ID of the destination asset.
	//
	// example:
	//
	// 032b618f-b220-4a0d-bd37-fbdc6*******
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the query. Valid values:
	//
	// 	- Set the value to **uuid**.
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeCustomBlockRecordsResponseBodyRecordListTargetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockRecordsResponseBodyRecordListTargetList) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordListTargetList) GetTarget() *string {
	return s.Target
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordListTargetList) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordListTargetList) SetTarget(v string) *DescribeCustomBlockRecordsResponseBodyRecordListTargetList {
	s.Target = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordListTargetList) SetTargetType(v string) *DescribeCustomBlockRecordsResponseBodyRecordListTargetList {
	s.TargetType = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponseBodyRecordListTargetList) Validate() error {
	return dara.Validate(s)
}
