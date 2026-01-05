// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginClassesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPluginClassesResponseBody
	GetCode() *string
	SetData(v *ListPluginClassesResponseBodyData) *ListPluginClassesResponseBody
	GetData() *ListPluginClassesResponseBodyData
	SetMessage(v string) *ListPluginClassesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPluginClassesResponseBody
	GetRequestId() *string
}

type ListPluginClassesResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListPluginClassesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 61EBF577-1601-51E1-B136-9CD6xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPluginClassesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginClassesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginClassesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPluginClassesResponseBody) GetData() *ListPluginClassesResponseBodyData {
	return s.Data
}

func (s *ListPluginClassesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPluginClassesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginClassesResponseBody) SetCode(v string) *ListPluginClassesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPluginClassesResponseBody) SetData(v *ListPluginClassesResponseBodyData) *ListPluginClassesResponseBody {
	s.Data = v
	return s
}

func (s *ListPluginClassesResponseBody) SetMessage(v string) *ListPluginClassesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPluginClassesResponseBody) SetRequestId(v string) *ListPluginClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginClassesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPluginClassesResponseBodyData struct {
	Items []*ListPluginClassesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s ListPluginClassesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPluginClassesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPluginClassesResponseBodyData) GetItems() []*ListPluginClassesResponseBodyDataItems {
	return s.Items
}

func (s *ListPluginClassesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPluginClassesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPluginClassesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListPluginClassesResponseBodyData) SetItems(v []*ListPluginClassesResponseBodyDataItems) *ListPluginClassesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPluginClassesResponseBodyData) SetPageNumber(v int32) *ListPluginClassesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPluginClassesResponseBodyData) SetPageSize(v int32) *ListPluginClassesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPluginClassesResponseBodyData) SetTotalSize(v int32) *ListPluginClassesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListPluginClassesResponseBodyData) Validate() error {
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

type ListPluginClassesResponseBodyDataItems struct {
	Alias       *string `json:"alias,omitempty" xml:"alias,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	Installed *bool `json:"installed,omitempty" xml:"installed,omitempty"`
	// example:
	//
	// oauth
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pls-d4ghv6em1hkixxxxxxxx
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	// example:
	//
	// pl-cvu6r4um1hkoxxxxxxxx
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// example:
	//
	// Success
	PublishStatus *string `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	// example:
	//
	// HigressOfficial
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// 2.0.0
	SupportedMinGatewayVersion *string `json:"supportedMinGatewayVersion,omitempty" xml:"supportedMinGatewayVersion,omitempty"`
	// example:
	//
	// Auth
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListPluginClassesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListPluginClassesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListPluginClassesResponseBodyDataItems) GetAlias() *string {
	return s.Alias
}

func (s *ListPluginClassesResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListPluginClassesResponseBodyDataItems) GetInstalled() *bool {
	return s.Installed
}

func (s *ListPluginClassesResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListPluginClassesResponseBodyDataItems) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *ListPluginClassesResponseBodyDataItems) GetPluginId() *string {
	return s.PluginId
}

func (s *ListPluginClassesResponseBodyDataItems) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *ListPluginClassesResponseBodyDataItems) GetSource() *string {
	return s.Source
}

func (s *ListPluginClassesResponseBodyDataItems) GetSupportedMinGatewayVersion() *string {
	return s.SupportedMinGatewayVersion
}

func (s *ListPluginClassesResponseBodyDataItems) GetType() *string {
	return s.Type
}

func (s *ListPluginClassesResponseBodyDataItems) GetVersion() *string {
	return s.Version
}

func (s *ListPluginClassesResponseBodyDataItems) SetAlias(v string) *ListPluginClassesResponseBodyDataItems {
	s.Alias = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetDescription(v string) *ListPluginClassesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetInstalled(v bool) *ListPluginClassesResponseBodyDataItems {
	s.Installed = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetName(v string) *ListPluginClassesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetPluginClassId(v string) *ListPluginClassesResponseBodyDataItems {
	s.PluginClassId = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetPluginId(v string) *ListPluginClassesResponseBodyDataItems {
	s.PluginId = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetPublishStatus(v string) *ListPluginClassesResponseBodyDataItems {
	s.PublishStatus = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetSource(v string) *ListPluginClassesResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetSupportedMinGatewayVersion(v string) *ListPluginClassesResponseBodyDataItems {
	s.SupportedMinGatewayVersion = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetType(v string) *ListPluginClassesResponseBodyDataItems {
	s.Type = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) SetVersion(v string) *ListPluginClassesResponseBodyDataItems {
	s.Version = &v
	return s
}

func (s *ListPluginClassesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
