// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *StartServiceRequest
	GetDryRun() *bool
}

type StartServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 5db780a0-9d3e-4f35-a84a-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: performs a dry run without executing the actual operation. If the request fails the dry run, an error is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - false: sends a normal request.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s StartServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartServiceRequest) GoString() string {
	return s.String()
}

func (s *StartServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StartServiceRequest) SetClientToken(v string) *StartServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartServiceRequest) SetDryRun(v bool) *StartServiceRequest {
	s.DryRun = &v
	return s
}

func (s *StartServiceRequest) Validate() error {
	return dara.Validate(s)
}
