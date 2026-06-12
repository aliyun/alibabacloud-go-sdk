// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *RenewServiceInstanceRequest
	GetDryRun() *bool
	SetServiceInstanceId(v string) *RenewServiceInstanceRequest
	GetServiceInstanceId() *string
}

type RenewServiceInstanceRequest struct {
	// Specifies whether to perform a dry run of the renewal request, including permission and instance status checks. Valid values:
	//
	// - **true**: Sends the request without renewing the service instance.
	//
	// - **false**: Sends the request and renews the service instance after the checks pass.
	//
	// Default value: false. The operation is allowed only when the service instance is in the Pending Renewal or Renewal Failed state.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-70a3b15bb626435b****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RenewServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RenewServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *RenewServiceInstanceRequest) SetDryRun(v bool) *RenewServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *RenewServiceInstanceRequest) SetServiceInstanceId(v string) *RenewServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *RenewServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
