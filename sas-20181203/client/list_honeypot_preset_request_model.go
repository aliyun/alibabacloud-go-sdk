// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotPresetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListHoneypotPresetRequest
	GetCurrentPage() *int32
	SetHoneypotImageName(v string) *ListHoneypotPresetRequest
	GetHoneypotImageName() *string
	SetLang(v string) *ListHoneypotPresetRequest
	GetLang() *string
	SetNodeId(v string) *ListHoneypotPresetRequest
	GetNodeId() *string
	SetNodeName(v string) *ListHoneypotPresetRequest
	GetNodeName() *string
	SetPageSize(v int32) *ListHoneypotPresetRequest
	GetPageSize() *int32
	SetPresetName(v string) *ListHoneypotPresetRequest
	GetPresetName() *string
}

type ListHoneypotPresetRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the honeypot image.
	//
	// example:
	//
	// ruoyi
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// d892b4fe-af0d-4486-ab2a-8a518045****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the management node.
	//
	// example:
	//
	// Node1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The custom name of the honeypot template.
	//
	// example:
	//
	// mx-rouyi
	PresetName *string `json:"PresetName,omitempty" xml:"PresetName,omitempty"`
}

func (s ListHoneypotPresetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotPresetRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotPresetRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotPresetRequest) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *ListHoneypotPresetRequest) GetLang() *string {
	return s.Lang
}

func (s *ListHoneypotPresetRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListHoneypotPresetRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *ListHoneypotPresetRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotPresetRequest) GetPresetName() *string {
	return s.PresetName
}

func (s *ListHoneypotPresetRequest) SetCurrentPage(v int32) *ListHoneypotPresetRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotPresetRequest) SetHoneypotImageName(v string) *ListHoneypotPresetRequest {
	s.HoneypotImageName = &v
	return s
}

func (s *ListHoneypotPresetRequest) SetLang(v string) *ListHoneypotPresetRequest {
	s.Lang = &v
	return s
}

func (s *ListHoneypotPresetRequest) SetNodeId(v string) *ListHoneypotPresetRequest {
	s.NodeId = &v
	return s
}

func (s *ListHoneypotPresetRequest) SetNodeName(v string) *ListHoneypotPresetRequest {
	s.NodeName = &v
	return s
}

func (s *ListHoneypotPresetRequest) SetPageSize(v int32) *ListHoneypotPresetRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotPresetRequest) SetPresetName(v string) *ListHoneypotPresetRequest {
	s.PresetName = &v
	return s
}

func (s *ListHoneypotPresetRequest) Validate() error {
	return dara.Validate(s)
}
