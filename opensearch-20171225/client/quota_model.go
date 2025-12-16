// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuota interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResource(v int32) *Quota
	GetComputeResource() *int32
	SetDocSize(v int32) *Quota
	GetDocSize() *int32
	SetOrderType(v string) *Quota
	GetOrderType() *string
	SetSpec(v string) *Quota
	GetSpec() *string
}

type Quota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	OrderType       *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s Quota) String() string {
	return dara.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *Quota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *Quota) GetOrderType() *string {
	return s.OrderType
}

func (s *Quota) GetSpec() *string {
	return s.Spec
}

func (s *Quota) SetComputeResource(v int32) *Quota {
	s.ComputeResource = &v
	return s
}

func (s *Quota) SetDocSize(v int32) *Quota {
	s.DocSize = &v
	return s
}

func (s *Quota) SetOrderType(v string) *Quota {
	s.OrderType = &v
	return s
}

func (s *Quota) SetSpec(v string) *Quota {
	s.Spec = &v
	return s
}

func (s *Quota) Validate() error {
	return dara.Validate(s)
}
