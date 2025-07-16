// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormComponentDefinitionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResult(v []*GetFormComponentDefinitionListResponseBodyResult) *GetFormComponentDefinitionListResponseBody
	GetResult() []*GetFormComponentDefinitionListResponseBodyResult
	SetRequestId(v string) *GetFormComponentDefinitionListResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetFormComponentDefinitionListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetFormComponentDefinitionListResponseBody
	GetVendorType() *string
}

type GetFormComponentDefinitionListResponseBody struct {
	Result []*GetFormComponentDefinitionListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetFormComponentDefinitionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFormComponentDefinitionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListResponseBody) GetResult() []*GetFormComponentDefinitionListResponseBodyResult {
	return s.Result
}

func (s *GetFormComponentDefinitionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFormComponentDefinitionListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetFormComponentDefinitionListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetFormComponentDefinitionListResponseBody) SetResult(v []*GetFormComponentDefinitionListResponseBodyResult) *GetFormComponentDefinitionListResponseBody {
	s.Result = v
	return s
}

func (s *GetFormComponentDefinitionListResponseBody) SetRequestId(v string) *GetFormComponentDefinitionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBody) SetVendorRequestId(v string) *GetFormComponentDefinitionListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBody) SetVendorType(v string) *GetFormComponentDefinitionListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFormComponentDefinitionListResponseBodyResult struct {
	// example:
	//
	// FooterYida
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// formContainer_kksjiuk
	FieldId *string `json:"FieldId,omitempty" xml:"FieldId,omitempty"`
	// example:
	//
	// {"en_US":""}
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// formContainer_kksjiuk
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s GetFormComponentDefinitionListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFormComponentDefinitionListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListResponseBodyResult) GetComponentName() *string {
	return s.ComponentName
}

func (s *GetFormComponentDefinitionListResponseBodyResult) GetFieldId() *string {
	return s.FieldId
}

func (s *GetFormComponentDefinitionListResponseBodyResult) GetLabel() *string {
	return s.Label
}

func (s *GetFormComponentDefinitionListResponseBodyResult) GetParentId() *string {
	return s.ParentId
}

func (s *GetFormComponentDefinitionListResponseBodyResult) SetComponentName(v string) *GetFormComponentDefinitionListResponseBodyResult {
	s.ComponentName = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBodyResult) SetFieldId(v string) *GetFormComponentDefinitionListResponseBodyResult {
	s.FieldId = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBodyResult) SetLabel(v string) *GetFormComponentDefinitionListResponseBodyResult {
	s.Label = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBodyResult) SetParentId(v string) *GetFormComponentDefinitionListResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
