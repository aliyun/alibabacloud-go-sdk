// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPluginsResponseBody
	GetCode() *string
	SetData(v *ListPluginsResponseBodyData) *ListPluginsResponseBody
	GetData() *ListPluginsResponseBodyData
	SetMessage(v string) *ListPluginsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPluginsResponseBody
	GetRequestId() *string
}

type ListPluginsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response payload.
	Data *ListPluginsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 168BA42D-F822-569D-A67F-FC59E6ABC2B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPluginsResponseBody) GetData() *ListPluginsResponseBodyData {
	return s.Data
}

func (s *ListPluginsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginsResponseBody) SetCode(v string) *ListPluginsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPluginsResponseBody) SetData(v *ListPluginsResponseBodyData) *ListPluginsResponseBody {
	s.Data = v
	return s
}

func (s *ListPluginsResponseBody) SetMessage(v string) *ListPluginsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPluginsResponseBody) SetRequestId(v string) *ListPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPluginsResponseBodyData struct {
	// The list of plug-in information.
	Items []*ListPluginsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListPluginsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyData) GetItems() []*ListPluginsResponseBodyDataItems {
	return s.Items
}

func (s *ListPluginsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPluginsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPluginsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListPluginsResponseBodyData) SetItems(v []*ListPluginsResponseBodyDataItems) *ListPluginsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPluginsResponseBodyData) SetPageNumber(v int32) *ListPluginsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPluginsResponseBodyData) SetPageSize(v int32) *ListPluginsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPluginsResponseBodyData) SetTotalSize(v int32) *ListPluginsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListPluginsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPluginsResponseBodyDataItems struct {
	// The attachment information.
	AttachmentInfo *ListPluginsResponseBodyDataItemsAttachmentInfo `json:"attachmentInfo,omitempty" xml:"attachmentInfo,omitempty" type:"Struct"`
	// The gateway instance information.
	GatewayInfo *ListPluginsResponseBodyDataItemsGatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	// The plug-in type information.
	PluginClassInfo *ListPluginsResponseBodyDataItemsPluginClassInfo `json:"pluginClassInfo,omitempty" xml:"pluginClassInfo,omitempty" type:"Struct"`
	// The plug-in ID.
	//
	// example:
	//
	// pl-cvu6r4um1hko3b3ti0a0
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s ListPluginsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyDataItems) GetAttachmentInfo() *ListPluginsResponseBodyDataItemsAttachmentInfo {
	return s.AttachmentInfo
}

func (s *ListPluginsResponseBodyDataItems) GetGatewayInfo() *ListPluginsResponseBodyDataItemsGatewayInfo {
	return s.GatewayInfo
}

func (s *ListPluginsResponseBodyDataItems) GetPluginClassInfo() *ListPluginsResponseBodyDataItemsPluginClassInfo {
	return s.PluginClassInfo
}

func (s *ListPluginsResponseBodyDataItems) GetPluginId() *string {
	return s.PluginId
}

func (s *ListPluginsResponseBodyDataItems) SetAttachmentInfo(v *ListPluginsResponseBodyDataItemsAttachmentInfo) *ListPluginsResponseBodyDataItems {
	s.AttachmentInfo = v
	return s
}

func (s *ListPluginsResponseBodyDataItems) SetGatewayInfo(v *ListPluginsResponseBodyDataItemsGatewayInfo) *ListPluginsResponseBodyDataItems {
	s.GatewayInfo = v
	return s
}

func (s *ListPluginsResponseBodyDataItems) SetPluginClassInfo(v *ListPluginsResponseBodyDataItemsPluginClassInfo) *ListPluginsResponseBodyDataItems {
	s.PluginClassInfo = v
	return s
}

func (s *ListPluginsResponseBodyDataItems) SetPluginId(v string) *ListPluginsResponseBodyDataItems {
	s.PluginId = &v
	return s
}

