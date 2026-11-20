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
	// The business user identity. It is provided based on the identity configuration bound to the gateway credentials.
	Communication *OpenServiceRequestCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// Specifies whether to actually activate the benefits. A value of true initiates the activation process. A value of false only verifies the call and returns TEST_OK.
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
	// The object that carries the customer business user identifier.
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
	// The customer business user ID.
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
