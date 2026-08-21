// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *GetConnectorClientRequest
	GetConnectorId() *string
	SetDevTag(v string) *GetConnectorClientRequest
	GetDevTag() *string
}

type GetConnectorClientRequest struct {
	// The connector ID. You can call [ListConnectors](~~ListConnectors~~) to query connectors.
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The unique device identifier of the ConnectorClient. You can call [ListConnectors](~~ListConnectors~~) to query connectors.
	//
	// This parameter is required.
	//
	// example:
	//
	// E4BD65C4-58F6-5127-AD2F-319CF020F549
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
}

func (s GetConnectorClientRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorClientRequest) GoString() string {
	return s.String()
}

func (s *GetConnectorClientRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetConnectorClientRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *GetConnectorClientRequest) SetConnectorId(v string) *GetConnectorClientRequest {
	s.ConnectorId = &v
	return s
}

func (s *GetConnectorClientRequest) SetDevTag(v string) *GetConnectorClientRequest {
	s.DevTag = &v
	return s
}

func (s *GetConnectorClientRequest) Validate() error {
	return dara.Validate(s)
}
