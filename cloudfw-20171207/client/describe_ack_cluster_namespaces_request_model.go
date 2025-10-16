// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *DescribeAckClusterNamespacesRequest
	GetConnectorId() *string
}

type DescribeAckClusterNamespacesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s DescribeAckClusterNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterNamespacesRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DescribeAckClusterNamespacesRequest) SetConnectorId(v string) *DescribeAckClusterNamespacesRequest {
	s.ConnectorId = &v
	return s
}

func (s *DescribeAckClusterNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
