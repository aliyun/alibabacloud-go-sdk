// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *DeleteConnectorRequest
	GetConnectorId() *string
}

type DeleteConnectorRequest struct {
	// ConnectorID。
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-d02b62911b2fb2d4
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s DeleteConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DeleteConnectorRequest) SetConnectorId(v string) *DeleteConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *DeleteConnectorRequest) Validate() error {
	return dara.Validate(s)
}
