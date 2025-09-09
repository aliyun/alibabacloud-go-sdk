// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShardTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeShardTaskInfoResponseBodyData) *DescribeShardTaskInfoResponseBody
	GetData() *DescribeShardTaskInfoResponseBodyData
	SetRequestId(v string) *DescribeShardTaskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeShardTaskInfoResponseBody
	GetSuccess() *bool
}

type DescribeShardTaskInfoResponseBody struct {
	// Indicates the data that is returned.
	Data *DescribeShardTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates the unique ID of the request. If the request fails, provide this ID for technical support to troubleshoot the failure.
	//
	// example:
	//
	// 5D64DE5944A1E541E0CB908A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeShardTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBody) GetData() *DescribeShardTaskInfoResponseBodyData {
	return s.Data
}

func (s *DescribeShardTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeShardTaskInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeShardTaskInfoResponseBody) SetData(v *DescribeShardTaskInfoResponseBodyData) *DescribeShardTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) SetRequestId(v string) *DescribeShardTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) SetSuccess(v bool) *DescribeShardTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeShardTaskInfoResponseBodyData struct {
	// Indicates the number of remaining days before the tasks to shard tables or convert tables expire.
	//
	// example:
	//
	// 0
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates information about full migration tasks.
	Full *DescribeShardTaskInfoResponseBodyDataFull `json:"Full,omitempty" xml:"Full,omitempty" type:"Struct"`
	// Indicates information about full check tasks.
	FullCheck *DescribeShardTaskInfoResponseBodyDataFullCheck `json:"FullCheck,omitempty" xml:"FullCheck,omitempty" type:"Struct"`
	// Indicates information about full correction tasks.
	FullRevise *DescribeShardTaskInfoResponseBodyDataFullRevise `json:"FullRevise,omitempty" xml:"FullRevise,omitempty" type:"Struct"`
	// Indicates information about incremental data synchronization.
	Increment *DescribeShardTaskInfoResponseBodyDataIncrement `json:"Increment,omitempty" xml:"Increment,omitempty" type:"Struct"`
	// Indicates the incremental data synchronization progress.
	//
	// example:
	//
	// 70
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates check tasks.
	Review *DescribeShardTaskInfoResponseBodyDataReview `json:"Review,omitempty" xml:"Review,omitempty" type:"Struct"`
	// Indicates the name of the table that you convert or shard.
	//
	// example:
	//
	// a1
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// Indicates the current stage of the task.
	//
	// example:
	//
	// 2
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// Indicates the state of the tasks to shard tables or convert tables.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates the name of the table after you convert or shard the table.
	//
	// example:
	//
	// a2
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyData) GetExpired() *string {
	return s.Expired
}

func (s *DescribeShardTaskInfoResponseBodyData) GetFull() *DescribeShardTaskInfoResponseBodyDataFull {
	return s.Full
}

func (s *DescribeShardTaskInfoResponseBodyData) GetFullCheck() *DescribeShardTaskInfoResponseBodyDataFullCheck {
	return s.FullCheck
}

func (s *DescribeShardTaskInfoResponseBodyData) GetFullRevise() *DescribeShardTaskInfoResponseBodyDataFullRevise {
	return s.FullRevise
}

func (s *DescribeShardTaskInfoResponseBodyData) GetIncrement() *DescribeShardTaskInfoResponseBodyDataIncrement {
	return s.Increment
}

func (s *DescribeShardTaskInfoResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *DescribeShardTaskInfoResponseBodyData) GetReview() *DescribeShardTaskInfoResponseBodyDataReview {
	return s.Review
}

func (s *DescribeShardTaskInfoResponseBodyData) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *DescribeShardTaskInfoResponseBodyData) GetStage() *string {
	return s.Stage
}

