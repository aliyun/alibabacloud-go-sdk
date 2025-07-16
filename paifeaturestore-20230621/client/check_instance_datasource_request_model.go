// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CheckInstanceDatasourceRequest
	GetConfig() *string
	SetType(v string) *CheckInstanceDatasourceRequest
	GetType() *string
	SetUri(v string) *CheckInstanceDatasourceRequest
	GetUri() *string
}

type CheckInstanceDatasourceRequest struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// igraph1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CheckInstanceDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceDatasourceRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceDatasourceRequest) GetConfig() *string {
	return s.Config
}

func (s *CheckInstanceDatasourceRequest) GetType() *string {
	return s.Type
}

func (s *CheckInstanceDatasourceRequest) GetUri() *string {
	return s.Uri
}

func (s *CheckInstanceDatasourceRequest) SetConfig(v string) *CheckInstanceDatasourceRequest {
	s.Config = &v
	return s
}

func (s *CheckInstanceDatasourceRequest) SetType(v string) *CheckInstanceDatasourceRequest {
	s.Type = &v
	return s
}

func (s *CheckInstanceDatasourceRequest) SetUri(v string) *CheckInstanceDatasourceRequest {
	s.Uri = &v
	return s
}

func (s *CheckInstanceDatasourceRequest) Validate() error {
	return dara.Validate(s)
}
