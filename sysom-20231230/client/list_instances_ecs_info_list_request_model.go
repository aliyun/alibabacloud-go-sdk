// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesEcsInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListInstancesEcsInfoListRequest
	GetXDebugId() *string
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
	SetXSysomInvokeSource(v string) *ListInstancesEcsInfoListRequest
	GetXSysomInvokeSource() *string
}

type ListInstancesEcsInfoListRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The type of information to retrieve.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip
	InfoType *string `json:"info_type,omitempty" xml:"info_type,omitempty"`
	// Specifies the instance ID to filter the Agent installation status of the specified instance.
	//
	// example:
	//
	// i-bp118piqcio9tiwgh84b
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// The management status of the instance.
	//
	// example:
	//
	// managed
	ManagedType *string `json:"managed_type,omitempty" xml:"managed_type,omitempty"`
	// Specifies the component ID to filter the instance information list for the corresponding component.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// The region used to filter instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region             *string `json:"region,omitempty" xml:"region,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListInstancesEcsInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesEcsInfoListRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesEcsInfoListRequest) GetXDebugId() *string {
	return s.XDebugId
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

func (s *ListInstancesEcsInfoListRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListInstancesEcsInfoListRequest) SetXDebugId(v string) *ListInstancesEcsInfoListRequest {
	s.XDebugId = &v
	return s
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

func (s *ListInstancesEcsInfoListRequest) SetXSysomInvokeSource(v string) *ListInstancesEcsInfoListRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListInstancesEcsInfoListRequest) Validate() error {
	return dara.Validate(s)
}
