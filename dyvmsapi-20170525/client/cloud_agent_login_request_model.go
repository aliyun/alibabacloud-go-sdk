// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentLoginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindTel(v string) *CloudAgentLoginRequest
	GetBindTel() *string
	SetBindType(v int64) *CloudAgentLoginRequest
	GetBindType() *int64
	SetCno(v string) *CloudAgentLoginRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudAgentLoginRequest
	GetEnterpriseId() *int64
	SetLoginStatus(v int64) *CloudAgentLoginRequest
	GetLoginStatus() *int64
	SetPauseDescription(v string) *CloudAgentLoginRequest
	GetPauseDescription() *string
}

type CloudAgentLoginRequest struct {
	// 绑定电话
	//
	// This parameter is required.
	//
	// example:
	//
	// 41008502
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 取值 1.普通电话,2.分机,3.webrtc </br>说明：绑定类型如果是分机则必须先让分机电话设备注册成功。如果绑定类型为webrtc，就算调用接口成功也是无法呼叫的。如果在企业设置页面开启了自适应绑定电话类型，则系统会根据绑定的号码来找到对应的绑定类型，使用系统找到的绑定类型，此处的设置优先级低。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BindType *int64 `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 座席工号；取值3-10位正整数
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 登录状态，1：置闲，2：置忙，默认为1
	//
	// example:
	//
	// 1
	LoginStatus *int64 `json:"LoginStatus,omitempty" xml:"LoginStatus,omitempty"`
	// 置忙描述
	//
	// example:
	//
	// description
	PauseDescription *string `json:"PauseDescription,omitempty" xml:"PauseDescription,omitempty"`
}

func (s CloudAgentLoginRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentLoginRequest) GoString() string {
	return s.String()
}

func (s *CloudAgentLoginRequest) GetBindTel() *string {
	return s.BindTel
}

func (s *CloudAgentLoginRequest) GetBindType() *int64 {
	return s.BindType
}

func (s *CloudAgentLoginRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudAgentLoginRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudAgentLoginRequest) GetLoginStatus() *int64 {
	return s.LoginStatus
}

func (s *CloudAgentLoginRequest) GetPauseDescription() *string {
	return s.PauseDescription
}

func (s *CloudAgentLoginRequest) SetBindTel(v string) *CloudAgentLoginRequest {
	s.BindTel = &v
	return s
}

func (s *CloudAgentLoginRequest) SetBindType(v int64) *CloudAgentLoginRequest {
	s.BindType = &v
	return s
}

func (s *CloudAgentLoginRequest) SetCno(v string) *CloudAgentLoginRequest {
	s.Cno = &v
	return s
}

func (s *CloudAgentLoginRequest) SetEnterpriseId(v int64) *CloudAgentLoginRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudAgentLoginRequest) SetLoginStatus(v int64) *CloudAgentLoginRequest {
	s.LoginStatus = &v
	return s
}

func (s *CloudAgentLoginRequest) SetPauseDescription(v string) *CloudAgentLoginRequest {
	s.PauseDescription = &v
	return s
}

func (s *CloudAgentLoginRequest) Validate() error {
	return dara.Validate(s)
}
