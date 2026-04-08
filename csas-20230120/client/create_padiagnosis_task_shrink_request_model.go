// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePADiagnosisTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevTag(v string) *CreatePADiagnosisTaskShrinkRequest
	GetDevTag() *string
	SetDiagnoseType(v string) *CreatePADiagnosisTaskShrinkRequest
	GetDiagnoseType() *string
	SetHost(v string) *CreatePADiagnosisTaskShrinkRequest
	GetHost() *string
	SetPopId(v string) *CreatePADiagnosisTaskShrinkRequest
	GetPopId() *string
	SetPopMode(v string) *CreatePADiagnosisTaskShrinkRequest
	GetPopMode() *string
	SetPort(v string) *CreatePADiagnosisTaskShrinkRequest
	GetPort() *string
	SetProtocol(v string) *CreatePADiagnosisTaskShrinkRequest
	GetProtocol() *string
	SetUdpExtraConfigsShrink(v string) *CreatePADiagnosisTaskShrinkRequest
	GetUdpExtraConfigsShrink() *string
	SetUserGroupId(v string) *CreatePADiagnosisTaskShrinkRequest
	GetUserGroupId() *string
	SetUsername(v string) *CreatePADiagnosisTaskShrinkRequest
	GetUsername() *string
}

type CreatePADiagnosisTaskShrinkRequest struct {
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
	Protocol              *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	UdpExtraConfigsShrink *string `json:"UdpExtraConfigs,omitempty" xml:"UdpExtraConfigs,omitempty"`
	UserGroupId           *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// example:
	//
	// Zhaosi
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreatePADiagnosisTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePADiagnosisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetHost() *string {
	return s.Host
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetPopId() *string {
	return s.PopId
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetPopMode() *string {
	return s.PopMode
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetPort() *string {
	return s.Port
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetUdpExtraConfigsShrink() *string {
	return s.UdpExtraConfigsShrink
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *CreatePADiagnosisTaskShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetDevTag(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.DevTag = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetDiagnoseType(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.DiagnoseType = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetHost(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.Host = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetPopId(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.PopId = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetPopMode(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.PopMode = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetPort(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.Port = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetProtocol(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetUdpExtraConfigsShrink(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.UdpExtraConfigsShrink = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetUserGroupId(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.UserGroupId = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) SetUsername(v string) *CreatePADiagnosisTaskShrinkRequest {
	s.Username = &v
	return s
}

func (s *CreatePADiagnosisTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
