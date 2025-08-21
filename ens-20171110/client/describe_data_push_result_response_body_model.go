// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataPushResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDataPushResultResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataPushResultResponseBody
	GetPageSize() *int32
	SetPushResults(v *DescribeDataPushResultResponseBodyPushResults) *DescribeDataPushResultResponseBody
	GetPushResults() *DescribeDataPushResultResponseBodyPushResults
	SetRequestId(v string) *DescribeDataPushResultResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataPushResultResponseBody
	GetTotalCount() *int32
}

type DescribeDataPushResultResponseBody struct {
	// The page number.
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
	// The push results of data files.
	PushResults *DescribeDataPushResultResponseBodyPushResults `json:"PushResults,omitempty" xml:"PushResults,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataPushResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataPushResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataPushResultResponseBody) GetPushResults() *DescribeDataPushResultResponseBodyPushResults {
	return s.PushResults
}

func (s *DescribeDataPushResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataPushResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataPushResultResponseBody) SetPageNumber(v int32) *DescribeDataPushResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataPushResultResponseBody) SetPageSize(v int32) *DescribeDataPushResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataPushResultResponseBody) SetPushResults(v *DescribeDataPushResultResponseBodyPushResults) *DescribeDataPushResultResponseBody {
	s.PushResults = v
	return s
}

func (s *DescribeDataPushResultResponseBody) SetRequestId(v string) *DescribeDataPushResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataPushResultResponseBody) SetTotalCount(v int32) *DescribeDataPushResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataPushResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDataPushResultResponseBodyPushResults struct {
	PushResult []*DescribeDataPushResultResponseBodyPushResultsPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Repeated"`
}

func (s DescribeDataPushResultResponseBodyPushResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResults) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResults) GetPushResult() []*DescribeDataPushResultResponseBodyPushResultsPushResult {
	return s.PushResult
}

func (s *DescribeDataPushResultResponseBodyPushResults) SetPushResult(v []*DescribeDataPushResultResponseBodyPushResultsPushResult) *DescribeDataPushResultResponseBodyPushResults {
	s.PushResult = v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResults) Validate() error {
	return dara.Validate(s)
}

type DescribeDataPushResultResponseBodyPushResultsPushResult struct {
	// The name of the data file.
	//
	// example:
	//
	// gcs-pre-websocket-eip-telecom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The push status of data files.
	StatusStatS *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS `json:"StatusStatS,omitempty" xml:"StatusStatS,omitempty" type:"Struct"`
	// The version number of the data file.
	//
	// example:
	//
	// standard
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResult) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) GetName() *string {
	return s.Name
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) GetStatusStatS() *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS {
	return s.StatusStatS
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) GetVersion() *string {
	return s.Version
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) SetName(v string) *DescribeDataPushResultResponseBodyPushResultsPushResult {
	s.Name = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) SetStatusStatS(v *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) *DescribeDataPushResultResponseBodyPushResultsPushResult {
	s.StatusStatS = v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) SetVersion(v string) *DescribeDataPushResultResponseBodyPushResultsPushResult {
	s.Version = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) Validate() error {
	return dara.Validate(s)
}

type DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS struct {
	StatusStat []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat `json:"StatusStat,omitempty" xml:"StatusStat,omitempty" type:"Repeated"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) GetStatusStat() []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat {
	return s.StatusStat
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) SetStatusStat(v []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS {
	s.StatusStat = v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) Validate() error {
	return dara.Validate(s)
}

type DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat struct {
	// The total number of ENS nodes.
	//
	// example:
	//
	// 2
	RegionIdCount *int32 `json:"RegionIdCount,omitempty" xml:"RegionIdCount,omitempty"`
	// The push status of data files on the ENS node.
	RegionIds *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
	// The push status. The value is of the enumeration type. Valid values: SUCCESS FAILED PUSHING
	//
	// example:
	//
	// FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) GetRegionIdCount() *int32 {
	return s.RegionIdCount
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) GetRegionIds() *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds {
	return s.RegionIds
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) SetRegionIdCount(v int32) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat {
	s.RegionIdCount = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) SetRegionIds(v *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat {
	s.RegionIds = v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) SetStatus(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat {
	s.Status = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) Validate() error {
	return dara.Validate(s)
}

type DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds struct {
	RegionId []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) GetRegionId() []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	return s.RegionId
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) SetRegionId(v []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds {
	s.RegionId = v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) Validate() error {
	return dara.Validate(s)
}

type DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId struct {
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-wuxi-5
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the push operation. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-28T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The description of the status.
	//
	// example:
	//
	// Success
	StatusDescrip *string `json:"StatusDescrip,omitempty" xml:"StatusDescrip,omitempty"`
	// The time when the status was last updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-07T02:37:10Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) GetStatusDescrip() *string {
	return s.StatusDescrip
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetRegionId(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.RegionId = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetStartTime(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.StartTime = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetStatusDescrip(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.StatusDescrip = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetUpdateTime(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) Validate() error {
	return dara.Validate(s)
}
