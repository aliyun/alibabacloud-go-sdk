// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePADiagnosisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevTag(v string) *CreatePADiagnosisTaskRequest
	GetDevTag() *string
	SetDiagnoseType(v string) *CreatePADiagnosisTaskRequest
	GetDiagnoseType() *string
	SetHost(v string) *CreatePADiagnosisTaskRequest
	GetHost() *string
	SetPopId(v string) *CreatePADiagnosisTaskRequest
	GetPopId() *string
	SetPopMode(v string) *CreatePADiagnosisTaskRequest
	GetPopMode() *string
	SetPort(v string) *CreatePADiagnosisTaskRequest
	GetPort() *string
	SetProtocol(v string) *CreatePADiagnosisTaskRequest
	GetProtocol() *string
	SetUdpExtraConfigs(v *CreatePADiagnosisTaskRequestUdpExtraConfigs) *CreatePADiagnosisTaskRequest
	GetUdpExtraConfigs() *CreatePADiagnosisTaskRequestUdpExtraConfigs
	SetUserGroupId(v string) *CreatePADiagnosisTaskRequest
	GetUserGroupId() *string
	SetUsername(v string) *CreatePADiagnosisTaskRequest
	GetUsername() *string
}

type CreatePADiagnosisTaskRequest struct {
	// example:
	//
	// 2987b3e0-8108-2f99-4d18-3b4f1c1c36d7
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FullLink
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// socialapp-gateway.client9.me
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// pop-3e244b62357dcafc
	PopId *string `json:"PopId,omitempty" xml:"PopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ManualSelect
	PopMode *string `json:"PopMode,omitempty" xml:"PopMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TCP
	Protocol        *string                                      `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	UdpExtraConfigs *CreatePADiagnosisTaskRequestUdpExtraConfigs `json:"UdpExtraConfigs,omitempty" xml:"UdpExtraConfigs,omitempty" type:"Struct"`
	UserGroupId     *string                                      `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// example:
	//
	// Zhaosi
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreatePADiagnosisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePADiagnosisTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePADiagnosisTaskRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *CreatePADiagnosisTaskRequest) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *CreatePADiagnosisTaskRequest) GetHost() *string {
	return s.Host
}

func (s *CreatePADiagnosisTaskRequest) GetPopId() *string {
	return s.PopId
}

func (s *CreatePADiagnosisTaskRequest) GetPopMode() *string {
	return s.PopMode
}

func (s *CreatePADiagnosisTaskRequest) GetPort() *string {
	return s.Port
}

func (s *CreatePADiagnosisTaskRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreatePADiagnosisTaskRequest) GetUdpExtraConfigs() *CreatePADiagnosisTaskRequestUdpExtraConfigs {
	return s.UdpExtraConfigs
}

func (s *CreatePADiagnosisTaskRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *CreatePADiagnosisTaskRequest) GetUsername() *string {
	return s.Username
}

func (s *CreatePADiagnosisTaskRequest) SetDevTag(v string) *CreatePADiagnosisTaskRequest {
	s.DevTag = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetDiagnoseType(v string) *CreatePADiagnosisTaskRequest {
	s.DiagnoseType = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetHost(v string) *CreatePADiagnosisTaskRequest {
	s.Host = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetPopId(v string) *CreatePADiagnosisTaskRequest {
	s.PopId = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetPopMode(v string) *CreatePADiagnosisTaskRequest {
	s.PopMode = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetPort(v string) *CreatePADiagnosisTaskRequest {
	s.Port = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetProtocol(v string) *CreatePADiagnosisTaskRequest {
	s.Protocol = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetUdpExtraConfigs(v *CreatePADiagnosisTaskRequestUdpExtraConfigs) *CreatePADiagnosisTaskRequest {
	s.UdpExtraConfigs = v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetUserGroupId(v string) *CreatePADiagnosisTaskRequest {
	s.UserGroupId = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) SetUsername(v string) *CreatePADiagnosisTaskRequest {
	s.Username = &v
	return s
}

func (s *CreatePADiagnosisTaskRequest) Validate() error {
	if s.UdpExtraConfigs != nil {
		if err := s.UdpExtraConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePADiagnosisTaskRequestUdpExtraConfigs struct {
	// example:
	//
	// hello
	ExpectedResponse *string `json:"ExpectedResponse,omitempty" xml:"ExpectedResponse,omitempty"`
	// example:
	//
	// hello
	RequestContent *string `json:"RequestContent,omitempty" xml:"RequestContent,omitempty"`
}

func (s CreatePADiagnosisTaskRequestUdpExtraConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreatePADiagnosisTaskRequestUdpExtraConfigs) GoString() string {
	return s.String()
}

func (s *CreatePADiagnosisTaskRequestUdpExtraConfigs) GetExpectedResponse() *string {
	return s.ExpectedResponse
}

func (s *CreatePADiagnosisTaskRequestUdpExtraConfigs) GetRequestContent() *string {
	return s.RequestContent
}

func (s *CreatePADiagnosisTaskRequestUdpExtraConfigs) SetExpectedResponse(v string) *CreatePADiagnosisTaskRequestUdpExtraConfigs {
	s.ExpectedResponse = &v
	return s
}

func (s *CreatePADiagnosisTaskRequestUdpExtraConfigs) SetRequestContent(v string) *CreatePADiagnosisTaskRequestUdpExtraConfigs {
	s.RequestContent = &v
	return s
}

func (s *CreatePADiagnosisTaskRequestUdpExtraConfigs) Validate() error {
	return dara.Validate(s)
}
