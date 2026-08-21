// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *GetConnectorRequest
	GetConnectorId() *string
}

type GetConnectorRequest struct {
	// ConnectorID。
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-fcd9c35583087b2f
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s GetConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorRequest) GoString() string {
	return s.String()
}

func (s *GetConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetConnectorRequest) SetConnectorId(v string) *GetConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *GetConnectorRequest) Validate() error {
	return dara.Validate(s)
}
