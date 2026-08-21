// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetConfigByNameRequest
	GetXDebugId() *string
	SetConfigName(v string) *GetConfigByNameRequest
	GetConfigName() *string
	SetConfigType(v string) *GetConfigByNameRequest
	GetConfigType() *string
	SetEntityId(v string) *GetConfigByNameRequest
	GetEntityId() *string
	SetUseGlobalUid(v bool) *GetConfigByNameRequest
	GetUseGlobalUid() *bool
	SetVersionId(v int64) *GetConfigByNameRequest
	GetVersionId() *int64
	SetXSysomInvokeSource(v string) *GetConfigByNameRequest
	GetXSysomInvokeSource() *string
}

type GetConfigByNameRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The configuration name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 节点网络延时检测
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// The type of the configuration parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// attention
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// The entity ID. Default value: "default".
	//
	// example:
	//
	// i-bptest
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Specifies whether to use the global UID.
	//
	// example:
	//
	// false
	UseGlobalUid *bool `json:"useGlobalUid,omitempty" xml:"useGlobalUid,omitempty"`
	// The version ID.
	//
	// example:
	//
	// 1
	VersionId          *int64  `json:"versionId,omitempty" xml:"versionId,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetConfigByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigByNameRequest) GoString() string {
	return s.String()
}

func (s *GetConfigByNameRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetConfigByNameRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *GetConfigByNameRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetConfigByNameRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *GetConfigByNameRequest) GetUseGlobalUid() *bool {
	return s.UseGlobalUid
}

func (s *GetConfigByNameRequest) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetConfigByNameRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetConfigByNameRequest) SetXDebugId(v string) *GetConfigByNameRequest {
	s.XDebugId = &v
	return s
}

func (s *GetConfigByNameRequest) SetConfigName(v string) *GetConfigByNameRequest {
	s.ConfigName = &v
	return s
}

func (s *GetConfigByNameRequest) SetConfigType(v string) *GetConfigByNameRequest {
	s.ConfigType = &v
	return s
}

func (s *GetConfigByNameRequest) SetEntityId(v string) *GetConfigByNameRequest {
	s.EntityId = &v
	return s
}

func (s *GetConfigByNameRequest) SetUseGlobalUid(v bool) *GetConfigByNameRequest {
	s.UseGlobalUid = &v
	return s
}

func (s *GetConfigByNameRequest) SetVersionId(v int64) *GetConfigByNameRequest {
	s.VersionId = &v
	return s
}

func (s *GetConfigByNameRequest) SetXSysomInvokeSource(v string) *GetConfigByNameRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetConfigByNameRequest) Validate() error {
	return dara.Validate(s)
}
