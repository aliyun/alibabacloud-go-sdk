// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListOnlineAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListOnlineAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListOnlineAgentResponseBody
	GetCode() *string
	SetData(v *CloudListOnlineAgentResponseBodyData) *CloudListOnlineAgentResponseBody
	GetData() *CloudListOnlineAgentResponseBodyData
	SetMessage(v string) *CloudListOnlineAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListOnlineAgentResponseBody
	GetRequestId() *string
}

type CloudListOnlineAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListOnlineAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListOnlineAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListOnlineAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListOnlineAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListOnlineAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListOnlineAgentResponseBody) GetData() *CloudListOnlineAgentResponseBodyData {
	return s.Data
}

func (s *CloudListOnlineAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListOnlineAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListOnlineAgentResponseBody) SetAccessDeniedDetail(v string) *CloudListOnlineAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListOnlineAgentResponseBody) SetCode(v string) *CloudListOnlineAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListOnlineAgentResponseBody) SetData(v *CloudListOnlineAgentResponseBodyData) *CloudListOnlineAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudListOnlineAgentResponseBody) SetMessage(v string) *CloudListOnlineAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListOnlineAgentResponseBody) SetRequestId(v string) *CloudListOnlineAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListOnlineAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListOnlineAgentResponseBodyData struct {
	// 座席状态数组
	AgentStatuses []*CloudListOnlineAgentResponseBodyDataAgentStatuses `json:"AgentStatuses,omitempty" xml:"AgentStatuses,omitempty" type:"Repeated"`
}

func (s CloudListOnlineAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListOnlineAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListOnlineAgentResponseBodyData) GetAgentStatuses() []*CloudListOnlineAgentResponseBodyDataAgentStatuses {
	return s.AgentStatuses
}

func (s *CloudListOnlineAgentResponseBodyData) SetAgentStatuses(v []*CloudListOnlineAgentResponseBodyDataAgentStatuses) *CloudListOnlineAgentResponseBodyData {
	s.AgentStatuses = v
	return s
}

func (s *CloudListOnlineAgentResponseBodyData) Validate() error {
	if s.AgentStatuses != nil {
		for _, item := range s.AgentStatuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudListOnlineAgentResponseBodyDataAgentStatuses struct {
	// 绑定电话
	//
	// example:
	//
	// 22223333
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 设备状态，-1 失效，0 空闲，1 已锁定，2 呼叫中，3 响铃，4 通话中
	//
	// example:
	//
	// 0
	DeviceStatus *int64 `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// 设备状态时长
	//
	// example:
	//
	// 0
	DeviceStatusDuration *int64 `json:"DeviceStatusDuration,omitempty" xml:"DeviceStatusDuration,omitempty"`
	// 登录状态，0离线，1在线，2置忙，3整理
	//
	// example:
	//
	// 1
	LoginStatus *int64 `json:"LoginStatus,omitempty" xml:"LoginStatus,omitempty"`
	// 登录状态时长
	//
	// example:
	//
	// 7
	LoginStatusDuration *int64 `json:"LoginStatusDuration,omitempty" xml:"LoginStatusDuration,omitempty"`
	// 登录时间
	//
	// example:
	//
	// 1774821736
	LoginTime *string `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	// 座席名称
	//
	// example:
	//
	// 示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 座席状态（结合了登录状态和设备状态），失效，空闲，置忙，整理，呼叫中，响铃，通话
	//
	// example:
	//
	// 示例值示例值示例值
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CloudListOnlineAgentResponseBodyDataAgentStatuses) String() string {
	return dara.Prettify(s)
}

func (s CloudListOnlineAgentResponseBodyDataAgentStatuses) GoString() string {
	return s.String()
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetBindTel() *string {
	return s.BindTel
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetCno() *string {
	return s.Cno
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetDeviceStatus() *int64 {
	return s.DeviceStatus
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetDeviceStatusDuration() *int64 {
	return s.DeviceStatusDuration
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetLoginStatus() *int64 {
	return s.LoginStatus
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetLoginStatusDuration() *int64 {
	return s.LoginStatusDuration
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetLoginTime() *string {
	return s.LoginTime
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetName() *string {
	return s.Name
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) GetState() *string {
	return s.State
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetBindTel(v string) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.BindTel = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetCno(v string) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.Cno = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetDeviceStatus(v int64) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.DeviceStatus = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetDeviceStatusDuration(v int64) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.DeviceStatusDuration = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetLoginStatus(v int64) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.LoginStatus = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetLoginStatusDuration(v int64) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.LoginStatusDuration = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetLoginTime(v string) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.LoginTime = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetName(v string) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.Name = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) SetState(v string) *CloudListOnlineAgentResponseBodyDataAgentStatuses {
	s.State = &v
	return s
}

func (s *CloudListOnlineAgentResponseBodyDataAgentStatuses) Validate() error {
	return dara.Validate(s)
}
