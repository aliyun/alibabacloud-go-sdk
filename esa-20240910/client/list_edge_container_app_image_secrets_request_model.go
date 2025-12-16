// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppImageSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListEdgeContainerAppImageSecretsRequest
	GetAppId() *string
}

type ListEdgeContainerAppImageSecretsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListEdgeContainerAppImageSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppImageSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppImageSecretsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListEdgeContainerAppImageSecretsRequest) SetAppId(v string) *ListEdgeContainerAppImageSecretsRequest {
	s.AppId = &v
	return s
}

func (s *ListEdgeContainerAppImageSecretsRequest) Validate() error {
	return dara.Validate(s)
}
