// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppImageSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateEdgeContainerAppImageSecretRequest
	GetAppId() *string
	SetPassword(v string) *CreateEdgeContainerAppImageSecretRequest
	GetPassword() *string
	SetRegistry(v string) *CreateEdgeContainerAppImageSecretRequest
	GetRegistry() *string
	SetUsername(v string) *CreateEdgeContainerAppImageSecretRequest
	GetUsername() *string
}

type CreateEdgeContainerAppImageSecretRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cr-cn-shanghai.edas.aliyuncs.com
	Registry *string `json:"Registry,omitempty" xml:"Registry,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateEdgeContainerAppImageSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppImageSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppImageSecretRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateEdgeContainerAppImageSecretRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateEdgeContainerAppImageSecretRequest) GetRegistry() *string {
	return s.Registry
}

func (s *CreateEdgeContainerAppImageSecretRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateEdgeContainerAppImageSecretRequest) SetAppId(v string) *CreateEdgeContainerAppImageSecretRequest {
	s.AppId = &v
	return s
}

func (s *CreateEdgeContainerAppImageSecretRequest) SetPassword(v string) *CreateEdgeContainerAppImageSecretRequest {
	s.Password = &v
	return s
}

func (s *CreateEdgeContainerAppImageSecretRequest) SetRegistry(v string) *CreateEdgeContainerAppImageSecretRequest {
	s.Registry = &v
	return s
}

func (s *CreateEdgeContainerAppImageSecretRequest) SetUsername(v string) *CreateEdgeContainerAppImageSecretRequest {
	s.Username = &v
	return s
}

func (s *CreateEdgeContainerAppImageSecretRequest) Validate() error {
	return dara.Validate(s)
}