func (s *DescribeShardTaskInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeShardTaskInfoResponseBodyData) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *DescribeShardTaskInfoResponseBodyData) SetExpired(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFull(v *DescribeShardTaskInfoResponseBodyDataFull) *DescribeShardTaskInfoResponseBodyData {
	s.Full = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFullCheck(v *DescribeShardTaskInfoResponseBodyDataFullCheck) *DescribeShardTaskInfoResponseBodyData {
	s.FullCheck = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFullRevise(v *DescribeShardTaskInfoResponseBodyDataFullRevise) *DescribeShardTaskInfoResponseBodyData {
	s.FullRevise = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetIncrement(v *DescribeShardTaskInfoResponseBodyDataIncrement) *DescribeShardTaskInfoResponseBodyData {
	s.Increment = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetProgress(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetReview(v *DescribeShardTaskInfoResponseBodyDataReview) *DescribeShardTaskInfoResponseBodyData {
	s.Review = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetSourceTableName(v string) *DescribeShardTaskInfoResponseBodyData {
	s.SourceTableName = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetStage(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Stage = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetStatus(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetTargetTableName(v string) *DescribeShardTaskInfoResponseBodyData {
	s.TargetTableName = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeShardTaskInfoResponseBodyDataFull struct {
	// Indicates the number of remaining days before the tasks expire.
	//
	// example:
	//
	// 0
	Expired *int32 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates the progress of the tasks.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates the start time when the tasks are performed.
	//
	// example:
	//
	// 2019-09-16 15:11:53
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of tasks.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	//
	// example:
	//
	// 0
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFull) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFull) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) GetExpired() *int32 {
	return s.Expired
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) GetTps() *int32 {
	return s.Tps
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFull {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) Validate() error {
	return dara.Validate(s)
}

type DescribeShardTaskInfoResponseBodyDataFullCheck struct {
	// Indicates the number of remaining days before the tasks expire.
	//
	// example:
	//
	// 0
	Expired *int32 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates the progress of the tasks.
	//
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates the start time when the tasks are performed.
	//
	// example:
	//
	// 2019-09-16 15:12:53
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of tasks.
	//
	// example:
	//
	// 0
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	//
	// example:
	//
	// 0
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFullCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFullCheck) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) GetExpired() *int32 {
	return s.Expired
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) GetTps() *int32 {
	return s.Tps
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) Validate() error {
	return dara.Validate(s)
}

type DescribeShardTaskInfoResponseBodyDataFullRevise struct {
	// Indicates the number of remaining days before the tasks expire.
	//
	// example:
	//
	// 0
	Expired *int32 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates the progress of the tasks.
	//
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates the start time when the tasks are performed.
	//
	// example:
	//
	// 2019-09-16 15:12:53
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of tasks.
	//
	// example:
	//
	// 0
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	//
	// example:
	//
	// 0
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFullRevise) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFullRevise) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) GetExpired() *int32 {
	return s.Expired
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) GetTps() *int32 {
	return s.Tps
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) Validate() error {
	return dara.Validate(s)
}

type DescribeShardTaskInfoResponseBodyDataIncrement struct {
	// Indicates the latency of the incremental data synchronization.
	//
	// example:
	//
	// 1568617906
	Delay *int32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// Indicates the start time when the incremental data synchronization is performed.
	//
	// example:
	//
	// 2019-09-16 15:12:53
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	//
	// example:
	//
	// 10
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataIncrement) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataIncrement) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) GetDelay() *int32 {
	return s.Delay
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) GetTps() *int32 {
	return s.Tps
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetDelay(v int32) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.Delay = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) Validate() error {
	return dara.Validate(s)
}

type DescribeShardTaskInfoResponseBodyDataReview struct {
	// Indicates the number of remaining days before the tasks expire.
	//
	// example:
	//
	// 0
	Expired *int32 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates the progress of the tasks.
	//
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates the start time when the tasks are performed.
	//
	// example:
	//
	// 2019-09-16 15:12:53
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of tasks.
	//
	// example:
	//
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	//
	// example:
	//
	// 10
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataReview) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataReview) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) GetExpired() *int32 {
	return s.Expired
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) GetTps() *int32 {
	return s.Tps
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataReview {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) Validate() error {
	return dara.Validate(s)
}
