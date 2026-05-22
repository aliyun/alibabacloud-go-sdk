// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppImageSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteEdgeContainerAppImageSecretRequest
	GetAppId() *string
	SetName(v string) *DeleteEdgeContainerAppImageSecretRequest
	GetName() *string
}

type DeleteEdgeContainerAppImageSecretRequest struct {
	// Application ID, which can be obtained using the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Name of the image secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// reg-123*****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteEdgeContainerAppImageSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppImageSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppImageSecretRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteEdgeContainerAppImageSecretRequest) GetName() *string {
	return s.Name
}

func (s *DeleteEdgeContainerAppImageSecretRequest) SetAppId(v string) *DeleteEdgeContainerAppImageSecretRequest {
	s.AppId = &v
	return s
}

func (s *DeleteEdgeContainerAppImageSecretRequest) SetName(v string) *DeleteEdgeContainerAppImageSecretRequest {
	s.Name = &v
	return s
}

func (s *DeleteEdgeContainerAppImageSecretRequest) Validate() error {
	return dara.Validate(s)
}
