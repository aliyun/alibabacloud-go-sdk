// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainer(v string) *DeleteServiceInstancesRequest
	GetContainer() *string
	SetInstanceList(v string) *DeleteServiceInstancesRequest
	GetInstanceList() *string
	SetSoftRestart(v bool) *DeleteServiceInstancesRequest
	GetSoftRestart() *bool
}

type DeleteServiceInstancesRequest struct {
	// The name of the container whose process needs to be restarted. This parameter takes effect only if the SoftRestart parameter is set to true.
	//
	// example:
	//
	// worker0
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The instances that you want to restart. Separate multiple instance names with commas (,). For more information about how to query the instance name, see [ListServiceInstances](https://help.aliyun.com/document_detail/412108.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// foo-rdsbxxxx,foo-rdsaxxxx
	InstanceList *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
	// Specifies whether to restart only the container process without recreating the instance. Default value: false. Valid values: true and false.
	//
	// example:
	//
	// true
	SoftRestart *bool `json:"SoftRestart,omitempty" xml:"SoftRestart,omitempty"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) GetContainer() *string {
	return s.Container
}

func (s *DeleteServiceInstancesRequest) GetInstanceList() *string {
	return s.InstanceList
}

func (s *DeleteServiceInstancesRequest) GetSoftRestart() *bool {
	return s.SoftRestart
}

func (s *DeleteServiceInstancesRequest) SetContainer(v string) *DeleteServiceInstancesRequest {
	s.Container = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetInstanceList(v string) *DeleteServiceInstancesRequest {
	s.InstanceList = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetSoftRestart(v bool) *DeleteServiceInstancesRequest {
	s.SoftRestart = &v
	return s
}

func (s *DeleteServiceInstancesRequest) Validate() error {
	return dara.Validate(s)
}
