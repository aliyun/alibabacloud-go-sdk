// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListDiagnoseReportResponseBodyHeaders) *ListDiagnoseReportResponseBody
	GetHeaders() *ListDiagnoseReportResponseBodyHeaders
	SetRequestId(v string) *ListDiagnoseReportResponseBody
	GetRequestId() *string
	SetResult(v []*ListDiagnoseReportResponseBodyResult) *ListDiagnoseReportResponseBody
	GetResult() []*ListDiagnoseReportResponseBodyResult
}

type ListDiagnoseReportResponseBody struct {
	// The total number of entries returned.
	Headers *ListDiagnoseReportResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The header of the response.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trigger mode of health diagnostics. Valid values:
	//
	// 	- SYSTEM: The system is automatically triggered.
	//
	// 	- INNER: internal trigger
	//
	// 	- USER: manually triggered by the user
	Result []*ListDiagnoseReportResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiagnoseReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBody) GetHeaders() *ListDiagnoseReportResponseBodyHeaders {
	return s.Headers
}

func (s *ListDiagnoseReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnoseReportResponseBody) GetResult() []*ListDiagnoseReportResponseBodyResult {
	return s.Result
}

func (s *ListDiagnoseReportResponseBody) SetHeaders(v *ListDiagnoseReportResponseBodyHeaders) *ListDiagnoseReportResponseBody {
	s.Headers = v
	return s
}

func (s *ListDiagnoseReportResponseBody) SetRequestId(v string) *ListDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnoseReportResponseBody) SetResult(v []*ListDiagnoseReportResponseBodyResult) *ListDiagnoseReportResponseBody {
	s.Result = v
	return s
}

func (s *ListDiagnoseReportResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDiagnoseReportResponseBodyHeaders struct {
	// The returned results.
	//
	// example:
	//
	// 15
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListDiagnoseReportResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListDiagnoseReportResponseBodyHeaders) SetXTotalCount(v int32) *ListDiagnoseReportResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListDiagnoseReportResponseBodyResult struct {
	// The ID of the report.
	//
	// example:
	//
	// 1535745731000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the item.
	DiagnoseItems []*ListDiagnoseReportResponseBodyResultDiagnoseItems `json:"diagnoseItems,omitempty" xml:"diagnoseItems,omitempty" type:"Repeated"`
	// Reports the list of diagnostic item information.
	//
	// example:
	//
	// YELLOW
	Health *string `json:"health,omitempty" xml:"health,omitempty"`
	// The overall health of the cluster in the report. Supported: GREEN, YELLOW, RED, and UNKNOWN.
	//
	// example:
	//
	// es-cn-abc
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The diagnosis status. Valid values: Supported: SUCCESS, FAILED, and RUNNING.
	//
	// example:
	//
	// trigger__2020-08-17T17:09:02f
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The ID of the instance for diagnosis.
	//
	// example:
	//
	// SUCCESS
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The timestamp when the report was created.
	//
	// example:
	//
	// USER
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s ListDiagnoseReportResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDiagnoseReportResponseBodyResult) GetDiagnoseItems() []*ListDiagnoseReportResponseBodyResultDiagnoseItems {
	return s.DiagnoseItems
}

func (s *ListDiagnoseReportResponseBodyResult) GetHealth() *string {
	return s.Health
}

func (s *ListDiagnoseReportResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDiagnoseReportResponseBodyResult) GetReportId() *string {
	return s.ReportId
}

func (s *ListDiagnoseReportResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListDiagnoseReportResponseBodyResult) GetTrigger() *string {
	return s.Trigger
}

func (s *ListDiagnoseReportResponseBodyResult) SetCreateTime(v int64) *ListDiagnoseReportResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetDiagnoseItems(v []*ListDiagnoseReportResponseBodyResultDiagnoseItems) *ListDiagnoseReportResponseBodyResult {
	s.DiagnoseItems = v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetHealth(v string) *ListDiagnoseReportResponseBodyResult {
	s.Health = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetInstanceId(v string) *ListDiagnoseReportResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetReportId(v string) *ListDiagnoseReportResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetState(v string) *ListDiagnoseReportResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetTrigger(v string) *ListDiagnoseReportResponseBodyResult {
	s.Trigger = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) Validate() error {
	if s.DiagnoseItems != nil {
		for _, item := range s.DiagnoseItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDiagnoseReportResponseBodyResultDiagnoseItems struct {
	// The type of the diagnostic result. Valid values:
	//
	// 	- TEXT: text description
	//
	// 	- CONSOLE_API: console-triggered
	//
	// 	- ES_API: API triggered
	Detail *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
	// The details of the diagnostic item.
	//
	// example:
	//
	// YELLOW
	Health *string `json:"health,omitempty" xml:"health,omitempty"`
	// The health of the diagnostic item. Supported: GREEN, YELLOW, RED, and UNKNOWN.
	//
	// example:
	//
	// IndexAliasUseDiagnostic
	Item *string `json:"item,omitempty" xml:"item,omitempty"`
}

func (s ListDiagnoseReportResponseBodyResultDiagnoseItems) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportResponseBodyResultDiagnoseItems) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) GetDetail() *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	return s.Detail
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) GetHealth() *string {
	return s.Health
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) GetItem() *string {
	return s.Item
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) SetDetail(v *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) *ListDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Detail = v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) SetHealth(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Health = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) SetItem(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Item = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) Validate() error {
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail struct {
	// The diagnosis.
	//
	// example:
	//
	// Check whether the number of replica shards is optimal and easy to maintain
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The description of the diagnostic item.
	//
	// example:
	//
	// Number of Replica Shards
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The suggestion for the diagnosis.
	//
	// example:
	//
	// You may need to adjust the numbers of replica shards of some indices as follows:  [geoname08 : 0 -&gt; 1][geoname09 : 0 -&gt; 1][geonametest01 : 0 -&gt; 1]
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// You can call the following function in the Elasticsearch API....
	Suggest *string `json:"suggest,omitempty" xml:"suggest,omitempty"`
	// The full name of the diagnostic item.
	//
	// example:
	//
	// ES_API
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetDesc() *string {
	return s.Desc
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetName() *string {
	return s.Name
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetResult() *string {
	return s.Result
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetSuggest() *string {
	return s.Suggest
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GetType() *string {
	return s.Type
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetDesc(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Desc = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetName(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Name = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetResult(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Result = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetSuggest(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Suggest = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetType(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Type = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) Validate() error {
	return dara.Validate(s)
}
