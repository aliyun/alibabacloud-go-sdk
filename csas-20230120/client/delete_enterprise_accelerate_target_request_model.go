// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseAccelerateTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEapId(v string) *DeleteEnterpriseAccelerateTargetRequest
	GetEapId() *string
	SetTarget(v []*string) *DeleteEnterpriseAccelerateTargetRequest
	GetTarget() []*string
}

type DeleteEnterpriseAccelerateTargetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eap-ed1f0e392a28a4e6
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
	// This parameter is required.
	Target []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s DeleteEnterpriseAccelerateTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseAccelerateTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseAccelerateTargetRequest) GetEapId() *string {
	return s.EapId
}

func (s *DeleteEnterpriseAccelerateTargetRequest) GetTarget() []*string {
	return s.Target
}

func (s *DeleteEnterpriseAccelerateTargetRequest) SetEapId(v string) *DeleteEnterpriseAccelerateTargetRequest {
	s.EapId = &v
	return s
}

func (s *DeleteEnterpriseAccelerateTargetRequest) SetTarget(v []*string) *DeleteEnterpriseAccelerateTargetRequest {
	s.Target = v
	return s
}

func (s *DeleteEnterpriseAccelerateTargetRequest) Validate() error {
	return dara.Validate(s)
}
