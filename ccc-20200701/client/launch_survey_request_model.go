// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchSurveyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *LaunchSurveyRequest
	GetContactFlowId() *string
	SetContactFlowVariables(v string) *LaunchSurveyRequest
	GetContactFlowVariables() *string
	SetDeviceId(v string) *LaunchSurveyRequest
	GetDeviceId() *string
	SetInstanceId(v string) *LaunchSurveyRequest
	GetInstanceId() *string
	SetJobId(v string) *LaunchSurveyRequest
	GetJobId() *string
	SetSmsMetadataId(v string) *LaunchSurveyRequest
	GetSmsMetadataId() *string
	SetSurveyChannel(v string) *LaunchSurveyRequest
	GetSurveyChannel() *string
	SetSurveyTemplateId(v string) *LaunchSurveyRequest
	GetSurveyTemplateId() *string
	SetSurveyTemplateVariables(v string) *LaunchSurveyRequest
	GetSurveyTemplateVariables() *string
	SetUserId(v string) *LaunchSurveyRequest
	GetUserId() *string
}

type LaunchSurveyRequest struct {
	// example:
	//
	// 4685b65a-eb8f-11ec-8ea0-0242ac120002
	ContactFlowId        *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	ContactFlowVariables *string `json:"ContactFlowVariables,omitempty" xml:"ContactFlowVariables,omitempty"`
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-6580466654649****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 4685b65a-eb8f-11ec-8ea0-0242ac120002
	SmsMetadataId *string `json:"SmsMetadataId,omitempty" xml:"SmsMetadataId,omitempty"`
	// example:
	//
	// IVR
	SurveyChannel           *string `json:"SurveyChannel,omitempty" xml:"SurveyChannel,omitempty"`
	SurveyTemplateId        *string `json:"SurveyTemplateId,omitempty" xml:"SurveyTemplateId,omitempty"`
	SurveyTemplateVariables *string `json:"SurveyTemplateVariables,omitempty" xml:"SurveyTemplateVariables,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LaunchSurveyRequest) String() string {
	return dara.Prettify(s)
}

func (s LaunchSurveyRequest) GoString() string {
	return s.String()
}

func (s *LaunchSurveyRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *LaunchSurveyRequest) GetContactFlowVariables() *string {
	return s.ContactFlowVariables
}

func (s *LaunchSurveyRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *LaunchSurveyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LaunchSurveyRequest) GetJobId() *string {
	return s.JobId
}

func (s *LaunchSurveyRequest) GetSmsMetadataId() *string {
	return s.SmsMetadataId
}

func (s *LaunchSurveyRequest) GetSurveyChannel() *string {
	return s.SurveyChannel
}

func (s *LaunchSurveyRequest) GetSurveyTemplateId() *string {
	return s.SurveyTemplateId
}

func (s *LaunchSurveyRequest) GetSurveyTemplateVariables() *string {
	return s.SurveyTemplateVariables
}

func (s *LaunchSurveyRequest) GetUserId() *string {
	return s.UserId
}

func (s *LaunchSurveyRequest) SetContactFlowId(v string) *LaunchSurveyRequest {
	s.ContactFlowId = &v
	return s
}

func (s *LaunchSurveyRequest) SetContactFlowVariables(v string) *LaunchSurveyRequest {
	s.ContactFlowVariables = &v
	return s
}

func (s *LaunchSurveyRequest) SetDeviceId(v string) *LaunchSurveyRequest {
	s.DeviceId = &v
	return s
}

func (s *LaunchSurveyRequest) SetInstanceId(v string) *LaunchSurveyRequest {
	s.InstanceId = &v
	return s
}

func (s *LaunchSurveyRequest) SetJobId(v string) *LaunchSurveyRequest {
	s.JobId = &v
	return s
}

func (s *LaunchSurveyRequest) SetSmsMetadataId(v string) *LaunchSurveyRequest {
	s.SmsMetadataId = &v
	return s
}

func (s *LaunchSurveyRequest) SetSurveyChannel(v string) *LaunchSurveyRequest {
	s.SurveyChannel = &v
	return s
}

func (s *LaunchSurveyRequest) SetSurveyTemplateId(v string) *LaunchSurveyRequest {
	s.SurveyTemplateId = &v
	return s
}

func (s *LaunchSurveyRequest) SetSurveyTemplateVariables(v string) *LaunchSurveyRequest {
	s.SurveyTemplateVariables = &v
	return s
}

func (s *LaunchSurveyRequest) SetUserId(v string) *LaunchSurveyRequest {
	s.UserId = &v
	return s
}

func (s *LaunchSurveyRequest) Validate() error {
	return dara.Validate(s)
}
