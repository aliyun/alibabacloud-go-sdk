// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeOrderMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetChangeOrderMetricResponseBody
	GetCode() *string
	SetData(v []*GetChangeOrderMetricResponseBodyData) *GetChangeOrderMetricResponseBody
	GetData() []*GetChangeOrderMetricResponseBodyData
	SetMessage(v string) *GetChangeOrderMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChangeOrderMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetChangeOrderMetricResponseBody
	GetSuccess() *bool
}

type GetChangeOrderMetricResponseBody struct {
	// The HTTP status code. The following limits are imposed on the ID:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of applications.
	Data []*GetChangeOrderMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The additional information that is returned. The following limits are imposed on the ID:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3B763F98-0BA2-5C23-B6B8-558568D2C1C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the microservice list was obtained. The following limits are imposed on the ID:
	//
	// 	- **true**: The namespaces were obtained.
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetChangeOrderMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetChangeOrderMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChangeOrderMetricResponseBody) GetData() []*GetChangeOrderMetricResponseBodyData {
	return s.Data
}

func (s *GetChangeOrderMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChangeOrderMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChangeOrderMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetChangeOrderMetricResponseBody) SetCode(v string) *GetChangeOrderMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetChangeOrderMetricResponseBody) SetData(v []*GetChangeOrderMetricResponseBodyData) *GetChangeOrderMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetChangeOrderMetricResponseBody) SetMessage(v string) *GetChangeOrderMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetChangeOrderMetricResponseBody) SetRequestId(v string) *GetChangeOrderMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChangeOrderMetricResponseBody) SetSuccess(v bool) *GetChangeOrderMetricResponseBody {
	s.Success = &v
	return s
}

func (s *GetChangeOrderMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetChangeOrderMetricResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId         *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AvgTimeCostMs *float32 `json:"AvgTimeCostMs,omitempty" xml:"AvgTimeCostMs,omitempty"`
	// The number of abnormal change orders.
	//
	// example:
	//
	// 1
	Error *int64 `json:"Error,omitempty" xml:"Error,omitempty"`
	// The percentage of change failures.
	//
	// example:
	//
	// 0.25
	ErrorPercent  *float32 `json:"ErrorPercent,omitempty" xml:"ErrorPercent,omitempty"`
	MaxTimeCostMs *float32 `json:"MaxTimeCostMs,omitempty" xml:"MaxTimeCostMs,omitempty"`
	// The application name.
	//
	// example:
	//
	// test
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OptimizeSuggestions *string `json:"OptimizeSuggestions,omitempty" xml:"OptimizeSuggestions,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskTimeCostMsAvg *string `json:"TaskTimeCostMsAvg,omitempty" xml:"TaskTimeCostMsAvg,omitempty"`
	// The total number of change orders.
	//
	// example:
	//
	// 4
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetChangeOrderMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChangeOrderMetricResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetChangeOrderMetricResponseBodyData) GetAvgTimeCostMs() *float32 {
	return s.AvgTimeCostMs
}

func (s *GetChangeOrderMetricResponseBodyData) GetError() *int64 {
	return s.Error
}

func (s *GetChangeOrderMetricResponseBodyData) GetErrorPercent() *float32 {
	return s.ErrorPercent
}

func (s *GetChangeOrderMetricResponseBodyData) GetMaxTimeCostMs() *float32 {
	return s.MaxTimeCostMs
}

func (s *GetChangeOrderMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetChangeOrderMetricResponseBodyData) GetOptimizeSuggestions() *string {
	return s.OptimizeSuggestions
}

func (s *GetChangeOrderMetricResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetChangeOrderMetricResponseBodyData) GetTaskTimeCostMsAvg() *string {
	return s.TaskTimeCostMsAvg
}

func (s *GetChangeOrderMetricResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetChangeOrderMetricResponseBodyData) SetAppId(v string) *GetChangeOrderMetricResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetAvgTimeCostMs(v float32) *GetChangeOrderMetricResponseBodyData {
	s.AvgTimeCostMs = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetError(v int64) *GetChangeOrderMetricResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetErrorPercent(v float32) *GetChangeOrderMetricResponseBodyData {
	s.ErrorPercent = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetMaxTimeCostMs(v float32) *GetChangeOrderMetricResponseBodyData {
	s.MaxTimeCostMs = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetName(v string) *GetChangeOrderMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetOptimizeSuggestions(v string) *GetChangeOrderMetricResponseBodyData {
	s.OptimizeSuggestions = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetRegionId(v string) *GetChangeOrderMetricResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetTaskTimeCostMsAvg(v string) *GetChangeOrderMetricResponseBodyData {
	s.TaskTimeCostMsAvg = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) SetTotal(v int64) *GetChangeOrderMetricResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetChangeOrderMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
