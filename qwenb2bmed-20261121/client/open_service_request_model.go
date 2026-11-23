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
	// 按接入时约定填写业务用户信息。如果当前业务凭据无需提供用户信息，可省略整个对象。字段名为 user_info 和 user_id。
	Communication *OpenServiceRequestCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// 必填 JSON 布尔值。false 只验证本次调用，成功返回 TEST_OK；true 提交服务开通申请。
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
	// 业务用户信息对象。提供时必须填写其中的 user_id。
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
	// 非空字符串，填写你自己的应用中用于识别该用户的 ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// user-001
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
