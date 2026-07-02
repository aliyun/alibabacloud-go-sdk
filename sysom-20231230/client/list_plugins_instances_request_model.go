// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListPluginsInstancesRequest
	GetCurrent() *int64
	SetInstanceIdName(v string) *ListPluginsInstancesRequest
	GetInstanceIdName() *string
	SetInstanceTag(v string) *ListPluginsInstancesRequest
	GetInstanceTag() *string
	SetOperationType(v string) *ListPluginsInstancesRequest
	GetOperationType() *string
	SetPageSize(v int64) *ListPluginsInstancesRequest
	GetPageSize() *int64
	SetPluginId(v string) *ListPluginsInstancesRequest
	GetPluginId() *string
	SetRegion(v string) *ListPluginsInstancesRequest
	GetRegion() *string
}

type ListPluginsInstancesRequest struct {
	// The current page number. This field is present when pagination is used.
	//
	// example:
	//
	// 5
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// Filters instances by instance ID or instance name. Fuzzy match is supported.
	//
	// example:
	//
	// i-bpxx
	InstanceIdName *string `json:"instance_id_name,omitempty" xml:"instance_id_name,omitempty"`
	// Filters instances by instance tag.
	//
	// example:
	//
	// {"key":"app","value":"sysom-aliyun-com"}
	InstanceTag *string `json:"instance_tag,omitempty" xml:"instance_tag,omitempty"`
	// Filters instances by plug-in installation status.
	//
	// This parameter is required.
	//
	// example:
	//
	// install
	OperationType *string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters the instance list by the specified agent. If this parameter is specified, only instances associated with the specified agent are returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// Filters instances by region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s ListPluginsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListPluginsInstancesRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListPluginsInstancesRequest) GetInstanceIdName() *string {
	return s.InstanceIdName
}

func (s *ListPluginsInstancesRequest) GetInstanceTag() *string {
	return s.InstanceTag
}

func (s *ListPluginsInstancesRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListPluginsInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPluginsInstancesRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListPluginsInstancesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListPluginsInstancesRequest) SetCurrent(v int64) *ListPluginsInstancesRequest {
	s.Current = &v
	return s
}

func (s *ListPluginsInstancesRequest) SetInstanceIdName(v string) *ListPluginsInstancesRequest {
	s.InstanceIdName = &v
	return s
}

func (s *ListPluginsInstancesRequest) SetInstanceTag(v string) *ListPluginsInstancesRequest {
	s.InstanceTag = &v
	return s
}

func (s *ListPluginsInstancesRequest) SetOperationType(v string) *ListPluginsInstancesRequest {
	s.OperationType = &v
	return s
}

func (s *ListPluginsInstancesRequest) SetPageSize(v int64) *ListPluginsInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPluginsInstancesRequest) SetPluginId(v string) *ListPluginsInstancesRequest {
	s.PluginId = &v
	return s
}

func (s *ListPluginsInstancesRequest) SetRegion(v string) *ListPluginsInstancesRequest {
	s.Region = &v
	return s
}

func (s *ListPluginsInstancesRequest) Validate() error {
	return dara.Validate(s)
}
