// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkAgentStatusDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkAgentStatusDetailResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkAgentStatusDetailResponseBody
	GetCode() *string
	SetData(v *ClinkAgentStatusDetailResponseBodyData) *ClinkAgentStatusDetailResponseBody
	GetData() *ClinkAgentStatusDetailResponseBodyData
	SetMessage(v string) *ClinkAgentStatusDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkAgentStatusDetailResponseBody
	GetRequestId() *string
}

type ClinkAgentStatusDetailResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkAgentStatusDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9927CD06-C33A-50CC-A9BB-A3427967A66F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkAgentStatusDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkAgentStatusDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkAgentStatusDetailResponseBody) GetData() *ClinkAgentStatusDetailResponseBodyData {
	return s.Data
}

func (s *ClinkAgentStatusDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkAgentStatusDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkAgentStatusDetailResponseBody) SetAccessDeniedDetail(v string) *ClinkAgentStatusDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBody) SetCode(v string) *ClinkAgentStatusDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBody) SetData(v *ClinkAgentStatusDetailResponseBodyData) *ClinkAgentStatusDetailResponseBody {
	s.Data = v
	return s
}

func (s *ClinkAgentStatusDetailResponseBody) SetMessage(v string) *ClinkAgentStatusDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBody) SetRequestId(v string) *ClinkAgentStatusDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkAgentStatusDetailResponseBodyData struct {
	// 座席状态详情
	AgentStatusDetail *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail `json:"AgentStatusDetail,omitempty" xml:"AgentStatusDetail,omitempty" type:"Struct"`
	// 请求 id
	//
	// example:
	//
	// 示例值示例值示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkAgentStatusDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusDetailResponseBodyData) GetAgentStatusDetail() *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	return s.AgentStatusDetail
}

func (s *ClinkAgentStatusDetailResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkAgentStatusDetailResponseBodyData) SetAgentStatusDetail(v *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) *ClinkAgentStatusDetailResponseBodyData {
	s.AgentStatusDetail = v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyData) SetClinkRequestId(v string) *ClinkAgentStatusDetailResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyData) Validate() error {
	if s.AgentStatusDetail != nil {
		if err := s.AgentStatusDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail struct {
	// 绑定号码
	//
	// example:
	//
	// 示例值示例值
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 电话类型，1:电话；2:IP话机；3:软电话
	//
	// example:
	//
	// 1
	BindType *int64 `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 座席号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 企业id
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 登录终端，1:工具条；2:座席；3:管理员；4:接口；6:APP；9:移动SDK
	//
	// example:
	//
	// 2
	LoginType *int64 `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	// 座席名称
	//
	// example:
	//
	// name3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 座席状态， IDLE：空闲； PAUSE：置忙； WRAPUP：整理； CALLING： 呼叫中； RINGING：响铃； BUSY：通话
	//
	// example:
	//
	// IDLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 座席状态详情，空闲，忙碌，整理，呼叫中，响铃，呼入振铃，外呼振铃，通话，呼入通话，外呼通话，自定义置忙状态
	//
	// example:
	//
	// 示例值示例值示例值
	StatusDetails *string `json:"StatusDetails,omitempty" xml:"StatusDetails,omitempty"`
}

func (s ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GetBindTel() *string {
	return s.BindTel
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GetBindType() *int64 {
	return s.BindType
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GetCno() *string {
	return s.Cno
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GetLoginType() *int64 {
	return s.LoginType
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GetName() *string {
	return s.Name
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GetStatus() *string {
	return s.Status
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) GetStatusDetails() *string {
	return s.StatusDetails
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) SetBindTel(v string) *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	s.BindTel = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) SetBindType(v int64) *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	s.BindType = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) SetCno(v string) *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	s.Cno = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) SetEnterpriseId(v int64) *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) SetLoginType(v int64) *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	s.LoginType = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) SetName(v string) *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	s.Name = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) SetStatus(v string) *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	s.Status = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) SetStatusDetails(v string) *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail {
	s.StatusDetails = &v
	return s
}

func (s *ClinkAgentStatusDetailResponseBodyDataAgentStatusDetail) Validate() error {
	return dara.Validate(s)
}
