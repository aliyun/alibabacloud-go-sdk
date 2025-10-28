// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJvmConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateJvmConfigurationRequest
	GetAppId() *string
	SetGroupId(v string) *UpdateJvmConfigurationRequest
	GetGroupId() *string
	SetMaxHeapSize(v int32) *UpdateJvmConfigurationRequest
	GetMaxHeapSize() *int32
	SetMaxPermSize(v int32) *UpdateJvmConfigurationRequest
	GetMaxPermSize() *int32
	SetMinHeapSize(v int32) *UpdateJvmConfigurationRequest
	GetMinHeapSize() *int32
	SetOptions(v string) *UpdateJvmConfigurationRequest
	GetOptions() *string
}

type UpdateJvmConfigurationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// c627c157-560d-*************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance group where the application is deployed. You can call the ListDeployGroup operation to query the group ID. For more information, see [ListDeployGroup](https://help.aliyun.com/document_detail/62077.html).
	//
	// >
	//
	// 	- To configure the JVM parameters for an instance group, set this parameter to a specific ID.
	//
	// 	- To configure the JVM parameters for an application, leave this parameter empty.
	//
	// example:
	//
	// 0afc726e-077e-4357-98b2-db9f7145****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum size of the heap memory. Unit: MB.
	//
	// >
	//
	// 	- If this parameter is not specified in the group configuration, the value specified in the application configuration is used.
	//
	// 	- If this parameter is not specified in the application configuration, the default value is used.
	//
	// example:
	//
	// 500
	MaxHeapSize *int32 `json:"MaxHeapSize,omitempty" xml:"MaxHeapSize,omitempty"`
	// The size of the permanent generation heap memory. Unit: MB.
	//
	// >
	//
	// 	- If this parameter is not specified in the group configuration, the value specified in the application configuration is used.
	//
	// 	- If this parameter is not specified in the application configuration, the default value is used.
	//
	// example:
	//
	// 1000
	MaxPermSize *int32 `json:"MaxPermSize,omitempty" xml:"MaxPermSize,omitempty"`
	// The initial size of the heap memory. Unit: MB.
	//
	// >
	//
	// 	- If this parameter is not specified in the group configuration, the value specified in the application configuration is used.
	//
	// 	- If this parameter is not specified in the application configuration, the default value is used.
	//
	// example:
	//
	// 500
	MinHeapSize *int32 `json:"MinHeapSize,omitempty" xml:"MinHeapSize,omitempty"`
	// The custom JVM parameters.
	//
	// >
	//
	// 	- If this parameter is not specified in the group configuration, the value specified in the application configuration is used.
	//
	// 	- If this parameter is not specified in the application configuration, the default value is used.
	//
	// example:
	//
	// -Dproperty=value
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s UpdateJvmConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJvmConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateJvmConfigurationRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateJvmConfigurationRequest) GetMaxHeapSize() *int32 {
	return s.MaxHeapSize
}

func (s *UpdateJvmConfigurationRequest) GetMaxPermSize() *int32 {
	return s.MaxPermSize
}

func (s *UpdateJvmConfigurationRequest) GetMinHeapSize() *int32 {
	return s.MinHeapSize
}

func (s *UpdateJvmConfigurationRequest) GetOptions() *string {
	return s.Options
}

func (s *UpdateJvmConfigurationRequest) SetAppId(v string) *UpdateJvmConfigurationRequest {
	s.AppId = &v
	return s
}

func (s *UpdateJvmConfigurationRequest) SetGroupId(v string) *UpdateJvmConfigurationRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateJvmConfigurationRequest) SetMaxHeapSize(v int32) *UpdateJvmConfigurationRequest {
	s.MaxHeapSize = &v
	return s
}

func (s *UpdateJvmConfigurationRequest) SetMaxPermSize(v int32) *UpdateJvmConfigurationRequest {
	s.MaxPermSize = &v
	return s
}

func (s *UpdateJvmConfigurationRequest) SetMinHeapSize(v int32) *UpdateJvmConfigurationRequest {
	s.MinHeapSize = &v
	return s
}

func (s *UpdateJvmConfigurationRequest) SetOptions(v string) *UpdateJvmConfigurationRequest {
	s.Options = &v
	return s
}

func (s *UpdateJvmConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
