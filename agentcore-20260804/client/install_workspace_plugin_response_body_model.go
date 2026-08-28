// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallWorkspacePluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallWorkspacePluginResponseBody
	GetCode() *string
	SetData(v *InstallWorkspacePluginResponseBodyData) *InstallWorkspacePluginResponseBody
	GetData() *InstallWorkspacePluginResponseBodyData
	SetHttpStatusCode(v int32) *InstallWorkspacePluginResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InstallWorkspacePluginResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallWorkspacePluginResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallWorkspacePluginResponseBody
	GetSuccess() *bool
}

type InstallWorkspacePluginResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the plugin installation operation.
	Data *InstallWorkspacePluginResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InstallWorkspacePluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallWorkspacePluginResponseBody) GetData() *InstallWorkspacePluginResponseBodyData {
	return s.Data
}

func (s *InstallWorkspacePluginResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InstallWorkspacePluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallWorkspacePluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallWorkspacePluginResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallWorkspacePluginResponseBody) SetCode(v string) *InstallWorkspacePluginResponseBody {
	s.Code = &v
	return s
}

func (s *InstallWorkspacePluginResponseBody) SetData(v *InstallWorkspacePluginResponseBodyData) *InstallWorkspacePluginResponseBody {
	s.Data = v
	return s
}

func (s *InstallWorkspacePluginResponseBody) SetHttpStatusCode(v int32) *InstallWorkspacePluginResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InstallWorkspacePluginResponseBody) SetMessage(v string) *InstallWorkspacePluginResponseBody {
	s.Message = &v
	return s
}

func (s *InstallWorkspacePluginResponseBody) SetRequestId(v string) *InstallWorkspacePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallWorkspacePluginResponseBody) SetSuccess(v bool) *InstallWorkspacePluginResponseBody {
	s.Success = &v
	return s
}

func (s *InstallWorkspacePluginResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallWorkspacePluginResponseBodyData struct {
	// Indicates whether the plugin is enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The plugin name.
	//
	// example:
	//
	// collaboration
	PluginName *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
	// The plugin status. Valid values: DISABLED, ENABLING, ENABLED, ENABLE_FAILED, DISABLING, and DISABLE_FAILED.
	//
	// example:
	//
	// ENABLING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s InstallWorkspacePluginResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallWorkspacePluginResponseBodyData) GetPluginName() *string {
	return s.PluginName
}

func (s *InstallWorkspacePluginResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *InstallWorkspacePluginResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *InstallWorkspacePluginResponseBodyData) SetEnabled(v bool) *InstallWorkspacePluginResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *InstallWorkspacePluginResponseBodyData) SetPluginName(v string) *InstallWorkspacePluginResponseBodyData {
	s.PluginName = &v
	return s
}

func (s *InstallWorkspacePluginResponseBodyData) SetStatus(v string) *InstallWorkspacePluginResponseBodyData {
	s.Status = &v
	return s
}

func (s *InstallWorkspacePluginResponseBodyData) SetWorkspaceId(v string) *InstallWorkspacePluginResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *InstallWorkspacePluginResponseBodyData) Validate() error {
	return dara.Validate(s)
}
