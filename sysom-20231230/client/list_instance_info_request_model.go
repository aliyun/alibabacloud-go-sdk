// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListInstanceInfoRequest
	GetXDebugId() *string
	SetInfoType(v string) *ListInstanceInfoRequest
	GetInfoType() *string
	SetInstanceType(v string) *ListInstanceInfoRequest
	GetInstanceType() *string
	SetManagedType(v string) *ListInstanceInfoRequest
	GetManagedType() *string
	SetMaxResults(v int32) *ListInstanceInfoRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceInfoRequest
	GetNextToken() *string
	SetPluginId(v string) *ListInstanceInfoRequest
	GetPluginId() *string
	SetRegion(v string) *ListInstanceInfoRequest
	GetRegion() *string
	SetXSysomInvokeSource(v string) *ListInstanceInfoRequest
	GetXSysomInvokeSource() *string
}

type ListInstanceInfoRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// example:
	//
	// instance_tag
	InfoType *string `json:"infoType,omitempty" xml:"infoType,omitempty"`
	// example:
	//
	// ecs
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// example:
	//
	// managed
	ManagedType *string `json:"managedType,omitempty" xml:"managedType,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xiF/wDgaeitjjhVJYYzLwJ4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 01fc4a0b-f199-4885-9861-b4054a310fe7
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region             *string `json:"region,omitempty" xml:"region,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListInstanceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceInfoRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListInstanceInfoRequest) GetInfoType() *string {
	return s.InfoType
}

func (s *ListInstanceInfoRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstanceInfoRequest) GetManagedType() *string {
	return s.ManagedType
}

func (s *ListInstanceInfoRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceInfoRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceInfoRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListInstanceInfoRequest) GetRegion() *string {
	return s.Region
}

func (s *ListInstanceInfoRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListInstanceInfoRequest) SetXDebugId(v string) *ListInstanceInfoRequest {
	s.XDebugId = &v
	return s
}

func (s *ListInstanceInfoRequest) SetInfoType(v string) *ListInstanceInfoRequest {
	s.InfoType = &v
	return s
}

func (s *ListInstanceInfoRequest) SetInstanceType(v string) *ListInstanceInfoRequest {
	s.InstanceType = &v
	return s
}

func (s *ListInstanceInfoRequest) SetManagedType(v string) *ListInstanceInfoRequest {
	s.ManagedType = &v
	return s
}

func (s *ListInstanceInfoRequest) SetMaxResults(v int32) *ListInstanceInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceInfoRequest) SetNextToken(v string) *ListInstanceInfoRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstanceInfoRequest) SetPluginId(v string) *ListInstanceInfoRequest {
	s.PluginId = &v
	return s
}

func (s *ListInstanceInfoRequest) SetRegion(v string) *ListInstanceInfoRequest {
	s.Region = &v
	return s
}

func (s *ListInstanceInfoRequest) SetXSysomInvokeSource(v string) *ListInstanceInfoRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListInstanceInfoRequest) Validate() error {
	return dara.Validate(s)
}
