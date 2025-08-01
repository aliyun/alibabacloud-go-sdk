// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteMigrationTaskRequest
	GetAcceptLanguage() *string
	SetId(v string) *DeleteMigrationTaskRequest
	GetId() *string
	SetRequestPars(v string) *DeleteMigrationTaskRequest
	GetRequestPars() *string
}

type DeleteMigrationTaskRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s DeleteMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteMigrationTaskRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteMigrationTaskRequest) GetId() *string {
	return s.Id
}

func (s *DeleteMigrationTaskRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *DeleteMigrationTaskRequest) SetAcceptLanguage(v string) *DeleteMigrationTaskRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteMigrationTaskRequest) SetId(v string) *DeleteMigrationTaskRequest {
	s.Id = &v
	return s
}

func (s *DeleteMigrationTaskRequest) SetRequestPars(v string) *DeleteMigrationTaskRequest {
	s.RequestPars = &v
	return s
}

func (s *DeleteMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
