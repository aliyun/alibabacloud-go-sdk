// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrometheusManagedInstance interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *PrometheusManagedInstance
	GetCreateTime() *string
	SetInstanceType(v string) *PrometheusManagedInstance
	GetInstanceType() *string
	SetPrometheusInstanceId(v string) *PrometheusManagedInstance
	GetPrometheusInstanceId() *string
	SetPrometheusInstanceName(v string) *PrometheusManagedInstance
	GetPrometheusInstanceName() *string
	SetRegionId(v string) *PrometheusManagedInstance
	GetRegionId() *string
	SetStatus(v string) *PrometheusManagedInstance
	GetStatus() *string
	SetWorkspace(v string) *PrometheusManagedInstance
	GetWorkspace() *string
}

type PrometheusManagedInstance struct {
	CreateTime             *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	InstanceType           *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	PrometheusInstanceId   *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	RegionId               *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
	Workspace              *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s PrometheusManagedInstance) String() string {
	return dara.Prettify(s)
}

func (s PrometheusManagedInstance) GoString() string {
	return s.String()
}

func (s *PrometheusManagedInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PrometheusManagedInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *PrometheusManagedInstance) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *PrometheusManagedInstance) GetPrometheusInstanceName() *string {
	return s.PrometheusInstanceName
}

func (s *PrometheusManagedInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *PrometheusManagedInstance) GetStatus() *string {
	return s.Status
}

func (s *PrometheusManagedInstance) GetWorkspace() *string {
	return s.Workspace
}

func (s *PrometheusManagedInstance) SetCreateTime(v string) *PrometheusManagedInstance {
	s.CreateTime = &v
	return s
}

func (s *PrometheusManagedInstance) SetInstanceType(v string) *PrometheusManagedInstance {
	s.InstanceType = &v
	return s
}

func (s *PrometheusManagedInstance) SetPrometheusInstanceId(v string) *PrometheusManagedInstance {
	s.PrometheusInstanceId = &v
	return s
}

func (s *PrometheusManagedInstance) SetPrometheusInstanceName(v string) *PrometheusManagedInstance {
	s.PrometheusInstanceName = &v
	return s
}

func (s *PrometheusManagedInstance) SetRegionId(v string) *PrometheusManagedInstance {
	s.RegionId = &v
	return s
}

func (s *PrometheusManagedInstance) SetStatus(v string) *PrometheusManagedInstance {
	s.Status = &v
	return s
}

func (s *PrometheusManagedInstance) SetWorkspace(v string) *PrometheusManagedInstance {
	s.Workspace = &v
	return s
}

func (s *PrometheusManagedInstance) Validate() error {
	return dara.Validate(s)
}
