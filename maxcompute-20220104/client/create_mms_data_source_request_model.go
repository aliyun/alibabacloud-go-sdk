// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]interface{}) *CreateMmsDataSourceRequest
	GetConfig() map[string]interface{}
	SetName(v string) *CreateMmsDataSourceRequest
	GetName() *string
	SetNetworklink(v string) *CreateMmsDataSourceRequest
	GetNetworklink() *string
	SetType(v string) *CreateMmsDataSourceRequest
	GetType() *string
}

type CreateMmsDataSourceRequest struct {
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc-uf6pc2vordian33gobzfr:cn-shanghai
	Networklink *string `json:"networklink,omitempty" xml:"networklink,omitempty"`
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMmsDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateMmsDataSourceRequest) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *CreateMmsDataSourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateMmsDataSourceRequest) GetNetworklink() *string {
	return s.Networklink
}

func (s *CreateMmsDataSourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateMmsDataSourceRequest) SetConfig(v map[string]interface{}) *CreateMmsDataSourceRequest {
	s.Config = v
	return s
}

func (s *CreateMmsDataSourceRequest) SetName(v string) *CreateMmsDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateMmsDataSourceRequest) SetNetworklink(v string) *CreateMmsDataSourceRequest {
	s.Networklink = &v
	return s
}

func (s *CreateMmsDataSourceRequest) SetType(v string) *CreateMmsDataSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateMmsDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
