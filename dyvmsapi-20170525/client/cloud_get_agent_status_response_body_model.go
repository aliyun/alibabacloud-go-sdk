// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetAgentStatusResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetAgentStatusResponseBody
	GetCode() *string
	SetData(v *CloudGetAgentStatusResponseBodyData) *CloudGetAgentStatusResponseBody
	GetData() *CloudGetAgentStatusResponseBodyData
	SetMessage(v string) *CloudGetAgentStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetAgentStatusResponseBody
	GetRequestId() *string
}

type CloudGetAgentStatusResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetAgentStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetAgentStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetAgentStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetAgentStatusResponseBody) GetData() *CloudGetAgentStatusResponseBodyData {
	return s.Data
}

func (s *CloudGetAgentStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetAgentStatusResponseBody) SetAccessDeniedDetail(v string) *CloudGetAgentStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetAgentStatusResponseBody) SetCode(v string) *CloudGetAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetAgentStatusResponseBody) SetData(v *CloudGetAgentStatusResponseBodyData) *CloudGetAgentStatusResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetAgentStatusResponseBody) SetMessage(v string) *CloudGetAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetAgentStatusResponseBody) SetRequestId(v string) *CloudGetAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetAgentStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetAgentStatusResponseBodyData struct {
	// 座席设备状态，座席登录状态不是离线时返回  -1：错误的状态  0：空闲状态  1：锁定状态   2：邀请中状态  3：响铃状态  4：使用中状态
	//
	// example:
	//
	// 1
	DeviceStatus *int64 `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// 座席登录状态   0：离线状态  1：在线状态  2：置忙状态  3：整理状态
	//
	// example:
	//
	// 1
	LoginStatus *int64 `json:"LoginStatus,omitempty" xml:"LoginStatus,omitempty"`
	// 座席当前参与的通话唯一标识，座席设备状态为1/2/3/4时返回
	//
	// example:
	//
	// medias_1-162046xxxx.12
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 座席状态描述（结合了登录状态和设备状态），离线，空闲，置忙，整理，呼叫中，响铃，保持，通话
	//
	// example:
	//
	// 示例值示例值
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// 状态对应的动作，座席登录状态不是离线时返回
	//
	// example:
	//
	// idle
	StateAction *string `json:"StateAction,omitempty" xml:"StateAction,omitempty"`
}

func (s CloudGetAgentStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetAgentStatusResponseBodyData) GetDeviceStatus() *int64 {
	return s.DeviceStatus
}

func (s *CloudGetAgentStatusResponseBodyData) GetLoginStatus() *int64 {
	return s.LoginStatus
}

func (s *CloudGetAgentStatusResponseBodyData) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *CloudGetAgentStatusResponseBodyData) GetState() *string {
	return s.State
}

func (s *CloudGetAgentStatusResponseBodyData) GetStateAction() *string {
	return s.StateAction
}

func (s *CloudGetAgentStatusResponseBodyData) SetDeviceStatus(v int64) *CloudGetAgentStatusResponseBodyData {
	s.DeviceStatus = &v
	return s
}

func (s *CloudGetAgentStatusResponseBodyData) SetLoginStatus(v int64) *CloudGetAgentStatusResponseBodyData {
	s.LoginStatus = &v
	return s
}

func (s *CloudGetAgentStatusResponseBodyData) SetMainUniqueId(v string) *CloudGetAgentStatusResponseBodyData {
	s.MainUniqueId = &v
	return s
}

func (s *CloudGetAgentStatusResponseBodyData) SetState(v string) *CloudGetAgentStatusResponseBodyData {
	s.State = &v
	return s
}

func (s *CloudGetAgentStatusResponseBodyData) SetStateAction(v string) *CloudGetAgentStatusResponseBodyData {
	s.StateAction = &v
	return s
}

func (s *CloudGetAgentStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