func (s *ListPluginsResponseBodyDataItems) Validate() error {
	if s.AttachmentInfo != nil {
		if err := s.AttachmentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.GatewayInfo != nil {
		if err := s.GatewayInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PluginClassInfo != nil {
		if err := s.PluginClassInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPluginsResponseBodyDataItemsAttachmentInfo struct {
	// Indicates if enabled.
	//
	// example:
	//
	// false
	Enable *string `json:"enable,omitempty" xml:"enable,omitempty"`
	// The attachment ID.
	//
	// example:
	//
	// pa-ct2irn6m1hkreaen0t40
	PluginAttachmentId *string `json:"pluginAttachmentId,omitempty" xml:"pluginAttachmentId,omitempty"`
}

func (s ListPluginsResponseBodyDataItemsAttachmentInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBodyDataItemsAttachmentInfo) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyDataItemsAttachmentInfo) GetEnable() *string {
	return s.Enable
}

func (s *ListPluginsResponseBodyDataItemsAttachmentInfo) GetPluginAttachmentId() *string {
	return s.PluginAttachmentId
}

func (s *ListPluginsResponseBodyDataItemsAttachmentInfo) SetEnable(v string) *ListPluginsResponseBodyDataItemsAttachmentInfo {
	s.Enable = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsAttachmentInfo) SetPluginAttachmentId(v string) *ListPluginsResponseBodyDataItemsAttachmentInfo {
	s.PluginAttachmentId = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsAttachmentInfo) Validate() error {
	return dara.Validate(s)
}

type ListPluginsResponseBodyDataItemsGatewayInfo struct {
	// The instance ID.
	//
	// example:
	//
	// gw-cq7og15lhtxx6qasrj60
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// apitest-gw
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPluginsResponseBodyDataItemsGatewayInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBodyDataItemsGatewayInfo) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyDataItemsGatewayInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListPluginsResponseBodyDataItemsGatewayInfo) GetName() *string {
	return s.Name
}

func (s *ListPluginsResponseBodyDataItemsGatewayInfo) SetGatewayId(v string) *ListPluginsResponseBodyDataItemsGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsGatewayInfo) SetName(v string) *ListPluginsResponseBodyDataItemsGatewayInfo {
	s.Name = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsGatewayInfo) Validate() error {
	return dara.Validate(s)
}

type ListPluginsResponseBodyDataItemsPluginClassInfo struct {
	// The alias.
	//
	// example:
	//
	// key-based rate limit
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The execution priority.
	//
	// example:
	//
	// 999
	ExecutePriority *string `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	// The execution stage.
	//
	// example:
	//
	// AUTHZ
	ExecuteStage *string `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// key-rate-limit
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The plug-in type ID.
	//
	// example:
	//
	// pls-cqebrgh46ppatmpri
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	// The source of the plug-in.
	//
	// example:
	//
	// HigressOfficial
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The version.
	//
	// example:
	//
	// 2.0.3
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The description of the version.
	//
	// example:
	//
	// version description
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
}

func (s ListPluginsResponseBodyDataItemsPluginClassInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBodyDataItemsPluginClassInfo) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) GetAlias() *string {
	return s.Alias
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) GetExecutePriority() *string {
	return s.ExecutePriority
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) GetExecuteStage() *string {
	return s.ExecuteStage
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) GetName() *string {
	return s.Name
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) GetSource() *string {
	return s.Source
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) GetVersion() *string {
	return s.Version
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) SetAlias(v string) *ListPluginsResponseBodyDataItemsPluginClassInfo {
	s.Alias = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) SetExecutePriority(v string) *ListPluginsResponseBodyDataItemsPluginClassInfo {
	s.ExecutePriority = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) SetExecuteStage(v string) *ListPluginsResponseBodyDataItemsPluginClassInfo {
	s.ExecuteStage = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) SetName(v string) *ListPluginsResponseBodyDataItemsPluginClassInfo {
	s.Name = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) SetPluginClassId(v string) *ListPluginsResponseBodyDataItemsPluginClassInfo {
	s.PluginClassId = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) SetSource(v string) *ListPluginsResponseBodyDataItemsPluginClassInfo {
	s.Source = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) SetVersion(v string) *ListPluginsResponseBodyDataItemsPluginClassInfo {
	s.Version = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) SetVersionDescription(v string) *ListPluginsResponseBodyDataItemsPluginClassInfo {
	s.VersionDescription = &v
	return s
}

func (s *ListPluginsResponseBodyDataItemsPluginClassInfo) Validate() error {
	return dara.Validate(s)
}
