// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseAccelerateTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEapId(v string) *CreateEnterpriseAccelerateTargetRequest
	GetEapId() *string
	SetTarget(v []*string) *CreateEnterpriseAccelerateTargetRequest
	GetTarget() []*string
}

type CreateEnterpriseAccelerateTargetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eap-eec34d4b12fcca61
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
	// This parameter is required.
	Target []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s CreateEnterpriseAccelerateTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseAccelerateTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseAccelerateTargetRequest) GetEapId() *string {
	return s.EapId
}

func (s *CreateEnterpriseAccelerateTargetRequest) GetTarget() []*string {
	return s.Target
}

func (s *CreateEnterpriseAccelerateTargetRequest) SetEapId(v string) *CreateEnterpriseAccelerateTargetRequest {
	s.EapId = &v
	return s
}

func (s *CreateEnterpriseAccelerateTargetRequest) SetTarget(v []*string) *CreateEnterpriseAccelerateTargetRequest {
	s.Target = v
	return s
}

func (s *CreateEnterpriseAccelerateTargetRequest) Validate() error {
	return dara.Validate(s)
}
