// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectorClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *DeleteConnectorClientRequest
	GetConnectorId() *string
	SetDevTag(v string) *DeleteConnectorClientRequest
	GetDevTag() *string
}

type DeleteConnectorClientRequest struct {
	// The connector ID. You can call [ListConnectors](~~ListConnectors~~) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The unique identifier of the ConnectorClient device. You can call [ListConnectors](~~ListConnectors~~) to obtain the identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// E4BD65C4-58F6-5127-AD2F-319CF020F549
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
}

func (s DeleteConnectorClientRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorClientRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectorClientRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DeleteConnectorClientRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *DeleteConnectorClientRequest) SetConnectorId(v string) *DeleteConnectorClientRequest {
	s.ConnectorId = &v
	return s
}

func (s *DeleteConnectorClientRequest) SetDevTag(v string) *DeleteConnectorClientRequest {
	s.DevTag = &v
	return s
}

func (s *DeleteConnectorClientRequest) Validate() error {
	return dara.Validate(s)
}
