// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEnterpriseAcceleratePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEapId(v string) *DisableEnterpriseAcceleratePolicyRequest
	GetEapId() *string
}

type DisableEnterpriseAcceleratePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eap-d50b45aa7dc04aef
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
}

func (s DisableEnterpriseAcceleratePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableEnterpriseAcceleratePolicyRequest) GoString() string {
	return s.String()
}

func (s *DisableEnterpriseAcceleratePolicyRequest) GetEapId() *string {
	return s.EapId
}

func (s *DisableEnterpriseAcceleratePolicyRequest) SetEapId(v string) *DisableEnterpriseAcceleratePolicyRequest {
	s.EapId = &v
	return s
}

func (s *DisableEnterpriseAcceleratePolicyRequest) Validate() error {
	return dara.Validate(s)
}
