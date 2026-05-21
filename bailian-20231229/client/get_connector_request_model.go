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
	SetConnectorName(v string) *GetConnectorRequest
	GetConnectorName() *string
}

type GetConnectorRequest struct {
	// example:
	//
	// file_conn_xxxx
	ConnectorId   *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
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

func (s *GetConnectorRequest) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *GetConnectorRequest) SetConnectorId(v string) *GetConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *GetConnectorRequest) SetConnectorName(v string) *GetConnectorRequest {
	s.ConnectorName = &v
	return s
}

func (s *GetConnectorRequest) Validate() error {
	return dara.Validate(s)
}
