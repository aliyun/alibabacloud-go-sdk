// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncPluginsResponseBody
	GetCode() *string
	SetData(v []*ListContextDatabaseSyncPluginsResponseBodyData) *ListContextDatabaseSyncPluginsResponseBody
	GetData() []*ListContextDatabaseSyncPluginsResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncPluginsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncPluginsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncPluginsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncPluginsResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncPluginsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query results.
	Data []*ListContextDatabaseSyncPluginsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContextDatabaseSyncPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncPluginsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncPluginsResponseBody) GetData() []*ListContextDatabaseSyncPluginsResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncPluginsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncPluginsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncPluginsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncPluginsResponseBody) SetCode(v string) *ListContextDatabaseSyncPluginsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBody) SetData(v []*ListContextDatabaseSyncPluginsResponseBodyData) *ListContextDatabaseSyncPluginsResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncPluginsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBody) SetMessage(v string) *ListContextDatabaseSyncPluginsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBody) SetRequestId(v string) *ListContextDatabaseSyncPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncPluginsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncPluginsResponseBodyData struct {
	// The authorization method.
	//
	// example:
	//
	// {"kind":"oauth","phases":["discovery"]}
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// The display name.
	//
	// example:
	//
	// R&D team DingTalk data source
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The icon.
	//
	// example:
	//
	// /icons/dingtalk.svg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The plugin ID.
	//
	// example:
	//
	// dingtalk
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The plugin type. This operation returns only plugins of the source type.
	//
	// example:
	//
	// source
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The datasource config template (JSON character string, structure varies depending on the plugin).
	//
	// example:
	//
	// {"hint":"Select import scope","fields":[{"key":"folderIds","type":"string-list","label":"Folder / document scope"}]}
	SourceConfig *string `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty"`
	// The version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListContextDatabaseSyncPluginsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncPluginsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) GetIcon() *string {
	return s.Icon
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) GetKind() *string {
	return s.Kind
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) GetSourceConfig() *string {
	return s.SourceConfig
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) SetAuthorization(v string) *ListContextDatabaseSyncPluginsResponseBodyData {
	s.Authorization = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) SetDisplayName(v string) *ListContextDatabaseSyncPluginsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) SetIcon(v string) *ListContextDatabaseSyncPluginsResponseBodyData {
	s.Icon = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) SetId(v string) *ListContextDatabaseSyncPluginsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) SetKind(v string) *ListContextDatabaseSyncPluginsResponseBodyData {
	s.Kind = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) SetSourceConfig(v string) *ListContextDatabaseSyncPluginsResponseBodyData {
	s.SourceConfig = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) SetVersion(v string) *ListContextDatabaseSyncPluginsResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
