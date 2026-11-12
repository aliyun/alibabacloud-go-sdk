// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestartServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *RestartServiceRequest
	GetDryRun() *bool
}

type RestartServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 5db780a0-9d3e-4f35-a84a-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: Performs a dry run without executing the actual operation. If the request fails the dry run, an error is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - false: Sends a normal request.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s RestartServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartServiceRequest) GoString() string {
	return s.String()
}

func (s *RestartServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RestartServiceRequest) SetClientToken(v string) *RestartServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartServiceRequest) SetDryRun(v bool) *RestartServiceRequest {
	s.DryRun = &v
	return s
}

func (s *RestartServiceRequest) Validate() error {
	return dara.Validate(s)
}
