// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *UpdateContextStoreRequestConfig) *UpdateContextStoreRequest
	GetConfig() *UpdateContextStoreRequestConfig
	SetContextType(v string) *UpdateContextStoreRequest
	GetContextType() *string
	SetDataset(v *UpdateContextStoreRequestDataset) *UpdateContextStoreRequest
	GetDataset() *UpdateContextStoreRequestDataset
	SetDescription(v string) *UpdateContextStoreRequest
	GetDescription() *string
}

type UpdateContextStoreRequest struct {
	// The configuration information.
	Config *UpdateContextStoreRequestConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The context type.
	//
	// example:
	//
	// memory
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// The dataset information.
	Dataset *UpdateContextStoreRequestDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateContextStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreRequest) GetConfig() *UpdateContextStoreRequestConfig {
	return s.Config
}

func (s *UpdateContextStoreRequest) GetContextType() *string {
	return s.ContextType
}

func (s *UpdateContextStoreRequest) GetDataset() *UpdateContextStoreRequestDataset {
	return s.Dataset
}

func (s *UpdateContextStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateContextStoreRequest) SetConfig(v *UpdateContextStoreRequestConfig) *UpdateContextStoreRequest {
	s.Config = v
	return s
}

func (s *UpdateContextStoreRequest) SetContextType(v string) *UpdateContextStoreRequest {
	s.ContextType = &v
	return s
}

func (s *UpdateContextStoreRequest) SetDataset(v *UpdateContextStoreRequestDataset) *UpdateContextStoreRequest {
	s.Dataset = v
	return s
}

func (s *UpdateContextStoreRequest) SetDescription(v string) *UpdateContextStoreRequest {
	s.Description = &v
	return s
}

func (s *UpdateContextStoreRequest) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextStoreRequestConfig struct {
	// The metadata field.
	MetadataField map[string]*string `json:"metadataField,omitempty" xml:"metadataField,omitempty"`
	// The reference path.
	Source *UpdateContextStoreRequestConfigSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s UpdateContextStoreRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreRequestConfig) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreRequestConfig) GetMetadataField() map[string]*string {
	return s.MetadataField
}

func (s *UpdateContextStoreRequestConfig) GetSource() *UpdateContextStoreRequestConfigSource {
	return s.Source
}

func (s *UpdateContextStoreRequestConfig) SetMetadataField(v map[string]*string) *UpdateContextStoreRequestConfig {
	s.MetadataField = v
	return s
}

func (s *UpdateContextStoreRequestConfig) SetSource(v *UpdateContextStoreRequestConfigSource) *UpdateContextStoreRequestConfig {
	s.Source = v
	return s
}

func (s *UpdateContextStoreRequestConfig) Validate() error {
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextStoreRequestConfigSource struct {
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// sls-test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// sls-test-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1776824891000
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateContextStoreRequestConfigSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreRequestConfigSource) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreRequestConfigSource) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateContextStoreRequestConfigSource) GetProject() *string {
	return s.Project
}

func (s *UpdateContextStoreRequestConfigSource) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateContextStoreRequestConfigSource) SetLogstore(v string) *UpdateContextStoreRequestConfigSource {
	s.Logstore = &v
	return s
}

func (s *UpdateContextStoreRequestConfigSource) SetProject(v string) *UpdateContextStoreRequestConfigSource {
	s.Project = &v
	return s
}

func (s *UpdateContextStoreRequestConfigSource) SetStartTime(v string) *UpdateContextStoreRequestConfigSource {
	s.StartTime = &v
	return s
}

func (s *UpdateContextStoreRequestConfigSource) Validate() error {
	return dara.Validate(s)
}

type UpdateContextStoreRequestDataset struct {
	// The name of the dataset.
	//
	// example:
	//
	// test_dataset
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateContextStoreRequestDataset) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreRequestDataset) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreRequestDataset) GetName() *string {
	return s.Name
}

func (s *UpdateContextStoreRequestDataset) SetName(v string) *UpdateContextStoreRequestDataset {
	s.Name = &v
	return s
}

func (s *UpdateContextStoreRequestDataset) Validate() error {
	return dara.Validate(s)
}
