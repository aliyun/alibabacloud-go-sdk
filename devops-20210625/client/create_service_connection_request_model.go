// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *CreateServiceConnectionRequest
	GetAuthType() *string
	SetConnectionName(v string) *CreateServiceConnectionRequest
	GetConnectionName() *string
	SetConnectionType(v string) *CreateServiceConnectionRequest
	GetConnectionType() *string
	SetScope(v string) *CreateServiceConnectionRequest
	GetScope() *string
	SetServiceAuthId(v int64) *CreateServiceConnectionRequest
	GetServiceAuthId() *int64
}

type CreateServiceConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// CREDENTIAL
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// This parameter is required.
	ConnectionName *string `json:"connectionName,omitempty" xml:"connectionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ConnectionType *string `json:"connectionType,omitempty" xml:"connectionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PERSON
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	ServiceAuthId *int64 `json:"serviceAuthId,omitempty" xml:"serviceAuthId,omitempty"`
}

func (s CreateServiceConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceConnectionRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateServiceConnectionRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateServiceConnectionRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *CreateServiceConnectionRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateServiceConnectionRequest) GetServiceAuthId() *int64 {
	return s.ServiceAuthId
}

func (s *CreateServiceConnectionRequest) SetAuthType(v string) *CreateServiceConnectionRequest {
	s.AuthType = &v
	return s
}

func (s *CreateServiceConnectionRequest) SetConnectionName(v string) *CreateServiceConnectionRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateServiceConnectionRequest) SetConnectionType(v string) *CreateServiceConnectionRequest {
	s.ConnectionType = &v
	return s
}

func (s *CreateServiceConnectionRequest) SetScope(v string) *CreateServiceConnectionRequest {
	s.Scope = &v
	return s
}

func (s *CreateServiceConnectionRequest) SetServiceAuthId(v int64) *CreateServiceConnectionRequest {
	s.ServiceAuthId = &v
	return s
}

func (s *CreateServiceConnectionRequest) Validate() error {
	return dara.Validate(s)
}
