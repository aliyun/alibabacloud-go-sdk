// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoThrottleRulesAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateAutoThrottleRulesAsyncResponseBody
	GetCode() *int64
	SetData(v *UpdateAutoThrottleRulesAsyncResponseBodyData) *UpdateAutoThrottleRulesAsyncResponseBody
	GetData() *UpdateAutoThrottleRulesAsyncResponseBodyData
	SetMessage(v string) *UpdateAutoThrottleRulesAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAutoThrottleRulesAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAutoThrottleRulesAsyncResponseBody
	GetSuccess() *bool
}

type UpdateAutoThrottleRulesAsyncResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *UpdateAutoThrottleRulesAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAutoThrottleRulesAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoThrottleRulesAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) GetData() *UpdateAutoThrottleRulesAsyncResponseBodyData {
	return s.Data
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) SetCode(v int64) *UpdateAutoThrottleRulesAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) SetData(v *UpdateAutoThrottleRulesAsyncResponseBodyData) *UpdateAutoThrottleRulesAsyncResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) SetMessage(v string) *UpdateAutoThrottleRulesAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) SetRequestId(v string) *UpdateAutoThrottleRulesAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) SetSuccess(v bool) *UpdateAutoThrottleRulesAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateAutoThrottleRulesAsyncResponseBodyData struct {
	// Indicates whether the asynchronous request was complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Complete *bool `json:"Complete,omitempty" xml:"Complete,omitempty"`
	// The returned data of the configuration.
	//
	// >  The data is returned only if the value of isFinish is **true**. This value indicates that the asynchronous request is complete.
	ConfigResponse *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse `json:"ConfigResponse,omitempty" xml:"ConfigResponse,omitempty" type:"Struct"`
	// Indicates whether the asynchronous request failed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Fail *bool `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// Indicates whether the asynchronous request was complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsFinish *bool `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// The ID of the asynchronous request.
	//
	// example:
	//
	// async__665ee69612f1627c7fd9f3c85075****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	// The state of the asynchronous request. Valid values:
	//
	// 	- **RUNNING**
	//
	// 	- **SUCCESS**
	//
	// 	- **FAIL**
	//
	// example:
	//
	// SUCCESS
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the asynchronous request was made. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1645668213000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s UpdateAutoThrottleRulesAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoThrottleRulesAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) GetComplete() *bool {
	return s.Complete
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) GetConfigResponse() *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse {
	return s.ConfigResponse
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) GetFail() *bool {
	return s.Fail
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) GetState() *string {
	return s.State
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) SetComplete(v bool) *UpdateAutoThrottleRulesAsyncResponseBodyData {
	s.Complete = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) SetConfigResponse(v *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) *UpdateAutoThrottleRulesAsyncResponseBodyData {
	s.ConfigResponse = v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) SetFail(v bool) *UpdateAutoThrottleRulesAsyncResponseBodyData {
	s.Fail = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) SetIsFinish(v bool) *UpdateAutoThrottleRulesAsyncResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) SetResultId(v string) *UpdateAutoThrottleRulesAsyncResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) SetState(v string) *UpdateAutoThrottleRulesAsyncResponseBodyData {
	s.State = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) SetTimestamp(v int64) *UpdateAutoThrottleRulesAsyncResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse struct {
	// The number of database instances for which the parameters failed to be configured.
	//
	// example:
	//
	// 1
	ConfigFailInstanceCount *int64 `json:"ConfigFailInstanceCount,omitempty" xml:"ConfigFailInstanceCount,omitempty"`
	// The database instances for which the parameters failed to be configured.
	ConfigFailInstanceList []*UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList `json:"ConfigFailInstanceList,omitempty" xml:"ConfigFailInstanceList,omitempty" type:"Repeated"`
	// The number of database instances for which the parameters are configured.
	//
	// example:
	//
	// 1
	ConfigSuccessInstanceCount *int64 `json:"ConfigSuccessInstanceCount,omitempty" xml:"ConfigSuccessInstanceCount,omitempty"`
	// The database instances for which the parameters are configured.
	ConfigSuccessInstanceList []*UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList `json:"ConfigSuccessInstanceList,omitempty" xml:"ConfigSuccessInstanceList,omitempty" type:"Repeated"`
	// The total number of database instances.
	//
	// example:
	//
	// 2
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" xml:"TotalInstanceCount,omitempty"`
}

func (s UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) GetConfigFailInstanceCount() *int64 {
	return s.ConfigFailInstanceCount
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) GetConfigFailInstanceList() []*UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList {
	return s.ConfigFailInstanceList
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) GetConfigSuccessInstanceCount() *int64 {
	return s.ConfigSuccessInstanceCount
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) GetConfigSuccessInstanceList() []*UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList {
	return s.ConfigSuccessInstanceList
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) GetTotalInstanceCount() *int64 {
	return s.TotalInstanceCount
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) SetConfigFailInstanceCount(v int64) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse {
	s.ConfigFailInstanceCount = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) SetConfigFailInstanceList(v []*UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse {
	s.ConfigFailInstanceList = v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) SetConfigSuccessInstanceCount(v int64) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse {
	s.ConfigSuccessInstanceCount = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) SetConfigSuccessInstanceList(v []*UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse {
	s.ConfigSuccessInstanceList = v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) SetTotalInstanceCount(v int64) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse {
	s.TotalInstanceCount = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponse) Validate() error {
	return dara.Validate(s)
}

type UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList struct {
	// Indicates whether the parameters are configured. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ConfigSuccess *bool `json:"ConfigSuccess,omitempty" xml:"ConfigSuccess,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// instance das autonomy service is off or can not find instance
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze9xrhze0709****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) GoString() string {
	return s.String()
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) GetConfigSuccess() *bool {
	return s.ConfigSuccess
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) SetConfigSuccess(v bool) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList {
	s.ConfigSuccess = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) SetErrorMessage(v string) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) SetInstanceId(v string) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList {
	s.InstanceId = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) Validate() error {
	return dara.Validate(s)
}

type UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList struct {
	// Indicates whether the parameters are configured. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ConfigSuccess *bool `json:"ConfigSuccess,omitempty" xml:"ConfigSuccess,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) GoString() string {
	return s.String()
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) GetConfigSuccess() *bool {
	return s.ConfigSuccess
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) SetConfigSuccess(v bool) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList {
	s.ConfigSuccess = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) SetInstanceId(v string) *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList {
	s.InstanceId = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) Validate() error {
	return dara.Validate(s)
}
