// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTrustedOriginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableTrustedOriginRequest
	GetInstanceId() *string
	SetTrustedOriginId(v string) *DisableTrustedOriginRequest
	GetTrustedOriginId() *string
}

type DisableTrustedOriginRequest struct {
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the trusted origin.
	//
	// This parameter is required.
	//
	// example:
	//
	// to_example
	TrustedOriginId *string `json:"TrustedOriginId,omitempty" xml:"TrustedOriginId,omitempty"`
}

func (s DisableTrustedOriginRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableTrustedOriginRequest) GoString() string {
	return s.String()
}

func (s *DisableTrustedOriginRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableTrustedOriginRequest) GetTrustedOriginId() *string {
	return s.TrustedOriginId
}

func (s *DisableTrustedOriginRequest) SetInstanceId(v string) *DisableTrustedOriginRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableTrustedOriginRequest) SetTrustedOriginId(v string) *DisableTrustedOriginRequest {
	s.TrustedOriginId = &v
	return s
}

func (s *DisableTrustedOriginRequest) Validate() error {
	return dara.Validate(s)
}
