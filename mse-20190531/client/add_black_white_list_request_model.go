// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlackWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddBlackWhiteListRequest
	GetAcceptLanguage() *string
	SetContent(v string) *AddBlackWhiteListRequest
	GetContent() *string
	SetGatewayUniqueId(v string) *AddBlackWhiteListRequest
	GetGatewayUniqueId() *string
	SetIsWhite(v bool) *AddBlackWhiteListRequest
	GetIsWhite() *bool
	SetName(v string) *AddBlackWhiteListRequest
	GetName() *string
	SetNote(v string) *AddBlackWhiteListRequest
	GetNote() *string
	SetResourceIdJsonList(v string) *AddBlackWhiteListRequest
	GetResourceIdJsonList() *string
	SetResourceType(v string) *AddBlackWhiteListRequest
	GetResourceType() *string
	SetStatus(v string) *AddBlackWhiteListRequest
	GetStatus() *string
	SetType(v string) *AddBlackWhiteListRequest
	GetType() *string
}

type AddBlackWhiteListRequest struct {
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
	// 1.117.115.51
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-0fe488252dc44d55a9dd57875193****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The whitelist. Default value: No.
	//
	// example:
	//
	// false
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
	// The effective scope of the blacklist or whitelist. Valid values:
	//
	// 	- GATEWAY
	//
	// 	- DOMAIN
	//
	// 	- ROUTE
	//
	// example:
	//
	// GATEWAY
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the blacklist.
	//
	// 	- on: enabled
	//
	// 	- off: disabled
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of object in the blacklist or whitelist.
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddBlackWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddBlackWhiteListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddBlackWhiteListRequest) GetContent() *string {
	return s.Content
}

func (s *AddBlackWhiteListRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddBlackWhiteListRequest) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *AddBlackWhiteListRequest) GetName() *string {
	return s.Name
}

func (s *AddBlackWhiteListRequest) GetNote() *string {
	return s.Note
}

func (s *AddBlackWhiteListRequest) GetResourceIdJsonList() *string {
	return s.ResourceIdJsonList
}

func (s *AddBlackWhiteListRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddBlackWhiteListRequest) GetStatus() *string {
	return s.Status
}

func (s *AddBlackWhiteListRequest) GetType() *string {
	return s.Type
}

func (s *AddBlackWhiteListRequest) SetAcceptLanguage(v string) *AddBlackWhiteListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetContent(v string) *AddBlackWhiteListRequest {
	s.Content = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetGatewayUniqueId(v string) *AddBlackWhiteListRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetIsWhite(v bool) *AddBlackWhiteListRequest {
	s.IsWhite = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetName(v string) *AddBlackWhiteListRequest {
	s.Name = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetNote(v string) *AddBlackWhiteListRequest {
	s.Note = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetResourceIdJsonList(v string) *AddBlackWhiteListRequest {
	s.ResourceIdJsonList = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetResourceType(v string) *AddBlackWhiteListRequest {
	s.ResourceType = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetStatus(v string) *AddBlackWhiteListRequest {
	s.Status = &v
	return s
}

func (s *AddBlackWhiteListRequest) SetType(v string) *AddBlackWhiteListRequest {
	s.Type = &v
	return s
}

func (s *AddBlackWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
