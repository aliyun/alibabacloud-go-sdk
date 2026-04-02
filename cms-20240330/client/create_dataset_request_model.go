// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateDatasetRequest
	GetDatasetName() *string
	SetDescription(v string) *CreateDatasetRequest
	GetDescription() *string
	SetSchema(v map[string]*IndexKey) *CreateDatasetRequest
	GetSchema() map[string]*IndexKey
}

type CreateDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_dataset
	DatasetName *string `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Schema map[string]*IndexKey `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetRequest) GetSchema() map[string]*IndexKey {
	return s.Schema
}

func (s *CreateDatasetRequest) SetDatasetName(v string) *CreateDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetRequest) SetDescription(v string) *CreateDatasetRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetRequest) SetSchema(v map[string]*IndexKey) *CreateDatasetRequest {
	s.Schema = v
	return s
}

func (s *CreateDatasetRequest) Validate() error {
	return dara.Validate(s)
}
