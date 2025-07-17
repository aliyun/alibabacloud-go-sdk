// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddPrometheusInstanceRequest
	GetName() *string
	SetRegionId(v string) *AddPrometheusInstanceRequest
	GetRegionId() *string
	SetType(v string) *AddPrometheusInstanceRequest
	GetType() *string
}

type AddPrometheusInstanceRequest struct {
	// The name of the Prometheus instance for Remote Write.
	//
	// This parameter is required.
	//
	// example:
	//
	// notificationpolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the Prometheus instance. Only Prometheus instances for Remote Write is supported. Set the value to RW.
	//
	// This parameter is required.
	//
	// example:
	//
	// RW
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddPrometheusInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *AddPrometheusInstanceRequest) GetName() *string {
	return s.Name
}

func (s *AddPrometheusInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPrometheusInstanceRequest) GetType() *string {
	return s.Type
}

func (s *AddPrometheusInstanceRequest) SetName(v string) *AddPrometheusInstanceRequest {
	s.Name = &v
	return s
}

func (s *AddPrometheusInstanceRequest) SetRegionId(v string) *AddPrometheusInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *AddPrometheusInstanceRequest) SetType(v string) *AddPrometheusInstanceRequest {
	s.Type = &v
	return s
}

func (s *AddPrometheusInstanceRequest) Validate() error {
	return dara.Validate(s)
}
