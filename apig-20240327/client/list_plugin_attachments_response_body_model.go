// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPluginAttachmentsResponseBody
	GetCode() *string
	SetData(v *ListPluginAttachmentsResponseBodyData) *ListPluginAttachmentsResponseBody
	GetData() *ListPluginAttachmentsResponseBodyData
	SetMessage(v string) *ListPluginAttachmentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPluginAttachmentsResponseBody
	GetRequestId() *string
}

type ListPluginAttachmentsResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListPluginAttachmentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9640D776-794A-5077-9184-A247CA4B45C1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPluginAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginAttachmentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPluginAttachmentsResponseBody) GetData() *ListPluginAttachmentsResponseBodyData {
	return s.Data
}

func (s *ListPluginAttachmentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPluginAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginAttachmentsResponseBody) SetCode(v string) *ListPluginAttachmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPluginAttachmentsResponseBody) SetData(v *ListPluginAttachmentsResponseBodyData) *ListPluginAttachmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListPluginAttachmentsResponseBody) SetMessage(v string) *ListPluginAttachmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPluginAttachmentsResponseBody) SetRequestId(v string) *ListPluginAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginAttachmentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPluginAttachmentsResponseBodyData struct {
	Items []*ListPluginAttachmentsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListPluginAttachmentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPluginAttachmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPluginAttachmentsResponseBodyData) GetItems() []*ListPluginAttachmentsResponseBodyDataItems {
	return s.Items
}

func (s *ListPluginAttachmentsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPluginAttachmentsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPluginAttachmentsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListPluginAttachmentsResponseBodyData) SetItems(v []*ListPluginAttachmentsResponseBodyDataItems) *ListPluginAttachmentsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPluginAttachmentsResponseBodyData) SetPageNumber(v int32) *ListPluginAttachmentsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPluginAttachmentsResponseBodyData) SetPageSize(v int32) *ListPluginAttachmentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPluginAttachmentsResponseBodyData) SetTotalSize(v int32) *ListPluginAttachmentsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListPluginAttachmentsResponseBodyData) Validate() error {
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

type ListPluginAttachmentsResponseBodyDataItems struct {
	// example:
	//
	// GatewayRoute
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// example:
	//
	// true
	Enable             *bool               `json:"enable,omitempty" xml:"enable,omitempty"`
	EnvironmentInfo    *EnvironmentInfo    `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty"`
	ParentResourceInfo *ParentResourceInfo `json:"parentResourceInfo,omitempty" xml:"parentResourceInfo,omitempty"`
	// example:
	//
	// pa-d0j9t5em1hkncrlo51mg
	PluginAttachmentId *string          `json:"pluginAttachmentId,omitempty" xml:"pluginAttachmentId,omitempty"`
	PluginClassInfo    *PluginClassInfo `json:"pluginClassInfo,omitempty" xml:"pluginClassInfo,omitempty"`
	// example:
	//
	// bGltaXRfYnlfaGVhZGVyOiB4LWFwaS1rZXkKbGltaXRfa2V5czoKLSBrZXk6IGV4YW1wbGUta2V5LWEKICBxdWVyeV9wZXJfc2Vjb25kOiAxMAotIGtleTogZXhhbXBsZS1rZXktYgogIHF1ZXJ5X3Blcl9zZWNvbmQ6IDEK
	PluginConfig *string `json:"pluginConfig,omitempty" xml:"pluginConfig,omitempty"`
	// example:
	//
	// pl-cvu6r4um1hko3b3ti0a0
	PluginId      *string         `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ResourceInfos []*ResourceInfo `json:"resourceInfos,omitempty" xml:"resourceInfos,omitempty" type:"Repeated"`
}

func (s ListPluginAttachmentsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListPluginAttachmentsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetEnable() *bool {
	return s.Enable
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetEnvironmentInfo() *EnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetParentResourceInfo() *ParentResourceInfo {
	return s.ParentResourceInfo
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetPluginAttachmentId() *string {
	return s.PluginAttachmentId
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetPluginClassInfo() *PluginClassInfo {
	return s.PluginClassInfo
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetPluginId() *string {
	return s.PluginId
}

func (s *ListPluginAttachmentsResponseBodyDataItems) GetResourceInfos() []*ResourceInfo {
	return s.ResourceInfos
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetAttachResourceType(v string) *ListPluginAttachmentsResponseBodyDataItems {
	s.AttachResourceType = &v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetEnable(v bool) *ListPluginAttachmentsResponseBodyDataItems {
	s.Enable = &v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetEnvironmentInfo(v *EnvironmentInfo) *ListPluginAttachmentsResponseBodyDataItems {
	s.EnvironmentInfo = v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetParentResourceInfo(v *ParentResourceInfo) *ListPluginAttachmentsResponseBodyDataItems {
	s.ParentResourceInfo = v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetPluginAttachmentId(v string) *ListPluginAttachmentsResponseBodyDataItems {
	s.PluginAttachmentId = &v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetPluginClassInfo(v *PluginClassInfo) *ListPluginAttachmentsResponseBodyDataItems {
	s.PluginClassInfo = v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetPluginConfig(v string) *ListPluginAttachmentsResponseBodyDataItems {
	s.PluginConfig = &v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetPluginId(v string) *ListPluginAttachmentsResponseBodyDataItems {
	s.PluginId = &v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) SetResourceInfos(v []*ResourceInfo) *ListPluginAttachmentsResponseBodyDataItems {
	s.ResourceInfos = v
	return s
}

func (s *ListPluginAttachmentsResponseBodyDataItems) Validate() error {
	if s.EnvironmentInfo != nil {
		if err := s.EnvironmentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ParentResourceInfo != nil {
		if err := s.ParentResourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PluginClassInfo != nil {
		if err := s.PluginClassInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceInfos != nil {
		for _, item := range s.ResourceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
