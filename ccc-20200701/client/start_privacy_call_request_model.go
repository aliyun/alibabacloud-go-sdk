// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPrivacyCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartPrivacyCallRequest
	GetAppId() *string
	SetCallee(v string) *StartPrivacyCallRequest
	GetCallee() *string
	SetCaller(v string) *StartPrivacyCallRequest
	GetCaller() *string
	SetInstanceId(v string) *StartPrivacyCallRequest
	GetInstanceId() *string
}

type StartPrivacyCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ALIPUBB7A727C170949
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1888888****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0109810****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartPrivacyCallRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPrivacyCallRequest) GoString() string {
	return s.String()
}

func (s *StartPrivacyCallRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartPrivacyCallRequest) GetCallee() *string {
	return s.Callee
}

func (s *StartPrivacyCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *StartPrivacyCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartPrivacyCallRequest) SetAppId(v string) *StartPrivacyCallRequest {
	s.AppId = &v
	return s
}

func (s *StartPrivacyCallRequest) SetCallee(v string) *StartPrivacyCallRequest {
	s.Callee = &v
	return s
}

func (s *StartPrivacyCallRequest) SetCaller(v string) *StartPrivacyCallRequest {
	s.Caller = &v
	return s
}

func (s *StartPrivacyCallRequest) SetInstanceId(v string) *StartPrivacyCallRequest {
	s.InstanceId = &v
	return s
}

func (s *StartPrivacyCallRequest) Validate() error {
	return dara.Validate(s)
}
