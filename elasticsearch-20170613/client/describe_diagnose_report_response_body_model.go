// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnoseReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDiagnoseReportResponseBody
	GetRequestId() *string
	SetResult(v *DescribeDiagnoseReportResponseBodyResult) *DescribeDiagnoseReportResponseBody
	GetResult() *DescribeDiagnoseReportResponseBodyResult
}

type DescribeDiagnoseReportResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDiagnoseReportResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDiagnoseReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnoseReportResponseBody) GetResult() *DescribeDiagnoseReportResponseBodyResult {
	return s.Result
}

func (s *DescribeDiagnoseReportResponseBody) SetRequestId(v string) *DescribeDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBody) SetResult(v *DescribeDiagnoseReportResponseBodyResult) *DescribeDiagnoseReportResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDiagnoseReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDiagnoseReportResponseBodyResult struct {
	// example:
	//
	// 1535745731000
	CreateTime    *int64                                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiagnoseItems []*DescribeDiagnoseReportResponseBodyResultDiagnoseItems `json:"diagnoseItems,omitempty" xml:"diagnoseItems,omitempty" type:"Repeated"`
	// example:
	//
	// YELLOW
	Health *string `json:"health,omitempty" xml:"health,omitempty"`
	// example:
	//
	// es-cn-abc
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// trigger__2020-08-17T17:09:02
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// example:
	//
	// SUCCESS
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// SYSTEM
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s DescribeDiagnoseReportResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeDiagnoseReportResponseBodyResult) GetDiagnoseItems() []*DescribeDiagnoseReportResponseBodyResultDiagnoseItems {
	return s.DiagnoseItems
}

func (s *DescribeDiagnoseReportResponseBodyResult) GetHealth() *string {
	return s.Health
}

func (s *DescribeDiagnoseReportResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDiagnoseReportResponseBodyResult) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeDiagnoseReportResponseBodyResult) GetState() *string {
	return s.State
}

func (s *DescribeDiagnoseReportResponseBodyResult) GetTrigger() *string {
	return s.Trigger
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetCreateTime(v int64) *DescribeDiagnoseReportResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetDiagnoseItems(v []*DescribeDiagnoseReportResponseBodyResultDiagnoseItems) *DescribeDiagnoseReportResponseBodyResult {
	s.DiagnoseItems = v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetHealth(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.Health = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetInstanceId(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetReportId(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetState(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetTrigger(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.Trigger = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type DescribeDiagnoseReportResponseBodyResultDiagnoseItems struct {
	Detail *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
	// example:
	//
	// YELLOW
	Health *string `json:"health,omitempty" xml:"health,omitempty"`
	// example:
	//
	// IndexAliasUseDiagnostic
	Item *string `json:"item,omitempty" xml:"item,omitempty"`
}

func (s DescribeDiagnoseReportResponseBodyResultDiagnoseItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBodyResultDiagnoseItems) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) GetDetail() *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	return s.Detail
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) GetHealth() *string {
	return s.Health
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) GetItem() *string {
	return s.Item
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) SetDetail(v *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) *DescribeDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Detail = v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) SetHealth(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Health = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) SetItem(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Item = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail struct {
	// example:
	//
	// Check whether the number of replica shards is optimal and easy to maintain
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// Number of Replica Shards
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// You may need to adjust the numbers of replica shards of some indices as follows: [geoname08 : 0 -&gt; 1][geoname09 : 0 -&gt; 1][geonametest01 : 0 -&gt; 1]
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// You can call the following function in the Elasticsearch API....
	Suggest *string `json:"suggest,omitempty" xml:"suggest,omitempty"`
	// example:
	//
	// ES_API
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetDesc() *string {
	return s.Desc
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetName() *string {
	return s.Name
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetResult() *string {
	return s.Result
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetSuggest() *string {
	return s.Suggest
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetType() *string {
	return s.Type
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetDesc(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Desc = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetName(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Name = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetResult(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Result = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetSuggest(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Suggest = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetType(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Type = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) Validate() error {
	return dara.Validate(s)
}
