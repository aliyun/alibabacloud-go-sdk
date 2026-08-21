// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetResourcesRequest
	GetXDebugId() *string
	SetCluster(v string) *GetResourcesRequest
	GetCluster() *string
	SetInstance(v string) *GetResourcesRequest
	GetInstance() *string
	SetType(v string) *GetResourcesRequest
	GetType() *string
	SetXSysomInvokeSource(v string) *GetResourcesRequest
	GetXSysomInvokeSource() *string
}

type GetResourcesRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The resource type.
	//
	// example:
	//
	// mem
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetResourcesRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetResourcesRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetResourcesRequest) GetType() *string {
	return s.Type
}

func (s *GetResourcesRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetResourcesRequest) SetXDebugId(v string) *GetResourcesRequest {
	s.XDebugId = &v
	return s
}

func (s *GetResourcesRequest) SetCluster(v string) *GetResourcesRequest {
	s.Cluster = &v
	return s
}

func (s *GetResourcesRequest) SetInstance(v string) *GetResourcesRequest {
	s.Instance = &v
	return s
}

func (s *GetResourcesRequest) SetType(v string) *GetResourcesRequest {
	s.Type = &v
	return s
}

func (s *GetResourcesRequest) SetXSysomInvokeSource(v string) *GetResourcesRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetResourcesRequest) Validate() error {
	return dara.Validate(s)
}
