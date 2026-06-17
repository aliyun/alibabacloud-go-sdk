// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *DescribeAckClusterConnectorRequest
	GetConnectorId() *string
	SetLang(v string) *DescribeAckClusterConnectorRequest
	GetLang() *string
}

type DescribeAckClusterConnectorRequest struct {
	// The ID of the ACK cluster connector. You can obtain the ID by calling the [DescribeAckClusterConnectors](~~DescribeAckClusterConnectors~~) operation to query a list of ACK cluster connectors.
	//
	// - [DescribeAckClusterConnectors](~~DescribeAckClusterConnectors~~): Queries a list of ACK cluster connectors.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The language of the error messages that are returned for the health check status of the ACK cluster connector.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAckClusterConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterConnectorRequest) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DescribeAckClusterConnectorRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAckClusterConnectorRequest) SetConnectorId(v string) *DescribeAckClusterConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *DescribeAckClusterConnectorRequest) SetLang(v string) *DescribeAckClusterConnectorRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAckClusterConnectorRequest) Validate() error {
	return dara.Validate(s)
}
