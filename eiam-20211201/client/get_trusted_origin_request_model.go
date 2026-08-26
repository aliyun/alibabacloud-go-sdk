// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrustedOriginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTrustedOriginRequest
	GetInstanceId() *string
	SetTrustedOriginId(v string) *GetTrustedOriginRequest
	GetTrustedOriginId() *string
}

type GetTrustedOriginRequest struct {
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

func (s GetTrustedOriginRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrustedOriginRequest) GoString() string {
	return s.String()
}

func (s *GetTrustedOriginRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTrustedOriginRequest) GetTrustedOriginId() *string {
	return s.TrustedOriginId
}

func (s *GetTrustedOriginRequest) SetInstanceId(v string) *GetTrustedOriginRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTrustedOriginRequest) SetTrustedOriginId(v string) *GetTrustedOriginRequest {
	s.TrustedOriginId = &v
	return s
}

func (s *GetTrustedOriginRequest) Validate() error {
	return dara.Validate(s)
}
