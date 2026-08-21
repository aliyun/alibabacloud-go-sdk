// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPodsOfInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListPodsOfInstanceRequest
	GetXDebugId() *string
	SetClusterId(v string) *ListPodsOfInstanceRequest
	GetClusterId() *string
	SetCurrent(v int64) *ListPodsOfInstanceRequest
	GetCurrent() *int64
	SetInstance(v string) *ListPodsOfInstanceRequest
	GetInstance() *string
	SetPageSize(v int64) *ListPodsOfInstanceRequest
	GetPageSize() *int64
	SetXSysomInvokeSource(v string) *ListPodsOfInstanceRequest
	GetXSysomInvokeSource() *string
}

type ListPodsOfInstanceRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c96e34d74eb6748f3b2a46552d5d653f6
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The current page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize           *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListPodsOfInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPodsOfInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListPodsOfInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPodsOfInstanceRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListPodsOfInstanceRequest) GetInstance() *string {
	return s.Instance
}

func (s *ListPodsOfInstanceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPodsOfInstanceRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListPodsOfInstanceRequest) SetXDebugId(v string) *ListPodsOfInstanceRequest {
	s.XDebugId = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetClusterId(v string) *ListPodsOfInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetCurrent(v int64) *ListPodsOfInstanceRequest {
	s.Current = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetInstance(v string) *ListPodsOfInstanceRequest {
	s.Instance = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetPageSize(v int64) *ListPodsOfInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetXSysomInvokeSource(v string) *ListPodsOfInstanceRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListPodsOfInstanceRequest) Validate() error {
	return dara.Validate(s)
}
