// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntegrationId(v int64) *DeleteIntegrationsRequest
	GetIntegrationId() *int64
}

type DeleteIntegrationsRequest struct {
	// The ID of the alert integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	IntegrationId *int64 `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
}

func (s DeleteIntegrationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationsRequest) GetIntegrationId() *int64 {
	return s.IntegrationId
}

func (s *DeleteIntegrationsRequest) SetIntegrationId(v int64) *DeleteIntegrationsRequest {
	s.IntegrationId = &v
	return s
}

func (s *DeleteIntegrationsRequest) Validate() error {
	return dara.Validate(s)
}
