// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchAuthenticationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *LaunchAuthenticationRequest
	GetContactFlowId() *string
	SetContactFlowVariables(v string) *LaunchAuthenticationRequest
	GetContactFlowVariables() *string
	SetDeviceId(v string) *LaunchAuthenticationRequest
	GetDeviceId() *string
	SetInstanceId(v string) *LaunchAuthenticationRequest
	GetInstanceId() *string
	SetJobId(v string) *LaunchAuthenticationRequest
	GetJobId() *string
	SetUserId(v string) *LaunchAuthenticationRequest
	GetUserId() *string
}

type LaunchAuthenticationRequest struct {
	// The contact flow ID for the IVR identity verification flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// af145gfc-1108-4d55-8fca-f719bd512ebb
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// Variables passed to the contact flow. This parameter is optional. The configured variables can be retrieved and used within the IVR flow. The format is a JSON string representing a collection of key-value pairs.
	//
	// example:
	//
	// {
	//
	//       "customerID": "208880281831****",
	//
	//       "operateType": "cipherCode",
	//
	//       "taskId": "1234567890",
	//
	//       "crmOther": "123"
	//
	// }
	ContactFlowVariables *string `json:"ContactFlowVariables,omitempty" xml:"ContactFlowVariables,omitempty"`
	// Device ID. This parameter is meaningless and can be filled in with any value.
	//
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The call ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The agent ID that initiates identity verification.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LaunchAuthenticationRequest) String() string {
	return dara.Prettify(s)
}

func (s LaunchAuthenticationRequest) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *LaunchAuthenticationRequest) GetContactFlowVariables() *string {
	return s.ContactFlowVariables
}

func (s *LaunchAuthenticationRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *LaunchAuthenticationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LaunchAuthenticationRequest) GetJobId() *string {
	return s.JobId
}

func (s *LaunchAuthenticationRequest) GetUserId() *string {
	return s.UserId
}

func (s *LaunchAuthenticationRequest) SetContactFlowId(v string) *LaunchAuthenticationRequest {
	s.ContactFlowId = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetContactFlowVariables(v string) *LaunchAuthenticationRequest {
	s.ContactFlowVariables = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetDeviceId(v string) *LaunchAuthenticationRequest {
	s.DeviceId = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetInstanceId(v string) *LaunchAuthenticationRequest {
	s.InstanceId = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetJobId(v string) *LaunchAuthenticationRequest {
	s.JobId = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetUserId(v string) *LaunchAuthenticationRequest {
	s.UserId = &v
	return s
}

func (s *LaunchAuthenticationRequest) Validate() error {
	return dara.Validate(s)
}
