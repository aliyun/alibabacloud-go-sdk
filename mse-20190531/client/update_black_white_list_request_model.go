// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBlackWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateBlackWhiteListRequest
	GetAcceptLanguage() *string
	SetContent(v string) *UpdateBlackWhiteListRequest
	GetContent() *string
	SetGatewayUniqueId(v string) *UpdateBlackWhiteListRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateBlackWhiteListRequest
	GetId() *int64
	SetIsWhite(v bool) *UpdateBlackWhiteListRequest
	GetIsWhite() *bool
	SetName(v string) *UpdateBlackWhiteListRequest
	GetName() *string
	SetNote(v string) *UpdateBlackWhiteListRequest
	GetNote() *string
	SetResourceIdJsonList(v string) *UpdateBlackWhiteListRequest
	GetResourceIdJsonList() *string
	SetResourceType(v string) *UpdateBlackWhiteListRequest
	GetResourceType() *string
	SetStatus(v string) *UpdateBlackWhiteListRequest
	GetStatus() *string
	SetType(v string) *UpdateBlackWhiteListRequest
	GetType() *string
}

type UpdateBlackWhiteListRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The content of the blacklist.
	//
	// example:
	//
	// 127.0.2.11
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-c70622ff52fe49beb29bea9a6f52****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the blacklist.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to enable the whitelist.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// The name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description.
	//
	// example:
	//
	// this is a note
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The resource IDs in the JSON format.
	//
	// example:
	//
	// [123]
	ResourceIdJsonList *string `json:"ResourceIdJsonList,omitempty" xml:"ResourceIdJsonList,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// GATEWAY
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Specifies whether to enable the blacklist or whitelist.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the blacklist or whitelist.
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateBlackWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateBlackWhiteListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateBlackWhiteListRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateBlackWhiteListRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateBlackWhiteListRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateBlackWhiteListRequest) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *UpdateBlackWhiteListRequest) GetName() *string {
	return s.Name
}

func (s *UpdateBlackWhiteListRequest) GetNote() *string {
	return s.Note
}

func (s *UpdateBlackWhiteListRequest) GetResourceIdJsonList() *string {
	return s.ResourceIdJsonList
}

func (s *UpdateBlackWhiteListRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateBlackWhiteListRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateBlackWhiteListRequest) GetType() *string {
	return s.Type
}

func (s *UpdateBlackWhiteListRequest) SetAcceptLanguage(v string) *UpdateBlackWhiteListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetContent(v string) *UpdateBlackWhiteListRequest {
	s.Content = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetGatewayUniqueId(v string) *UpdateBlackWhiteListRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetId(v int64) *UpdateBlackWhiteListRequest {
	s.Id = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetIsWhite(v bool) *UpdateBlackWhiteListRequest {
	s.IsWhite = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetName(v string) *UpdateBlackWhiteListRequest {
	s.Name = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetNote(v string) *UpdateBlackWhiteListRequest {
	s.Note = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetResourceIdJsonList(v string) *UpdateBlackWhiteListRequest {
	s.ResourceIdJsonList = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetResourceType(v string) *UpdateBlackWhiteListRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetStatus(v string) *UpdateBlackWhiteListRequest {
	s.Status = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) SetType(v string) *UpdateBlackWhiteListRequest {
	s.Type = &v
	return s
}

func (s *UpdateBlackWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
