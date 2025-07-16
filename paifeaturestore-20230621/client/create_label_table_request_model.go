// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLabelTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceId(v string) *CreateLabelTableRequest
	GetDatasourceId() *string
	SetFields(v []*CreateLabelTableRequestFields) *CreateLabelTableRequest
	GetFields() []*CreateLabelTableRequestFields
	SetName(v string) *CreateLabelTableRequest
	GetName() *string
	SetProjectId(v string) *CreateLabelTableRequest
	GetProjectId() *string
}

type CreateLabelTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// This parameter is required.
	Fields []*CreateLabelTableRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rec_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateLabelTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLabelTableRequest) GoString() string {
	return s.String()
}

func (s *CreateLabelTableRequest) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *CreateLabelTableRequest) GetFields() []*CreateLabelTableRequestFields {
	return s.Fields
}

func (s *CreateLabelTableRequest) GetName() *string {
	return s.Name
}

func (s *CreateLabelTableRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateLabelTableRequest) SetDatasourceId(v string) *CreateLabelTableRequest {
	s.DatasourceId = &v
	return s
}

func (s *CreateLabelTableRequest) SetFields(v []*CreateLabelTableRequestFields) *CreateLabelTableRequest {
	s.Fields = v
	return s
}

func (s *CreateLabelTableRequest) SetName(v string) *CreateLabelTableRequest {
	s.Name = &v
	return s
}

func (s *CreateLabelTableRequest) SetProjectId(v string) *CreateLabelTableRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateLabelTableRequest) Validate() error {
	return dara.Validate(s)
}

type CreateLabelTableRequestFields struct {
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
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateLabelTableRequestFields) String() string {
	return dara.Prettify(s)
}

func (s CreateLabelTableRequestFields) GoString() string {
	return s.String()
}

func (s *CreateLabelTableRequestFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *CreateLabelTableRequestFields) GetName() *string {
	return s.Name
}

func (s *CreateLabelTableRequestFields) GetType() *string {
	return s.Type
}

func (s *CreateLabelTableRequestFields) SetAttributes(v []*string) *CreateLabelTableRequestFields {
	s.Attributes = v
	return s
}

func (s *CreateLabelTableRequestFields) SetName(v string) *CreateLabelTableRequestFields {
	s.Name = &v
	return s
}

func (s *CreateLabelTableRequestFields) SetType(v string) *CreateLabelTableRequestFields {
	s.Type = &v
	return s
}

func (s *CreateLabelTableRequestFields) Validate() error {
	return dara.Validate(s)
}
