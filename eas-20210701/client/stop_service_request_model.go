// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *StopServiceRequest
	GetDryRun() *bool
}

type StopServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 5db780a0-9d3e-4f35-a84a-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: A dry run is performed and the operation is not executed. If the request fails the dry run, an error is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - false: A normal request is sent.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s StopServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopServiceRequest) GoString() string {
	return s.String()
}

func (s *StopServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StopServiceRequest) SetClientToken(v string) *StopServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopServiceRequest) SetDryRun(v bool) *StopServiceRequest {
	s.DryRun = &v
	return s
}

func (s *StopServiceRequest) Validate() error {
	return dara.Validate(s)
}
