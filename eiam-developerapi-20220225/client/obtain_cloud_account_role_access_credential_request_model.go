// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCloudAccountRoleAccessCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountRoleExternalId(v string) *ObtainCloudAccountRoleAccessCredentialRequest
	GetCloudAccountRoleExternalId() *string
	SetDurationSeconds(v int32) *ObtainCloudAccountRoleAccessCredentialRequest
	GetDurationSeconds() *int32
}

type ObtainCloudAccountRoleAccessCredentialRequest struct {
	// The business identifier of the cloud account role.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::xxx:role/role-test
	CloudAccountRoleExternalId *string `json:"cloudAccountRoleExternalId,omitempty" xml:"cloudAccountRoleExternalId,omitempty"`
	// Specifies the validity duration of the temporary security credentials (STS Token) for the cloud account role, in seconds. Valid values: 900 to 43200 (15 minutes to 12 hours).
	//
	// Constraints:
	//
	// - The minimum value cannot be less than 900 seconds.
	//
	// - The maximum value is subject to the maximum session duration of the cloud provider role or service account. For example, the default maximum session duration for an AWS role is 3600 seconds.
	//
	// example:
	//
	// 1800
	DurationSeconds *int32 `json:"durationSeconds,omitempty" xml:"durationSeconds,omitempty"`
}

func (s ObtainCloudAccountRoleAccessCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialRequest) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialRequest) GetCloudAccountRoleExternalId() *string {
	return s.CloudAccountRoleExternalId
}

func (s *ObtainCloudAccountRoleAccessCredentialRequest) GetDurationSeconds() *int32 {
	return s.DurationSeconds
}

func (s *ObtainCloudAccountRoleAccessCredentialRequest) SetCloudAccountRoleExternalId(v string) *ObtainCloudAccountRoleAccessCredentialRequest {
	s.CloudAccountRoleExternalId = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialRequest) SetDurationSeconds(v int32) *ObtainCloudAccountRoleAccessCredentialRequest {
	s.DurationSeconds = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialRequest) Validate() error {
	return dara.Validate(s)
}
