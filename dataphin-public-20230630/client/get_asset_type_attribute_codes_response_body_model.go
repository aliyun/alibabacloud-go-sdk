// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetTypeAttributeCodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAssetTypeAttributeCodesResponseBody
	GetCode() *string
	SetData(v []*GetAssetTypeAttributeCodesResponseBodyData) *GetAssetTypeAttributeCodesResponseBody
	GetData() []*GetAssetTypeAttributeCodesResponseBodyData
	SetHttpStatusCode(v int32) *GetAssetTypeAttributeCodesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAssetTypeAttributeCodesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAssetTypeAttributeCodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAssetTypeAttributeCodesResponseBody
	GetSuccess() *bool
}

type GetAssetTypeAttributeCodesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of property definitions.
	Data []*GetAssetTypeAttributeCodesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAssetTypeAttributeCodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetTypeAttributeCodesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetTypeAttributeCodesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAssetTypeAttributeCodesResponseBody) GetData() []*GetAssetTypeAttributeCodesResponseBodyData {
	return s.Data
}

func (s *GetAssetTypeAttributeCodesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAssetTypeAttributeCodesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAssetTypeAttributeCodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetTypeAttributeCodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAssetTypeAttributeCodesResponseBody) SetCode(v string) *GetAssetTypeAttributeCodesResponseBody {
	s.Code = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBody) SetData(v []*GetAssetTypeAttributeCodesResponseBodyData) *GetAssetTypeAttributeCodesResponseBody {
	s.Data = v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBody) SetHttpStatusCode(v int32) *GetAssetTypeAttributeCodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBody) SetMessage(v string) *GetAssetTypeAttributeCodesResponseBody {
	s.Message = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBody) SetRequestId(v string) *GetAssetTypeAttributeCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBody) SetSuccess(v bool) *GetAssetTypeAttributeCodesResponseBody {
	s.Success = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBody) Validate() error {
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

type GetAssetTypeAttributeCodesResponseBodyData struct {
	// The property code (unique identifier).
	//
	// example:
	//
	// data_level
	AttributeCode *string `json:"AttributeCode,omitempty" xml:"AttributeCode,omitempty"`
	// The property name (display name).
	//
	// example:
	//
	// Data level
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// The property source. Valid values:
	//
	// - SYSTEM: system preset.
	//
	// - CUSTOM: custom.
	//
	// example:
	//
	// CUSTOM
	AttributeSource *string `json:"AttributeSource,omitempty" xml:"AttributeSource,omitempty"`
	// The property type. Valid values:
	//
	// - MANAGEMENT: management property.
	//
	// - TECHNICAL: technical property.
	//
	// - BUSINESS: business property.
	//
	// example:
	//
	// MANAGEMENT
	AttributeType *string `json:"AttributeType,omitempty" xml:"AttributeType,omitempty"`
	// The property description.
	//
	// example:
	//
	// Data asset level classification
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The location where the property can be edited. Valid values:
	//
	// - ASSET: asset catalog.
	//
	// - DEVELOPMENT: development.
	EditableIn []*string `json:"EditableIn,omitempty" xml:"EditableIn,omitempty" type:"Repeated"`
	// The source of dropdown options. Valid values:
	//
	// - MANUAL: manual input.
	//
	// - SYSTEM_REFERENCE: reference to a system property.
	//
	// example:
	//
	// MANUAL
	EnumSourceType *string `json:"EnumSourceType,omitempty" xml:"EnumSourceType,omitempty"`
	// The list of dropdown options. This parameter has a value only when EnumSourceType is set to MANUAL.
	EnumValues []*GetAssetTypeAttributeCodesResponseBodyDataEnumValues `json:"EnumValues,omitempty" xml:"EnumValues,omitempty" type:"Repeated"`
	// The input mode. Valid values:
	//
	// - CUSTOM_INPUT: custom input.
	//
	// - DROPDOWN_SINGLE: single-select dropdown.
	//
	// - DROPDOWN_MULTI: multi-select dropdown.
	//
	// - HYPERLINK: hyperlink.
	//
	// example:
	//
	// DROPDOWN_SINGLE
	InputMode *string `json:"InputMode,omitempty" xml:"InputMode,omitempty"`
	// The hyperlink navigation method. This parameter has a value only when InputMode is set to HYPERLINK. Valid values:
	//
	// - CURRENT_PAGE: opens in the current page.
	//
	// - NEW_PAGE: opens in a new page.
	//
	// example:
	//
	// NEW_PAGE
	LinkTarget *string `json:"LinkTarget,omitempty" xml:"LinkTarget,omitempty"`
	// The maximum length. This parameter is valid only when ValueType is set to STRING.
	//
	// example:
	//
	// 1000
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// Indicates whether the property is required.
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// The type of the referenced system property. This parameter has a value only when EnumSourceType is set to SYSTEM_REFERENCE.
	//
	// example:
	//
	// USER
	SystemReferenceType *string `json:"SystemReferenceType,omitempty" xml:"SystemReferenceType,omitempty"`
	// The data type of the property value.
	//
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	// The location where the property is visible. Valid values:
	//
	// - ASSET: asset catalog.
	//
	// - DEVELOPMENT: development.
	VisibleIn []*string `json:"VisibleIn,omitempty" xml:"VisibleIn,omitempty" type:"Repeated"`
}

func (s GetAssetTypeAttributeCodesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAssetTypeAttributeCodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetAttributeCode() *string {
	return s.AttributeCode
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetAttributeName() *string {
	return s.AttributeName
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetAttributeSource() *string {
	return s.AttributeSource
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetAttributeType() *string {
	return s.AttributeType
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetEditableIn() []*string {
	return s.EditableIn
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetEnumSourceType() *string {
	return s.EnumSourceType
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetEnumValues() []*GetAssetTypeAttributeCodesResponseBodyDataEnumValues {
	return s.EnumValues
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetInputMode() *string {
	return s.InputMode
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetLinkTarget() *string {
	return s.LinkTarget
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetRequired() *bool {
	return s.Required
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetSystemReferenceType() *string {
	return s.SystemReferenceType
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetValueType() *string {
	return s.ValueType
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) GetVisibleIn() []*string {
	return s.VisibleIn
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetAttributeCode(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.AttributeCode = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetAttributeName(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.AttributeName = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetAttributeSource(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.AttributeSource = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetAttributeType(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.AttributeType = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetDescription(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetEditableIn(v []*string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.EditableIn = v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetEnumSourceType(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.EnumSourceType = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetEnumValues(v []*GetAssetTypeAttributeCodesResponseBodyDataEnumValues) *GetAssetTypeAttributeCodesResponseBodyData {
	s.EnumValues = v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetInputMode(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.InputMode = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetLinkTarget(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.LinkTarget = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetMaxLength(v int32) *GetAssetTypeAttributeCodesResponseBodyData {
	s.MaxLength = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetRequired(v bool) *GetAssetTypeAttributeCodesResponseBodyData {
	s.Required = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetSystemReferenceType(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.SystemReferenceType = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetValueType(v string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.ValueType = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) SetVisibleIn(v []*string) *GetAssetTypeAttributeCodesResponseBodyData {
	s.VisibleIn = v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyData) Validate() error {
	if s.EnumValues != nil {
		for _, item := range s.EnumValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAssetTypeAttributeCodesResponseBodyDataEnumValues struct {
	// The display name of the option.
	//
	// example:
	//
	// Core
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The option value.
	//
	// example:
	//
	// L1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAssetTypeAttributeCodesResponseBodyDataEnumValues) String() string {
	return dara.Prettify(s)
}

func (s GetAssetTypeAttributeCodesResponseBodyDataEnumValues) GoString() string {
	return s.String()
}

func (s *GetAssetTypeAttributeCodesResponseBodyDataEnumValues) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAssetTypeAttributeCodesResponseBodyDataEnumValues) GetValue() *string {
	return s.Value
}

func (s *GetAssetTypeAttributeCodesResponseBodyDataEnumValues) SetDisplayName(v string) *GetAssetTypeAttributeCodesResponseBodyDataEnumValues {
	s.DisplayName = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyDataEnumValues) SetValue(v string) *GetAssetTypeAttributeCodesResponseBodyDataEnumValues {
	s.Value = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponseBodyDataEnumValues) Validate() error {
	return dara.Validate(s)
}
