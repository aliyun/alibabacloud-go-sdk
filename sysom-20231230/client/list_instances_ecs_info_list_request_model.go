// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesEcsInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInfoType(v string) *ListInstancesEcsInfoListRequest
	GetInfoType() *string
	SetInstanceId(v string) *ListInstancesEcsInfoListRequest
	GetInstanceId() *string
	SetManagedType(v string) *ListInstancesEcsInfoListRequest
	GetManagedType() *string
	SetPluginId(v string) *ListInstancesEcsInfoListRequest
	GetPluginId() *string
	SetRegion(v string) *ListInstancesEcsInfoListRequest
	GetRegion() *string
}

type ListInstancesEcsInfoListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ip
	InfoType *string `json:"info_type,omitempty" xml:"info_type,omitempty"`
	// example:
	//
	// i-bp118piqcio9tiwgh84b
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// example:
	//
	// managed
	ManagedType *string `json:"managed_type,omitempty" xml:"managed_type,omitempty"`
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s ListInstancesEcsInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesEcsInfoListRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesEcsInfoListRequest) GetInfoType() *string {
	return s.InfoType
}

func (s *ListInstancesEcsInfoListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesEcsInfoListRequest) GetManagedType() *string {
	return s.ManagedType
}

func (s *ListInstancesEcsInfoListRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListInstancesEcsInfoListRequest) GetRegion() *string {
	return s.Region
}

func (s *ListInstancesEcsInfoListRequest) SetInfoType(v string) *ListInstancesEcsInfoListRequest {
	s.InfoType = &v
	return s
}

func (s *ListInstancesEcsInfoListRequest) SetInstanceId(v string) *ListInstancesEcsInfoListRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesEcsInfoListRequest) SetManagedType(v string) *ListInstancesEcsInfoListRequest {
	s.ManagedType = &v
	return s
}

func (s *ListInstancesEcsInfoListRequest) SetPluginId(v string) *ListInstancesEcsInfoListRequest {
	s.PluginId = &v
	return s
}

func (s *ListInstancesEcsInfoListRequest) SetRegion(v string) *ListInstancesEcsInfoListRequest {
	s.Region = &v
	return s
}

func (s *ListInstancesEcsInfoListRequest) Validate() error {
	return dara.Validate(s)
}
