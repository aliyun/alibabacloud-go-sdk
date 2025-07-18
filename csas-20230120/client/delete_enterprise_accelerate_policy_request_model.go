// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseAcceleratePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEapId(v string) *DeleteEnterpriseAcceleratePolicyRequest
	GetEapId() *string
}

type DeleteEnterpriseAcceleratePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eap-6edfb9d722ef8429
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
}

func (s DeleteEnterpriseAcceleratePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseAcceleratePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseAcceleratePolicyRequest) GetEapId() *string {
	return s.EapId
}

func (s *DeleteEnterpriseAcceleratePolicyRequest) SetEapId(v string) *DeleteEnterpriseAcceleratePolicyRequest {
	s.EapId = &v
	return s
}

func (s *DeleteEnterpriseAcceleratePolicyRequest) Validate() error {
	return dara.Validate(s)
}
