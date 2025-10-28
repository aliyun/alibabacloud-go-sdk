// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoResourceOptimizeRulesAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateAutoResourceOptimizeRulesAsyncResponseBody
	GetCode() *int64
	SetData(v *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) *UpdateAutoResourceOptimizeRulesAsyncResponseBody
	GetData() *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData
	SetMessage(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAutoResourceOptimizeRulesAsyncResponseBody
	GetSuccess() *bool
}

type UpdateAutoResourceOptimizeRulesAsyncResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 9CB97BC4-6479-55D0-B9D0-EA925AFE****
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

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) GetData() *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData {
	return s.Data
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) SetCode(v int64) *UpdateAutoResourceOptimizeRulesAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) SetData(v *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) *UpdateAutoResourceOptimizeRulesAsyncResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) SetMessage(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) SetRequestId(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) SetSuccess(v bool) *UpdateAutoResourceOptimizeRulesAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAutoResourceOptimizeRulesAsyncResponseBodyData struct {
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
	ConfigResponse *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse `json:"ConfigResponse,omitempty" xml:"ConfigResponse,omitempty" type:"Struct"`
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
	// async__20ee808e72257f16a4fe024057ca****
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

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) GetComplete() *bool {
	return s.Complete
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) GetConfigResponse() *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse {
	return s.ConfigResponse
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) GetFail() *bool {
	return s.Fail
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) GetState() *string {
	return s.State
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) SetComplete(v bool) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData {
	s.Complete = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) SetConfigResponse(v *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData {
	s.ConfigResponse = v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) SetFail(v bool) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData {
	s.Fail = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) SetIsFinish(v bool) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) SetResultId(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) SetState(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData {
	s.State = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) SetTimestamp(v int64) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyData) Validate() error {
	if s.ConfigResponse != nil {
		if err := s.ConfigResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse struct {
	// The number of database instances for which the parameters failed to be configured.
	//
	// example:
	//
	// 1
	ConfigFailInstanceCount *int64 `json:"ConfigFailInstanceCount,omitempty" xml:"ConfigFailInstanceCount,omitempty"`
	// The database instances for which the parameters failed to be configured.
	ConfigFailInstanceList []*UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList `json:"ConfigFailInstanceList,omitempty" xml:"ConfigFailInstanceList,omitempty" type:"Repeated"`
	// The number of database instances for which the parameters are configured.
	//
	// example:
	//
	// 1
	ConfigSuccessInstanceCount *int64 `json:"ConfigSuccessInstanceCount,omitempty" xml:"ConfigSuccessInstanceCount,omitempty"`
	// The database instances for which the parameters are configured.
	ConfigSuccessInstanceList []*UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList `json:"ConfigSuccessInstanceList,omitempty" xml:"ConfigSuccessInstanceList,omitempty" type:"Repeated"`
	// The total number of database instances.
	//
	// example:
	//
	// 2
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" xml:"TotalInstanceCount,omitempty"`
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) GetConfigFailInstanceCount() *int64 {
	return s.ConfigFailInstanceCount
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) GetConfigFailInstanceList() []*UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList {
	return s.ConfigFailInstanceList
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) GetConfigSuccessInstanceCount() *int64 {
	return s.ConfigSuccessInstanceCount
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) GetConfigSuccessInstanceList() []*UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList {
	return s.ConfigSuccessInstanceList
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) GetTotalInstanceCount() *int64 {
	return s.TotalInstanceCount
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) SetConfigFailInstanceCount(v int64) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse {
	s.ConfigFailInstanceCount = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) SetConfigFailInstanceList(v []*UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse {
	s.ConfigFailInstanceList = v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) SetConfigSuccessInstanceCount(v int64) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse {
	s.ConfigSuccessInstanceCount = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) SetConfigSuccessInstanceList(v []*UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse {
	s.ConfigSuccessInstanceList = v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) SetTotalInstanceCount(v int64) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse {
	s.TotalInstanceCount = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponse) Validate() error {
	if s.ConfigFailInstanceList != nil {
		for _, item := range s.ConfigFailInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConfigSuccessInstanceList != nil {
		for _, item := range s.ConfigSuccessInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList struct {
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
	// Only Support DAS Pro High-availability Edition RDS MySQL 5.6, 5.7, 8.0 instance, and CPU cores >= 4, innodb_file_per_table=ON
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze9xrhze0709****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) GetConfigSuccess() *bool {
	return s.ConfigSuccess
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) SetConfigSuccess(v bool) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList {
	s.ConfigSuccess = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) SetErrorMessage(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) SetInstanceId(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList {
	s.InstanceId = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigFailInstanceList) Validate() error {
	return dara.Validate(s)
}

type UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList struct {
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

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) GetConfigSuccess() *bool {
	return s.ConfigSuccess
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) SetConfigSuccess(v bool) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList {
	s.ConfigSuccess = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) SetInstanceId(v string) *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList {
	s.InstanceId = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponseBodyDataConfigResponseConfigSuccessInstanceList) Validate() error {
	return dara.Validate(s)
}
