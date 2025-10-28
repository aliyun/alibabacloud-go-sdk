// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sApplicationConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateK8sApplicationConfigRequest
	GetAppId() *string
	SetClusterId(v string) *UpdateK8sApplicationConfigRequest
	GetClusterId() *string
	SetCpuLimit(v string) *UpdateK8sApplicationConfigRequest
	GetCpuLimit() *string
	SetCpuRequest(v string) *UpdateK8sApplicationConfigRequest
	GetCpuRequest() *string
	SetEphemeralStorageLimit(v string) *UpdateK8sApplicationConfigRequest
	GetEphemeralStorageLimit() *string
	SetEphemeralStorageRequest(v string) *UpdateK8sApplicationConfigRequest
	GetEphemeralStorageRequest() *string
	SetMcpuLimit(v string) *UpdateK8sApplicationConfigRequest
	GetMcpuLimit() *string
	SetMcpuRequest(v string) *UpdateK8sApplicationConfigRequest
	GetMcpuRequest() *string
	SetMemoryLimit(v string) *UpdateK8sApplicationConfigRequest
	GetMemoryLimit() *string
	SetMemoryRequest(v string) *UpdateK8sApplicationConfigRequest
	GetMemoryRequest() *string
	SetTimeout(v int32) *UpdateK8sApplicationConfigRequest
	GetTimeout() *int32
}

type UpdateK8sApplicationConfigRequest struct {
	// The ID of the application. You can query the application ID by calling the ListApplication operation. For more information, see [ListApplication](https://help.aliyun.com/document_detail/423162.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 23bf94d9-****-4994-9917-616a827aa777
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the cluster. You can query the cluster ID by calling the ListCluster operation. For more information, see [ListCluster](https://help.aliyun.com/document_detail/411844.html).
	//
	// example:
	//
	// 9c28bbb9-****-44b3-b953-54ef8a2d0be2
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The maximum number of CPU cores allowed for each application instance when the application is running. The value 0 indicates that no limit is set on CPU cores.
	//
	// example:
	//
	// 1
	CpuLimit *string `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// The number of CPU cores requested for each application instance when the application is running. Unit: cores. We recommend that you set this parameter. The value 0 indicates that no limit is set on CPU cores.
	//
	// > You must set this parameter together with the CpuLimit parameter. Make sure that the value of this parameter does not exceed that of the CpuLimit parameter.
	//
	// example:
	//
	// 1
	CpuRequest *string `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	// The maximum size of space required by ephemeral storage. Unit: GB. The value 0 indicates that no limit is set on the ephemeral storage space.
	//
	// example:
	//
	// 4
	EphemeralStorageLimit *string `json:"EphemeralStorageLimit,omitempty" xml:"EphemeralStorageLimit,omitempty"`
	// The minimum size of space required by ephemeral storage. Unit: GB. The value 0 indicates that no limit is set on the ephemeral storage space.
	//
	// > You must set this parameter together with the EphemeralStorageLimit parameter. Make sure that the value of this parameter does not exceed that of the EphemeralStorageLimit parameter.
	//
	// example:
	//
	// 2
	EphemeralStorageRequest *string `json:"EphemeralStorageRequest,omitempty" xml:"EphemeralStorageRequest,omitempty"`
	// The maximum number of CPU cores allowed. The value 0 indicates that no limit is set on CPU cores.
	//
	// example:
	//
	// 1
	McpuLimit *string `json:"McpuLimit,omitempty" xml:"McpuLimit,omitempty"`
	// The minimum number of CPU cores required. Unit: cores. The value 0 indicates that no limit is set on CPU cores.
	//
	// > You must set this parameter together with the CpuLimit parameter. Make sure that the value of this parameter does not exceed that of the CpuLimit parameter.
	//
	// example:
	//
	// 1000
	McpuRequest *string `json:"McpuRequest,omitempty" xml:"McpuRequest,omitempty"`
	// The maximum size of memory allowed for each application instance when the application is running. Unit: MB. The value 0 indicates that no limit is set on the memory size.
	//
	// example:
	//
	// 4
	MemoryLimit *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The size of memory requested for each application instance when the application is running. Unit: MB. We recommend that you set this parameter. If you do not want to apply for a memory quota, set this parameter to 0.
	//
	// > You must set this parameter together with the MemoryLimit parameter. Make sure that the value of this parameter does not exceed that of the MemoryLimit parameter.
	//
	// example:
	//
	// 400
	MemoryRequest *string `json:"MemoryRequest,omitempty" xml:"MemoryRequest,omitempty"`
	// The timeout period of the change process. Valid values: 1 to 1800. Default value: 600. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateK8sApplicationConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sApplicationConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateK8sApplicationConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateK8sApplicationConfigRequest) GetCpuLimit() *string {
	return s.CpuLimit
}

func (s *UpdateK8sApplicationConfigRequest) GetCpuRequest() *string {
	return s.CpuRequest
}

func (s *UpdateK8sApplicationConfigRequest) GetEphemeralStorageLimit() *string {
	return s.EphemeralStorageLimit
}

func (s *UpdateK8sApplicationConfigRequest) GetEphemeralStorageRequest() *string {
	return s.EphemeralStorageRequest
}

func (s *UpdateK8sApplicationConfigRequest) GetMcpuLimit() *string {
	return s.McpuLimit
}

func (s *UpdateK8sApplicationConfigRequest) GetMcpuRequest() *string {
	return s.McpuRequest
}

func (s *UpdateK8sApplicationConfigRequest) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *UpdateK8sApplicationConfigRequest) GetMemoryRequest() *string {
	return s.MemoryRequest
}

func (s *UpdateK8sApplicationConfigRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateK8sApplicationConfigRequest) SetAppId(v string) *UpdateK8sApplicationConfigRequest {
	s.AppId = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetClusterId(v string) *UpdateK8sApplicationConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetCpuLimit(v string) *UpdateK8sApplicationConfigRequest {
	s.CpuLimit = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetCpuRequest(v string) *UpdateK8sApplicationConfigRequest {
	s.CpuRequest = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetEphemeralStorageLimit(v string) *UpdateK8sApplicationConfigRequest {
	s.EphemeralStorageLimit = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetEphemeralStorageRequest(v string) *UpdateK8sApplicationConfigRequest {
	s.EphemeralStorageRequest = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetMcpuLimit(v string) *UpdateK8sApplicationConfigRequest {
	s.McpuLimit = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetMcpuRequest(v string) *UpdateK8sApplicationConfigRequest {
	s.McpuRequest = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetMemoryLimit(v string) *UpdateK8sApplicationConfigRequest {
	s.MemoryLimit = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetMemoryRequest(v string) *UpdateK8sApplicationConfigRequest {
	s.MemoryRequest = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetTimeout(v int32) *UpdateK8sApplicationConfigRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) Validate() error {
	return dara.Validate(s)
}
