// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoThrottleRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DisableAutoThrottleRulesResponseBody
	GetCode() *int64
	SetData(v *DisableAutoThrottleRulesResponseBodyData) *DisableAutoThrottleRulesResponseBody
	GetData() *DisableAutoThrottleRulesResponseBodyData
	SetMessage(v string) *DisableAutoThrottleRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableAutoThrottleRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableAutoThrottleRulesResponseBody
	GetSuccess() *bool
}

type DisableAutoThrottleRulesResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DisableAutoThrottleRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DisableAutoThrottleRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoThrottleRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAutoThrottleRulesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DisableAutoThrottleRulesResponseBody) GetData() *DisableAutoThrottleRulesResponseBodyData {
	return s.Data
}

func (s *DisableAutoThrottleRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableAutoThrottleRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAutoThrottleRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableAutoThrottleRulesResponseBody) SetCode(v int64) *DisableAutoThrottleRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBody) SetData(v *DisableAutoThrottleRulesResponseBodyData) *DisableAutoThrottleRulesResponseBody {
	s.Data = v
	return s
}

func (s *DisableAutoThrottleRulesResponseBody) SetMessage(v string) *DisableAutoThrottleRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBody) SetRequestId(v string) *DisableAutoThrottleRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBody) SetSuccess(v bool) *DisableAutoThrottleRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DisableAutoThrottleRulesResponseBodyData struct {
	// The number of database instances for which the automatic SQL throttling feature failed to be disabled.
	//
	// example:
	//
	// 1
	ConfigFailInstanceCount *int64 `json:"ConfigFailInstanceCount,omitempty" xml:"ConfigFailInstanceCount,omitempty"`
	// The database instances for which the automatic SQL throttling feature failed to be disabled.
	ConfigFailInstanceList []*DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList `json:"ConfigFailInstanceList,omitempty" xml:"ConfigFailInstanceList,omitempty" type:"Repeated"`
	// The number of database instances for which the automatic SQL throttling feature is disabled.
	//
	// example:
	//
	// 1
	ConfigSuccessInstanceCount *int64 `json:"ConfigSuccessInstanceCount,omitempty" xml:"ConfigSuccessInstanceCount,omitempty"`
	// The database instances for which the automatic SQL throttling feature is disabled.
	ConfigSuccessInstanceList []*DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList `json:"ConfigSuccessInstanceList,omitempty" xml:"ConfigSuccessInstanceList,omitempty" type:"Repeated"`
	// The total number of database instances.
	//
	// example:
	//
	// 2
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" xml:"TotalInstanceCount,omitempty"`
}

func (s DisableAutoThrottleRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoThrottleRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableAutoThrottleRulesResponseBodyData) GetConfigFailInstanceCount() *int64 {
	return s.ConfigFailInstanceCount
}

func (s *DisableAutoThrottleRulesResponseBodyData) GetConfigFailInstanceList() []*DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList {
	return s.ConfigFailInstanceList
}

func (s *DisableAutoThrottleRulesResponseBodyData) GetConfigSuccessInstanceCount() *int64 {
	return s.ConfigSuccessInstanceCount
}

func (s *DisableAutoThrottleRulesResponseBodyData) GetConfigSuccessInstanceList() []*DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList {
	return s.ConfigSuccessInstanceList
}

func (s *DisableAutoThrottleRulesResponseBodyData) GetTotalInstanceCount() *int64 {
	return s.TotalInstanceCount
}

func (s *DisableAutoThrottleRulesResponseBodyData) SetConfigFailInstanceCount(v int64) *DisableAutoThrottleRulesResponseBodyData {
	s.ConfigFailInstanceCount = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyData) SetConfigFailInstanceList(v []*DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) *DisableAutoThrottleRulesResponseBodyData {
	s.ConfigFailInstanceList = v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyData) SetConfigSuccessInstanceCount(v int64) *DisableAutoThrottleRulesResponseBodyData {
	s.ConfigSuccessInstanceCount = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyData) SetConfigSuccessInstanceList(v []*DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList) *DisableAutoThrottleRulesResponseBodyData {
	s.ConfigSuccessInstanceList = v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyData) SetTotalInstanceCount(v int64) *DisableAutoThrottleRulesResponseBodyData {
	s.TotalInstanceCount = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList struct {
	// Indicates whether the automatic SQL throttling feature is disabled. Valid values:
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
	// cannot found instance by rm-2ze9xrhze0709****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze9xrhze0709****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) GoString() string {
	return s.String()
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) GetConfigSuccess() *bool {
	return s.ConfigSuccess
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) SetConfigSuccess(v bool) *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList {
	s.ConfigSuccess = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) SetErrorMessage(v string) *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList {
	s.ErrorMessage = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) SetInstanceId(v string) *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigFailInstanceList) Validate() error {
	return dara.Validate(s)
}

type DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList struct {
	// Indicates whether the automatic SQL throttling feature is disabled. Valid values:
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

func (s DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList) GoString() string {
	return s.String()
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList) GetConfigSuccess() *bool {
	return s.ConfigSuccess
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList) SetConfigSuccess(v bool) *DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList {
	s.ConfigSuccess = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList) SetInstanceId(v string) *DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DisableAutoThrottleRulesResponseBodyDataConfigSuccessInstanceList) Validate() error {
	return dara.Validate(s)
}
