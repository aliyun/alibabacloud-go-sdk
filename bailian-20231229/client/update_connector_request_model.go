// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorName(v string) *UpdateConnectorRequest
	GetConnectorName() *string
	SetDescription(v string) *UpdateConnectorRequest
	GetDescription() *string
}

type UpdateConnectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-connector
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// never_delete_aeip_95_us-west-1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectorRequest) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *UpdateConnectorRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConnectorRequest) SetConnectorName(v string) *UpdateConnectorRequest {
	s.ConnectorName = &v
	return s
}

func (s *UpdateConnectorRequest) SetDescription(v string) *UpdateConnectorRequest {
	s.Description = &v
	return s
}

func (s *UpdateConnectorRequest) Validate() error {
	return dara.Validate(s)
}
