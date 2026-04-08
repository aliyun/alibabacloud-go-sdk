// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePADiagnosisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosisTask(v *CreatePADiagnosisTaskResponseBodyDiagnosisTask) *CreatePADiagnosisTaskResponseBody
	GetDiagnosisTask() *CreatePADiagnosisTaskResponseBodyDiagnosisTask
	SetRequestId(v string) *CreatePADiagnosisTaskResponseBody
	GetRequestId() *string
}

type CreatePADiagnosisTaskResponseBody struct {
	DiagnosisTask *CreatePADiagnosisTaskResponseBodyDiagnosisTask `json:"DiagnosisTask,omitempty" xml:"DiagnosisTask,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 5F79AE39-6622-5292-87EF-DE45631DE4D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePADiagnosisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePADiagnosisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePADiagnosisTaskResponseBody) GetDiagnosisTask() *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	return s.DiagnosisTask
}

func (s *CreatePADiagnosisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePADiagnosisTaskResponseBody) SetDiagnosisTask(v *CreatePADiagnosisTaskResponseBodyDiagnosisTask) *CreatePADiagnosisTaskResponseBody {
	s.DiagnosisTask = v
	return s
}

func (s *CreatePADiagnosisTaskResponseBody) SetRequestId(v string) *CreatePADiagnosisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBody) Validate() error {
	if s.DiagnosisTask != nil {
		if err := s.DiagnosisTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePADiagnosisTaskResponseBodyDiagnosisTask struct {
	// example:
	//
	// E9EE1CE7-4AA0-521D-B8E1-E13E47F05E94
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// example:
	//
	// diag-3e0d36d6c15a0502
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// example:
	//
	// FullLink
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// example:
	//
	// 172.16.6.1
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// pop-8ded63ce9d3d317e
	PopId *string `json:"PopId,omitempty" xml:"PopId,omitempty"`
	// example:
	//
	// AutoSelect
	PopMode *string `json:"PopMode,omitempty" xml:"PopMode,omitempty"`
	// example:
	//
	// 443
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Disabled
	Status          *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	UdpExtraConfigs *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs `json:"UdpExtraConfigs,omitempty" xml:"UdpExtraConfigs,omitempty" type:"Struct"`
	UserGroup       *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup       `json:"UserGroup,omitempty" xml:"UserGroup,omitempty" type:"Struct"`
	// example:
	//
	// zhangsan
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreatePADiagnosisTaskResponseBodyDiagnosisTask) String() string {
	return dara.Prettify(s)
}

func (s CreatePADiagnosisTaskResponseBodyDiagnosisTask) GoString() string {
	return s.String()
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetDevTag() *string {
	return s.DevTag
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetHost() *string {
	return s.Host
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetPopId() *string {
	return s.PopId
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetPopMode() *string {
	return s.PopMode
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetPort() *string {
	return s.Port
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetProtocol() *string {
	return s.Protocol
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetStatus() *string {
	return s.Status
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetUdpExtraConfigs() *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs {
	return s.UdpExtraConfigs
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetUserGroup() *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup {
	return s.UserGroup
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) GetUsername() *string {
	return s.Username
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetDevTag(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.DevTag = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetDiagnoseId(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.DiagnoseId = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetDiagnoseType(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.DiagnoseType = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetHost(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.Host = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetPopId(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.PopId = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetPopMode(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.PopMode = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetPort(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.Port = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetProtocol(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.Protocol = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetStatus(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.Status = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetUdpExtraConfigs(v *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.UdpExtraConfigs = v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetUserGroup(v *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.UserGroup = v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) SetUsername(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTask {
	s.Username = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTask) Validate() error {
	if s.UdpExtraConfigs != nil {
		if err := s.UdpExtraConfigs.Validate(); err != nil {
			return err
		}
	}
	if s.UserGroup != nil {
		if err := s.UserGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs struct {
	// example:
	//
	// hello
	ExpectedResponse *string `json:"ExpectedResponse,omitempty" xml:"ExpectedResponse,omitempty"`
	// example:
	//
	// hello
	RequestContent *string `json:"RequestContent,omitempty" xml:"RequestContent,omitempty"`
}

func (s CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) GoString() string {
	return s.String()
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) GetExpectedResponse() *string {
	return s.ExpectedResponse
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) GetRequestContent() *string {
	return s.RequestContent
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) SetExpectedResponse(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs {
	s.ExpectedResponse = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) SetRequestContent(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs {
	s.RequestContent = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) Validate() error {
	return dara.Validate(s)
}

type CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup struct {
	// example:
	//
	// ug-xxxxxxxx
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// example:
	//
	// IT
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) String() string {
	return dara.Prettify(s)
}

func (s CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) GoString() string {
	return s.String()
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) SetUserGroupId(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup {
	s.UserGroupId = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) SetUserGroupName(v string) *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup {
	s.UserGroupName = &v
	return s
}

func (s *CreatePADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) Validate() error {
	return dara.Validate(s)
}
