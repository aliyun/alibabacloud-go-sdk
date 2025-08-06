// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonVariableList(v []*GetCardTemplateResponseBodyCommonVariableList) *GetCardTemplateResponseBody
	GetCommonVariableList() []*GetCardTemplateResponseBodyCommonVariableList
	SetRequestId(v string) *GetCardTemplateResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetCardTemplateResponseBody
	GetStatus() *string
	SetTemplateId(v string) *GetCardTemplateResponseBody
	GetTemplateId() *string
	SetVendorRequestId(v string) *GetCardTemplateResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetCardTemplateResponseBody
	GetVendorType() *string
}

type GetCardTemplateResponseBody struct {
	// example:
	//
	// []
	CommonVariableList []*GetCardTemplateResponseBodyCommonVariableList `json:"commonVariableList,omitempty" xml:"commonVariableList,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 123
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetCardTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardTemplateResponseBody) GetCommonVariableList() []*GetCardTemplateResponseBodyCommonVariableList {
	return s.CommonVariableList
}

func (s *GetCardTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCardTemplateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetCardTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetCardTemplateResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetCardTemplateResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetCardTemplateResponseBody) SetCommonVariableList(v []*GetCardTemplateResponseBodyCommonVariableList) *GetCardTemplateResponseBody {
	s.CommonVariableList = v
	return s
}

func (s *GetCardTemplateResponseBody) SetRequestId(v string) *GetCardTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardTemplateResponseBody) SetStatus(v string) *GetCardTemplateResponseBody {
	s.Status = &v
	return s
}

func (s *GetCardTemplateResponseBody) SetTemplateId(v string) *GetCardTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetCardTemplateResponseBody) SetVendorRequestId(v string) *GetCardTemplateResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetCardTemplateResponseBody) SetVendorType(v string) *GetCardTemplateResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetCardTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCardTemplateResponseBodyCommonVariableList struct {
	// example:
	//
	// 卡片摘要
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Id
	//
	// example:
	//
	// lastMessage
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IfPrivateFiled *bool `json:"IfPrivateFiled,omitempty" xml:"IfPrivateFiled,omitempty"`
	// example:
	//
	// lastMessage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCardTemplateResponseBodyCommonVariableList) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateResponseBodyCommonVariableList) GoString() string {
	return s.String()
}

func (s *GetCardTemplateResponseBodyCommonVariableList) GetDescription() *string {
	return s.Description
}

func (s *GetCardTemplateResponseBodyCommonVariableList) GetId() *string {
	return s.Id
}

func (s *GetCardTemplateResponseBodyCommonVariableList) GetIfPrivateFiled() *bool {
	return s.IfPrivateFiled
}

func (s *GetCardTemplateResponseBodyCommonVariableList) GetName() *string {
	return s.Name
}

func (s *GetCardTemplateResponseBodyCommonVariableList) GetType() *string {
	return s.Type
}

func (s *GetCardTemplateResponseBodyCommonVariableList) SetDescription(v string) *GetCardTemplateResponseBodyCommonVariableList {
	s.Description = &v
	return s
}

func (s *GetCardTemplateResponseBodyCommonVariableList) SetId(v string) *GetCardTemplateResponseBodyCommonVariableList {
	s.Id = &v
	return s
}

func (s *GetCardTemplateResponseBodyCommonVariableList) SetIfPrivateFiled(v bool) *GetCardTemplateResponseBodyCommonVariableList {
	s.IfPrivateFiled = &v
	return s
}

func (s *GetCardTemplateResponseBodyCommonVariableList) SetName(v string) *GetCardTemplateResponseBodyCommonVariableList {
	s.Name = &v
	return s
}

func (s *GetCardTemplateResponseBodyCommonVariableList) SetType(v string) *GetCardTemplateResponseBodyCommonVariableList {
	s.Type = &v
	return s
}

func (s *GetCardTemplateResponseBodyCommonVariableList) Validate() error {
	return dara.Validate(s)
}
