// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyHealthCheckTemplateToServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ApplyHealthCheckTemplateToServerGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *ApplyHealthCheckTemplateToServerGroupRequest
	GetDryRun() *bool
	SetHealthCheckTemplateId(v string) *ApplyHealthCheckTemplateToServerGroupRequest
	GetHealthCheckTemplateId() *string
	SetServerGroupId(v string) *ApplyHealthCheckTemplateToServerGroupRequest
	GetServerGroupId() *string
}

type ApplyHealthCheckTemplateToServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hct-bp1qjwo61pqz3ahltv****
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-n80wyad08u0tox****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ApplyHealthCheckTemplateToServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyHealthCheckTemplateToServerGroupRequest) GoString() string {
	return s.String()
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) GetHealthCheckTemplateId() *string {
	return s.HealthCheckTemplateId
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) SetClientToken(v string) *ApplyHealthCheckTemplateToServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) SetDryRun(v bool) *ApplyHealthCheckTemplateToServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) SetHealthCheckTemplateId(v string) *ApplyHealthCheckTemplateToServerGroupRequest {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) SetServerGroupId(v string) *ApplyHealthCheckTemplateToServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) Validate() error {
	return dara.Validate(s)
}
