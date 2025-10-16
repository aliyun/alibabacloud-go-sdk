// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAckClusterConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *DeleteAckClusterConnectorRequest
	GetConnectorId() *string
}

type DeleteAckClusterConnectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s DeleteAckClusterConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAckClusterConnectorRequest) GoString() string {
	return s.String()
}

func (s *DeleteAckClusterConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DeleteAckClusterConnectorRequest) SetConnectorId(v string) *DeleteAckClusterConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *DeleteAckClusterConnectorRequest) Validate() error {
	return dara.Validate(s)
}
