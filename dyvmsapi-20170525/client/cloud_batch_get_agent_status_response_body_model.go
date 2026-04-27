// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchGetAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudBatchGetAgentStatusResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudBatchGetAgentStatusResponseBody
	GetCode() *string
	SetData(v *CloudBatchGetAgentStatusResponseBodyData) *CloudBatchGetAgentStatusResponseBody
	GetData() *CloudBatchGetAgentStatusResponseBodyData
	SetMessage(v string) *CloudBatchGetAgentStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudBatchGetAgentStatusResponseBody
	GetRequestId() *string
}

type CloudBatchGetAgentStatusResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudBatchGetAgentStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudBatchGetAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchGetAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CloudBatchGetAgentStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudBatchGetAgentStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudBatchGetAgentStatusResponseBody) GetData() *CloudBatchGetAgentStatusResponseBodyData {
	return s.Data
}

func (s *CloudBatchGetAgentStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudBatchGetAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudBatchGetAgentStatusResponseBody) SetAccessDeniedDetail(v string) *CloudBatchGetAgentStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBody) SetCode(v string) *CloudBatchGetAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBody) SetData(v *CloudBatchGetAgentStatusResponseBodyData) *CloudBatchGetAgentStatusResponseBody {
	s.Data = v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBody) SetMessage(v string) *CloudBatchGetAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBody) SetRequestId(v string) *CloudBatchGetAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudBatchGetAgentStatusResponseBodyData struct {
	// 座席设备状态，座席登录状态不是离线时返回  -1：错误的状态  0：空闲状态  1：锁定状态   2：邀请中状态  3：响铃状态  4：使用中状态
	//
	// example:
	//
	// 0
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// 座席登录状态   0：离线状态  1：在线状态  2：置忙状态  3：整理状态
	//
	// example:
	//
	// 1
	LoginStatus *string `json:"LoginStatus,omitempty" xml:"LoginStatus,omitempty"`
	// 座席当前参与的通话唯一标识，座席设备状态为1/2/3/4时返回
	//
	// example:
	//
	// test3321.sip33
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 座席状态描述（结合了登录状态和设备状态），离线，空闲，置忙，整理，呼叫中，响铃，保持，通话
	//
	// example:
	//
	// 示例值
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// 状态对应的动作，座席登录状态不是离线时返回。详见[座席状态动作说明](../../工具条/座席状态动作.md)
	//
	// example:
	//
	// Idle
	StateAction *string `json:"StateAction,omitempty" xml:"StateAction,omitempty"`
}

func (s CloudBatchGetAgentStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchGetAgentStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudBatchGetAgentStatusResponseBodyData) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *CloudBatchGetAgentStatusResponseBodyData) GetLoginStatus() *string {
	return s.LoginStatus
}

func (s *CloudBatchGetAgentStatusResponseBodyData) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *CloudBatchGetAgentStatusResponseBodyData) GetState() *string {
	return s.State
}

func (s *CloudBatchGetAgentStatusResponseBodyData) GetStateAction() *string {
	return s.StateAction
}

func (s *CloudBatchGetAgentStatusResponseBodyData) SetDeviceStatus(v string) *CloudBatchGetAgentStatusResponseBodyData {
	s.DeviceStatus = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBodyData) SetLoginStatus(v string) *CloudBatchGetAgentStatusResponseBodyData {
	s.LoginStatus = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBodyData) SetMainUniqueId(v string) *CloudBatchGetAgentStatusResponseBodyData {
	s.MainUniqueId = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBodyData) SetState(v string) *CloudBatchGetAgentStatusResponseBodyData {
	s.State = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBodyData) SetStateAction(v string) *CloudBatchGetAgentStatusResponseBodyData {
	s.StateAction = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
