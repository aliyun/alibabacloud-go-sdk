// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAckClusterConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *UpdateAckClusterConnectorRequest
	GetConnectorId() *string
	SetConnectorName(v string) *UpdateAckClusterConnectorRequest
	GetConnectorName() *string
	SetTtl(v string) *UpdateAckClusterConnectorRequest
	GetTtl() *string
}

type UpdateAckClusterConnectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// ack-cluster-connector-name
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// example:
	//
	// 30
	Ttl *string `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateAckClusterConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAckClusterConnectorRequest) GoString() string {
	return s.String()
}

func (s *UpdateAckClusterConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *UpdateAckClusterConnectorRequest) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *UpdateAckClusterConnectorRequest) GetTtl() *string {
	return s.Ttl
}

func (s *UpdateAckClusterConnectorRequest) SetConnectorId(v string) *UpdateAckClusterConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *UpdateAckClusterConnectorRequest) SetConnectorName(v string) *UpdateAckClusterConnectorRequest {
	s.ConnectorName = &v
	return s
}

func (s *UpdateAckClusterConnectorRequest) SetTtl(v string) *UpdateAckClusterConnectorRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateAckClusterConnectorRequest) Validate() error {
	return dara.Validate(s)
}
