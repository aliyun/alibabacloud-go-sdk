// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLabelTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceId(v string) *UpdateLabelTableRequest
	GetDatasourceId() *string
	SetFields(v []*UpdateLabelTableRequestFields) *UpdateLabelTableRequest
	GetFields() []*UpdateLabelTableRequestFields
	SetName(v string) *UpdateLabelTableRequest
	GetName() *string
}

type UpdateLabelTableRequest struct {
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// This parameter is required.
	Fields []*UpdateLabelTableRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// rec_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateLabelTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLabelTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateLabelTableRequest) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *UpdateLabelTableRequest) GetFields() []*UpdateLabelTableRequestFields {
	return s.Fields
}

func (s *UpdateLabelTableRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLabelTableRequest) SetDatasourceId(v string) *UpdateLabelTableRequest {
	s.DatasourceId = &v
	return s
}

func (s *UpdateLabelTableRequest) SetFields(v []*UpdateLabelTableRequestFields) *UpdateLabelTableRequest {
	s.Fields = v
	return s
}

func (s *UpdateLabelTableRequest) SetName(v string) *UpdateLabelTableRequest {
	s.Name = &v
	return s
}

func (s *UpdateLabelTableRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateLabelTableRequestFields struct {
	// This parameter is required.
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// lat
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DOUBLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateLabelTableRequestFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateLabelTableRequestFields) GoString() string {
	return s.String()
}

func (s *UpdateLabelTableRequestFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *UpdateLabelTableRequestFields) GetName() *string {
	return s.Name
}

func (s *UpdateLabelTableRequestFields) GetType() *string {
	return s.Type
}

func (s *UpdateLabelTableRequestFields) SetAttributes(v []*string) *UpdateLabelTableRequestFields {
	s.Attributes = v
	return s
}

func (s *UpdateLabelTableRequestFields) SetName(v string) *UpdateLabelTableRequestFields {
	s.Name = &v
	return s
}

func (s *UpdateLabelTableRequestFields) SetType(v string) *UpdateLabelTableRequestFields {
	s.Type = &v
	return s
}

func (s *UpdateLabelTableRequestFields) Validate() error {
	return dara.Validate(s)
}
