// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommunication(v *OpenServiceRequestCommunication) *OpenServiceRequest
	GetCommunication() *OpenServiceRequestCommunication
	SetOpenBenefits(v bool) *OpenServiceRequest
	GetOpenBenefits() *bool
}

type OpenServiceRequest struct {
	// 按凭据的业务身份配置提供；固定用户身份或不要求正文用户身份时可省略。若提供，必须是只含 user_info 的对象，不能为 null 或空对象。测试与真实开通请求执行相同身份检查。
	Communication *OpenServiceRequestCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// 必填 JSON 布尔值，无默认值。false 仅验证接口、鉴权、身份和请求参数，成功返回 TEST_OK，不实际开通。true 发起服务开通；返回 OPENING 时仍在处理中。重复发送 true 仍是开通操作，不是只读状态查询。字符串 true/false 不合法。
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	OpenBenefits *bool `json:"openBenefits,omitempty" xml:"openBenefits,omitempty"`
}

func (s OpenServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenServiceRequest) GetCommunication() *OpenServiceRequestCommunication {
	return s.Communication
}

func (s *OpenServiceRequest) GetOpenBenefits() *bool {
	return s.OpenBenefits
}

func (s *OpenServiceRequest) SetCommunication(v *OpenServiceRequestCommunication) *OpenServiceRequest {
	s.Communication = v
	return s
}

func (s *OpenServiceRequest) SetOpenBenefits(v bool) *OpenServiceRequest {
	s.OpenBenefits = &v
	return s
}

func (s *OpenServiceRequest) Validate() error {
	if s.Communication != nil {
		if err := s.Communication.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenServiceRequestCommunication struct {
	// communication 存在时必填的对象，只包含 user_id，不能为 null 或空对象。
	//
	// This parameter is required.
	UserInfo *OpenServiceRequestCommunicationUserInfo `json:"user_info,omitempty" xml:"user_info,omitempty" type:"Struct"`
}

func (s OpenServiceRequestCommunication) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceRequestCommunication) GoString() string {
	return s.String()
}

func (s *OpenServiceRequestCommunication) GetUserInfo() *OpenServiceRequestCommunicationUserInfo {
	return s.UserInfo
}

func (s *OpenServiceRequestCommunication) SetUserInfo(v *OpenServiceRequestCommunicationUserInfo) *OpenServiceRequestCommunication {
	s.UserInfo = v
	return s
}

func (s *OpenServiceRequestCommunication) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenServiceRequestCommunicationUserInfo struct {
	// 客户应用内的业务用户标识，须符合凭据身份配置。非空，最多 256 个 UTF-16 代码单元，无首尾空白、控制字符或非法 Unicode。不是阿里云 UID、订单号或 SP 实例 ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// example-user-001
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s OpenServiceRequestCommunicationUserInfo) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceRequestCommunicationUserInfo) GoString() string {
	return s.String()
}

func (s *OpenServiceRequestCommunicationUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *OpenServiceRequestCommunicationUserInfo) SetUserId(v string) *OpenServiceRequestCommunicationUserInfo {
	s.UserId = &v
	return s
}

func (s *OpenServiceRequestCommunicationUserInfo) Validate() error {
	return dara.Validate(s)
}
