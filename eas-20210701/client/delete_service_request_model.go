// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteServiceRequest
	GetDryRun() *bool
}

type DeleteServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 5db780a0-9d3e-4f35-a84a-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request.
	//
	// - true: performs a dry run without executing the actual operation. If the dry run fails, the corresponding error is returned. If the dry run succeeds, the error code DryRunOperation is returned.
	//
	// - false (default): sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteServiceRequest) SetClientToken(v string) *DeleteServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceRequest) SetDryRun(v bool) *DeleteServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServiceRequest) Validate() error {
	return dara.Validate(s)
}
