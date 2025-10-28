// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoResourceOptimizeRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DisableAutoResourceOptimizeRulesResponseBody
	GetCode() *int64
	SetData(v *DisableAutoResourceOptimizeRulesResponseBodyData) *DisableAutoResourceOptimizeRulesResponseBody
	GetData() *DisableAutoResourceOptimizeRulesResponseBodyData
	SetMessage(v string) *DisableAutoResourceOptimizeRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableAutoResourceOptimizeRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableAutoResourceOptimizeRulesResponseBody
	GetSuccess() *bool
}

type DisableAutoResourceOptimizeRulesResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DisableAutoResourceOptimizeRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DisableAutoResourceOptimizeRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoResourceOptimizeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) GetData() *DisableAutoResourceOptimizeRulesResponseBodyData {
	return s.Data
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) SetCode(v int64) *DisableAutoResourceOptimizeRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) SetData(v *DisableAutoResourceOptimizeRulesResponseBodyData) *DisableAutoResourceOptimizeRulesResponseBody {
	s.Data = v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) SetMessage(v string) *DisableAutoResourceOptimizeRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) SetRequestId(v string) *DisableAutoResourceOptimizeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) SetSuccess(v bool) *DisableAutoResourceOptimizeRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableAutoResourceOptimizeRulesResponseBodyData struct {
	// The number of database instances for which the automatic tablespace fragment recycling feature failed to be disabled.
	//
	// example:
	//
	// 1
	ConfigFailInstanceCount *int64 `json:"ConfigFailInstanceCount,omitempty" xml:"ConfigFailInstanceCount,omitempty"`
	// The list of database instances for which the automatic tablespace fragment recycling feature failed to be disabled.
	ConfigFailInstanceList []*DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList `json:"ConfigFailInstanceList,omitempty" xml:"ConfigFailInstanceList,omitempty" type:"Repeated"`
	// The number of database instances for which the automatic tablespace fragment recycling feature is disabled.
	//
	// example:
	//
	// 1
	ConfigSuccessInstanceCount *int64 `json:"ConfigSuccessInstanceCount,omitempty" xml:"ConfigSuccessInstanceCount,omitempty"`
	// The list of database instances for which the automatic tablespace fragment recycling feature is disabled.
	ConfigSuccessInstanceList []*DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList `json:"ConfigSuccessInstanceList,omitempty" xml:"ConfigSuccessInstanceList,omitempty" type:"Repeated"`
	// The total number of database instances.
	//
	// example:
	//
	// 2
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" xml:"TotalInstanceCount,omitempty"`
}

func (s DisableAutoResourceOptimizeRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoResourceOptimizeRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) GetConfigFailInstanceCount() *int64 {
	return s.ConfigFailInstanceCount
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) GetConfigFailInstanceList() []*DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList {
	return s.ConfigFailInstanceList
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) GetConfigSuccessInstanceCount() *int64 {
	return s.ConfigSuccessInstanceCount
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) GetConfigSuccessInstanceList() []*DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList {
	return s.ConfigSuccessInstanceList
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) GetTotalInstanceCount() *int64 {
	return s.TotalInstanceCount
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) SetConfigFailInstanceCount(v int64) *DisableAutoResourceOptimizeRulesResponseBodyData {
	s.ConfigFailInstanceCount = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) SetConfigFailInstanceList(v []*DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) *DisableAutoResourceOptimizeRulesResponseBodyData {
	s.ConfigFailInstanceList = v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) SetConfigSuccessInstanceCount(v int64) *DisableAutoResourceOptimizeRulesResponseBodyData {
	s.ConfigSuccessInstanceCount = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) SetConfigSuccessInstanceList(v []*DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList) *DisableAutoResourceOptimizeRulesResponseBodyData {
	s.ConfigSuccessInstanceList = v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) SetTotalInstanceCount(v int64) *DisableAutoResourceOptimizeRulesResponseBodyData {
	s.TotalInstanceCount = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyData) Validate() error {
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

type DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList struct {
	// Indicates whether the automatic tablespace fragment recycling feature is disabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ConfigSuccess *bool `json:"ConfigSuccess,omitempty" xml:"ConfigSuccess,omitempty"`
	// The error message returned if the request failed.
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

func (s DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) GoString() string {
	return s.String()
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) GetConfigSuccess() *bool {
	return s.ConfigSuccess
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) SetConfigSuccess(v bool) *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList {
	s.ConfigSuccess = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) SetErrorMessage(v string) *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList {
	s.ErrorMessage = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) SetInstanceId(v string) *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigFailInstanceList) Validate() error {
	return dara.Validate(s)
}

type DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList struct {
	// Indicates whether the automatic tablespace fragment recycling feature is disabled. Valid values:
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

func (s DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList) GoString() string {
	return s.String()
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList) GetConfigSuccess() *bool {
	return s.ConfigSuccess
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList) SetConfigSuccess(v bool) *DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList {
	s.ConfigSuccess = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList) SetInstanceId(v string) *DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponseBodyDataConfigSuccessInstanceList) Validate() error {
	return dara.Validate(s)
}
