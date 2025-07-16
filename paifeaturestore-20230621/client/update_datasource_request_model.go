// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateDatasourceRequest
	GetConfig() *string
	SetName(v string) *UpdateDatasourceRequest
	GetName() *string
	SetUri(v string) *UpdateDatasourceRequest
	GetUri() *string
}

type UpdateDatasourceRequest struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// igraph_instance1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasourceRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateDatasourceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDatasourceRequest) GetUri() *string {
	return s.Uri
}

func (s *UpdateDatasourceRequest) SetConfig(v string) *UpdateDatasourceRequest {
	s.Config = &v
	return s
}

func (s *UpdateDatasourceRequest) SetName(v string) *UpdateDatasourceRequest {
	s.Name = &v
	return s
}

func (s *UpdateDatasourceRequest) SetUri(v string) *UpdateDatasourceRequest {
	s.Uri = &v
	return s
}

func (s *UpdateDatasourceRequest) Validate() error {
	return dara.Validate(s)
}
