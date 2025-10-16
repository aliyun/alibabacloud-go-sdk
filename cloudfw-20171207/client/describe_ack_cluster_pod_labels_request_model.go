// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterPodLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *DescribeAckClusterPodLabelsRequest
	GetConnectorId() *string
}

type DescribeAckClusterPodLabelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s DescribeAckClusterPodLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterPodLabelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterPodLabelsRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DescribeAckClusterPodLabelsRequest) SetConnectorId(v string) *DescribeAckClusterPodLabelsRequest {
	s.ConnectorId = &v
	return s
}

func (s *DescribeAckClusterPodLabelsRequest) Validate() error {
	return dara.Validate(s)
}
