// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataTopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeDataTopResponseBody
	GetCode() *string
	SetData(v *GetQueryOptimizeDataTopResponseBodyData) *GetQueryOptimizeDataTopResponseBody
	GetData() *GetQueryOptimizeDataTopResponseBodyData
	SetMessage(v string) *GetQueryOptimizeDataTopResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeDataTopResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeDataTopResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeDataTopResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *GetQueryOptimizeDataTopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeDataTopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTopResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeDataTopResponseBody) GetData() *GetQueryOptimizeDataTopResponseBodyData {
	return s.Data
}

func (s *GetQueryOptimizeDataTopResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeDataTopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeDataTopResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeDataTopResponseBody) SetCode(v string) *GetQueryOptimizeDataTopResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) SetData(v *GetQueryOptimizeDataTopResponseBodyData) *GetQueryOptimizeDataTopResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) SetMessage(v string) *GetQueryOptimizeDataTopResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) SetRequestId(v string) *GetQueryOptimizeDataTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) SetSuccess(v string) *GetQueryOptimizeDataTopResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQueryOptimizeDataTopResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The information about the instances.
	List []*GetQueryOptimizeDataTopResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeDataTopResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTopResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetQueryOptimizeDataTopResponseBodyData) GetList() []*GetQueryOptimizeDataTopResponseBodyDataList {
	return s.List
}

func (s *GetQueryOptimizeDataTopResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQueryOptimizeDataTopResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueryOptimizeDataTopResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetExtra(v string) *GetQueryOptimizeDataTopResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetList(v []*GetQueryOptimizeDataTopResponseBodyDataList) *GetQueryOptimizeDataTopResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeDataTopResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeDataTopResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetTotal(v int64) *GetQueryOptimizeDataTopResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetQueryOptimizeDataTopResponseBodyDataList struct {
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric name. Valid values:
	//
	// 	- **sqlExecuteCount**: the number of slow SQL executions.
	//
	// 	- **optimizedSqlExecuteCount**: the number of slow SQL executions that need to be optimized.
	//
	// example:
	//
	// sqlExecuteCount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 100
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetQueryOptimizeDataTopResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTopResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) GetType() *string {
	return s.Type
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) GetValue() *float64 {
	return s.Value
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) SetInstanceId(v string) *GetQueryOptimizeDataTopResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) SetType(v string) *GetQueryOptimizeDataTopResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) SetValue(v float64) *GetQueryOptimizeDataTopResponseBodyDataList {
	s.Value = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
