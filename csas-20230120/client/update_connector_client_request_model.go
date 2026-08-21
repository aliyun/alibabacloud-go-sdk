// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *UpdateConnectorClientRequest
	GetConnectorId() *string
	SetDevTag(v string) *UpdateConnectorClientRequest
	GetDevTag() *string
	SetStatus(v string) *UpdateConnectorClientRequest
	GetStatus() *string
}

type UpdateConnectorClientRequest struct {
	// The connector ID. You can call [ListConnectors](~~ListConnectors~~) to query connector IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The unique identifier of the ConnectorClient device. You can call [ListConnectors](~~ListConnectors~~) to query connector information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 672ECBEE-727B-5F43-8D22-90F2BD9E38A7
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The connection status of the ConnectorClient. Valid values:
	//
	// - **Enabled**: connected.
	//
	// - **Disabled**: disconnected.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateConnectorClientRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorClientRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectorClientRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *UpdateConnectorClientRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *UpdateConnectorClientRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateConnectorClientRequest) SetConnectorId(v string) *UpdateConnectorClientRequest {
	s.ConnectorId = &v
	return s
}

func (s *UpdateConnectorClientRequest) SetDevTag(v string) *UpdateConnectorClientRequest {
	s.DevTag = &v
	return s
}

func (s *UpdateConnectorClientRequest) SetStatus(v string) *UpdateConnectorClientRequest {
	s.Status = &v
	return s
}

func (s *UpdateConnectorClientRequest) Validate() error {
	return dara.Validate(s)
}
