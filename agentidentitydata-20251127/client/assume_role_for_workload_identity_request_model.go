// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleForWorkloadIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDurationSeconds(v string) *AssumeRoleForWorkloadIdentityRequest
	GetDurationSeconds() *string
	SetPolicy(v string) *AssumeRoleForWorkloadIdentityRequest
	GetPolicy() *string
	SetRoleSessionName(v string) *AssumeRoleForWorkloadIdentityRequest
	GetRoleSessionName() *string
	SetWorkloadAccessToken(v string) *AssumeRoleForWorkloadIdentityRequest
	GetWorkloadAccessToken() *string
}

type AssumeRoleForWorkloadIdentityRequest struct {
	// example:
	//
	// 3600
	DurationSeconds *string `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// example:
	//
	// {"Statement": [{"Action": ["*"],"Effect": "Allow","Resource": ["*"]}],"Version":"1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// session-name
	RoleSessionName *string `json:"RoleSessionName,omitempty" xml:"RoleSessionName,omitempty"`
	// example:
	//
	// eyAgImFsZyI6ICJSUzI1NiIsIC****
	WorkloadAccessToken *string `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
}

func (s AssumeRoleForWorkloadIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleForWorkloadIdentityRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleForWorkloadIdentityRequest) GetDurationSeconds() *string {
	return s.DurationSeconds
}

func (s *AssumeRoleForWorkloadIdentityRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AssumeRoleForWorkloadIdentityRequest) GetRoleSessionName() *string {
	return s.RoleSessionName
}

func (s *AssumeRoleForWorkloadIdentityRequest) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *AssumeRoleForWorkloadIdentityRequest) SetDurationSeconds(v string) *AssumeRoleForWorkloadIdentityRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityRequest) SetPolicy(v string) *AssumeRoleForWorkloadIdentityRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityRequest) SetRoleSessionName(v string) *AssumeRoleForWorkloadIdentityRequest {
	s.RoleSessionName = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityRequest) SetWorkloadAccessToken(v string) *AssumeRoleForWorkloadIdentityRequest {
	s.WorkloadAccessToken = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityRequest) Validate() error {
	return dara.Validate(s)
}
