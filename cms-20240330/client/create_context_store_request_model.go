// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *CreateContextStoreRequestConfig) *CreateContextStoreRequest
	GetConfig() *CreateContextStoreRequestConfig
	SetContextStoreName(v string) *CreateContextStoreRequest
	GetContextStoreName() *string
	SetContextType(v string) *CreateContextStoreRequest
	GetContextType() *string
	SetDataset(v *CreateContextStoreRequestDataset) *CreateContextStoreRequest
	GetDataset() *CreateContextStoreRequestDataset
	SetDescription(v string) *CreateContextStoreRequest
	GetDescription() *string
}

type CreateContextStoreRequest struct {
	// The configuration.
	Config *CreateContextStoreRequestConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The name of the context store.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// The type of the context store.
	//
	// This parameter is required.
	//
	// example:
	//
	// memory
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// The properties of the dataset.
	Dataset *CreateContextStoreRequestDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The description of the context store.
	//
	// example:
	//
	// desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateContextStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateContextStoreRequest) GetConfig() *CreateContextStoreRequestConfig {
	return s.Config
}

func (s *CreateContextStoreRequest) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *CreateContextStoreRequest) GetContextType() *string {
	return s.ContextType
}

func (s *CreateContextStoreRequest) GetDataset() *CreateContextStoreRequestDataset {
	return s.Dataset
}

func (s *CreateContextStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateContextStoreRequest) SetConfig(v *CreateContextStoreRequestConfig) *CreateContextStoreRequest {
	s.Config = v
	return s
}

func (s *CreateContextStoreRequest) SetContextStoreName(v string) *CreateContextStoreRequest {
	s.ContextStoreName = &v
	return s
}

func (s *CreateContextStoreRequest) SetContextType(v string) *CreateContextStoreRequest {
	s.ContextType = &v
	return s
}

func (s *CreateContextStoreRequest) SetDataset(v *CreateContextStoreRequestDataset) *CreateContextStoreRequest {
	s.Dataset = v
	return s
}

func (s *CreateContextStoreRequest) SetDescription(v string) *CreateContextStoreRequest {
	s.Description = &v
	return s
}

func (s *CreateContextStoreRequest) Validate() error {
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

type CreateContextStoreRequestConfig struct {
	// The metadata fields.
	MetadataField map[string]*string `json:"metadataField,omitempty" xml:"metadataField,omitempty"`
	// The configuration source.
	Source *CreateContextStoreRequestConfigSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s CreateContextStoreRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreRequestConfig) GoString() string {
	return s.String()
}

func (s *CreateContextStoreRequestConfig) GetMetadataField() map[string]*string {
	return s.MetadataField
}

func (s *CreateContextStoreRequestConfig) GetSource() *CreateContextStoreRequestConfigSource {
	return s.Source
}

func (s *CreateContextStoreRequestConfig) SetMetadataField(v map[string]*string) *CreateContextStoreRequestConfig {
	s.MetadataField = v
	return s
}

func (s *CreateContextStoreRequestConfig) SetSource(v *CreateContextStoreRequestConfigSource) *CreateContextStoreRequestConfig {
	s.Source = v
	return s
}

func (s *CreateContextStoreRequestConfig) Validate() error {
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateContextStoreRequestConfigSource struct {
	// The name of the Log Service Logstore.
	//
	// example:
	//
	// sls-test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Log Service project.
	//
	// example:
	//
	// sls-test-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The effective start time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1751508233000
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CreateContextStoreRequestConfigSource) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreRequestConfigSource) GoString() string {
	return s.String()
}

func (s *CreateContextStoreRequestConfigSource) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateContextStoreRequestConfigSource) GetProject() *string {
	return s.Project
}

func (s *CreateContextStoreRequestConfigSource) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateContextStoreRequestConfigSource) SetLogstore(v string) *CreateContextStoreRequestConfigSource {
	s.Logstore = &v
	return s
}

func (s *CreateContextStoreRequestConfigSource) SetProject(v string) *CreateContextStoreRequestConfigSource {
	s.Project = &v
	return s
}

func (s *CreateContextStoreRequestConfigSource) SetStartTime(v string) *CreateContextStoreRequestConfigSource {
	s.StartTime = &v
	return s
}

func (s *CreateContextStoreRequestConfigSource) Validate() error {
	return dara.Validate(s)
}

type CreateContextStoreRequestDataset struct {
	// The name of the dataset.
	//
	// example:
	//
	// test_dataset
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateContextStoreRequestDataset) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreRequestDataset) GoString() string {
	return s.String()
}

func (s *CreateContextStoreRequestDataset) GetName() *string {
	return s.Name
}

func (s *CreateContextStoreRequestDataset) SetName(v string) *CreateContextStoreRequestDataset {
	s.Name = &v
	return s
}

func (s *CreateContextStoreRequestDataset) Validate() error {
	return dara.Validate(s)
}
