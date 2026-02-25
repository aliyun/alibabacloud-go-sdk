// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceFeatureGateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceFeatureGateRequest
	GetInstanceId() *string
}

type GetInstanceFeatureGateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceFeatureGateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceFeatureGateRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceFeatureGateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceFeatureGateRequest) SetInstanceId(v string) *GetInstanceFeatureGateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceFeatureGateRequest) Validate() error {
	return dara.Validate(s)
}
