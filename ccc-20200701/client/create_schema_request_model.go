// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSchemaRequest
	GetDescription() *string
	SetId(v string) *CreateSchemaRequest
	GetId() *string
	SetInstanceId(v string) *CreateSchemaRequest
	GetInstanceId() *string
	SetProperties(v []*CreateSchemaRequestProperties) *CreateSchemaRequest
	GetProperties() []*CreateSchemaRequestProperties
	SetRequestId(v string) *CreateSchemaRequest
	GetRequestId() *string
}

type CreateSchemaRequest struct {
	// example:
	//
	// -
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// schema id
	//
	// example:
	//
	// profile
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b0eb2742-f37e-4c67-82d4-25c651c1c450
	InstanceId *string                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Properties []*CreateSchemaRequestProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemaRequest) GoString() string {
	return s.String()
}

func (s *CreateSchemaRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSchemaRequest) GetId() *string {
	return s.Id
}

func (s *CreateSchemaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSchemaRequest) GetProperties() []*CreateSchemaRequestProperties {
	return s.Properties
}

func (s *CreateSchemaRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSchemaRequest) SetDescription(v string) *CreateSchemaRequest {
	s.Description = &v
	return s
}

func (s *CreateSchemaRequest) SetId(v string) *CreateSchemaRequest {
	s.Id = &v
	return s
}

func (s *CreateSchemaRequest) SetInstanceId(v string) *CreateSchemaRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSchemaRequest) SetProperties(v []*CreateSchemaRequestProperties) *CreateSchemaRequest {
	s.Properties = v
	return s
}

func (s *CreateSchemaRequest) SetRequestId(v string) *CreateSchemaRequest {
	s.RequestId = &v
	return s
}

func (s *CreateSchemaRequest) Validate() error {
	return dara.Validate(s)
}

type CreateSchemaRequestProperties struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// {\\"Clusters\\": {\\"Description\\": \\"The list of clusters.\\"}, \\"ClusterIds\\": {\\"Description\\": \\"The list of cluster IDs.\\"}}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// -
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	Disabled    *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2
	DisplayOrder *int32 `json:"DisplayOrder,omitempty" xml:"DisplayOrder,omitempty"`
	// example:
	//
	// textbox
	EditorType *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	// example:
	//
	// 100
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// example:
	//
	// 1
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// example:
	//
	// 1
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// example:
	//
	// 1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// *
	Pattern             *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	PatternErrorMessage *string `json:"PatternErrorMessage,omitempty" xml:"PatternErrorMessage,omitempty"`
	// example:
	//
	// true
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s CreateSchemaRequestProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemaRequestProperties) GoString() string {
	return s.String()
}

func (s *CreateSchemaRequestProperties) GetArray() *bool {
	return s.Array
}

func (s *CreateSchemaRequestProperties) GetAttributes() *string {
	return s.Attributes
}

func (s *CreateSchemaRequestProperties) GetDataType() *string {
	return s.DataType
}

func (s *CreateSchemaRequestProperties) GetDescription() *string {
	return s.Description
}

func (s *CreateSchemaRequestProperties) GetDisabled() *bool {
	return s.Disabled
}

func (s *CreateSchemaRequestProperties) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateSchemaRequestProperties) GetDisplayOrder() *int32 {
	return s.DisplayOrder
}

func (s *CreateSchemaRequestProperties) GetEditorType() *string {
	return s.EditorType
}

func (s *CreateSchemaRequestProperties) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *CreateSchemaRequestProperties) GetMaximum() *float64 {
	return s.Maximum
}

func (s *CreateSchemaRequestProperties) GetMinLength() *int32 {
	return s.MinLength
}

func (s *CreateSchemaRequestProperties) GetMinimum() *float64 {
	return s.Minimum
}

func (s *CreateSchemaRequestProperties) GetName() *string {
	return s.Name
}

func (s *CreateSchemaRequestProperties) GetPattern() *string {
	return s.Pattern
}

func (s *CreateSchemaRequestProperties) GetPatternErrorMessage() *string {
	return s.PatternErrorMessage
}

func (s *CreateSchemaRequestProperties) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateSchemaRequestProperties) GetRequired() *bool {
	return s.Required
}

func (s *CreateSchemaRequestProperties) SetArray(v bool) *CreateSchemaRequestProperties {
	s.Array = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetAttributes(v string) *CreateSchemaRequestProperties {
	s.Attributes = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetDataType(v string) *CreateSchemaRequestProperties {
	s.DataType = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetDescription(v string) *CreateSchemaRequestProperties {
	s.Description = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetDisabled(v bool) *CreateSchemaRequestProperties {
	s.Disabled = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetDisplayName(v string) *CreateSchemaRequestProperties {
	s.DisplayName = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetDisplayOrder(v int32) *CreateSchemaRequestProperties {
	s.DisplayOrder = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetEditorType(v string) *CreateSchemaRequestProperties {
	s.EditorType = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetMaxLength(v int32) *CreateSchemaRequestProperties {
	s.MaxLength = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetMaximum(v float64) *CreateSchemaRequestProperties {
	s.Maximum = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetMinLength(v int32) *CreateSchemaRequestProperties {
	s.MinLength = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetMinimum(v float64) *CreateSchemaRequestProperties {
	s.Minimum = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetName(v string) *CreateSchemaRequestProperties {
	s.Name = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetPattern(v string) *CreateSchemaRequestProperties {
	s.Pattern = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetPatternErrorMessage(v string) *CreateSchemaRequestProperties {
	s.PatternErrorMessage = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetReadOnly(v bool) *CreateSchemaRequestProperties {
	s.ReadOnly = &v
	return s
}

func (s *CreateSchemaRequestProperties) SetRequired(v bool) *CreateSchemaRequestProperties {
	s.Required = &v
	return s
}

func (s *CreateSchemaRequestProperties) Validate() error {
	return dara.Validate(s)
}
