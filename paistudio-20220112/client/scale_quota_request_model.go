// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMin(v *ResourceSpec) *ScaleQuotaRequest
	GetMin() *ResourceSpec
	SetResourceGroupIds(v []*string) *ScaleQuotaRequest
	GetResourceGroupIds() []*string
}

type ScaleQuotaRequest struct {
	Min              *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	ResourceGroupIds []*string     `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s ScaleQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleQuotaRequest) GoString() string {
	return s.String()
}

func (s *ScaleQuotaRequest) GetMin() *ResourceSpec {
	return s.Min
}

func (s *ScaleQuotaRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *ScaleQuotaRequest) SetMin(v *ResourceSpec) *ScaleQuotaRequest {
	s.Min = v
	return s
}

func (s *ScaleQuotaRequest) SetResourceGroupIds(v []*string) *ScaleQuotaRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *ScaleQuotaRequest) Validate() error {
	return dara.Validate(s)
}
