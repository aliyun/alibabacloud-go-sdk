// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateExternalDatabaseRequest
	GetComment() *string
	SetDatabaseName(v string) *CreateExternalDatabaseRequest
	GetDatabaseName() *string
	SetDefaultUserMapping(v string) *CreateExternalDatabaseRequest
	GetDefaultUserMapping() *string
	SetExternalConfig(v []*CreateExternalDatabaseRequestExternalConfig) *CreateExternalDatabaseRequest
	GetExternalConfig() []*CreateExternalDatabaseRequestExternalConfig
	SetMetastoreType(v string) *CreateExternalDatabaseRequest
	GetMetastoreType() *string
}

type CreateExternalDatabaseRequest struct {
	// example:
	//
	// for log stat
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// my_db
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// example:
	//
	// true
	DefaultUserMapping *string                                        `json:"defaultUserMapping,omitempty" xml:"defaultUserMapping,omitempty"`
	ExternalConfig     []*CreateExternalDatabaseRequestExternalConfig `json:"externalConfig,omitempty" xml:"externalConfig,omitempty" type:"Repeated"`
	// example:
	//
	// maxcompute
	MetastoreType *string `json:"metastoreType,omitempty" xml:"metastoreType,omitempty"`
}

func (s CreateExternalDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateExternalDatabaseRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateExternalDatabaseRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CreateExternalDatabaseRequest) GetDefaultUserMapping() *string {
	return s.DefaultUserMapping
}

func (s *CreateExternalDatabaseRequest) GetExternalConfig() []*CreateExternalDatabaseRequestExternalConfig {
	return s.ExternalConfig
}

func (s *CreateExternalDatabaseRequest) GetMetastoreType() *string {
	return s.MetastoreType
}

func (s *CreateExternalDatabaseRequest) SetComment(v string) *CreateExternalDatabaseRequest {
	s.Comment = &v
	return s
}

func (s *CreateExternalDatabaseRequest) SetDatabaseName(v string) *CreateExternalDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateExternalDatabaseRequest) SetDefaultUserMapping(v string) *CreateExternalDatabaseRequest {
	s.DefaultUserMapping = &v
	return s
}

func (s *CreateExternalDatabaseRequest) SetExternalConfig(v []*CreateExternalDatabaseRequestExternalConfig) *CreateExternalDatabaseRequest {
	s.ExternalConfig = v
	return s
}

func (s *CreateExternalDatabaseRequest) SetMetastoreType(v string) *CreateExternalDatabaseRequest {
	s.MetastoreType = &v
	return s
}

func (s *CreateExternalDatabaseRequest) Validate() error {
	if s.ExternalConfig != nil {
		for _, item := range s.ExternalConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateExternalDatabaseRequestExternalConfig struct {
	// example:
	//
	// mc_project
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// log_sum
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateExternalDatabaseRequestExternalConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalDatabaseRequestExternalConfig) GoString() string {
	return s.String()
}

func (s *CreateExternalDatabaseRequestExternalConfig) GetKey() *string {
	return s.Key
}

func (s *CreateExternalDatabaseRequestExternalConfig) GetValue() *string {
	return s.Value
}

func (s *CreateExternalDatabaseRequestExternalConfig) SetKey(v string) *CreateExternalDatabaseRequestExternalConfig {
	s.Key = &v
	return s
}

func (s *CreateExternalDatabaseRequestExternalConfig) SetValue(v string) *CreateExternalDatabaseRequestExternalConfig {
	s.Value = &v
	return s
}

func (s *CreateExternalDatabaseRequestExternalConfig) Validate() error {
	return dara.Validate(s)
}
