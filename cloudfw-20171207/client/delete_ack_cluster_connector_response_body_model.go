// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAckClusterConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAckClusterConnectorResponseBody
	GetRequestId() *string
}

type DeleteAckClusterConnectorResponseBody struct {
	// example:
	//
	// 133173B9-8010-5DF5-8B93-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAckClusterConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAckClusterConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAckClusterConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAckClusterConnectorResponseBody) SetRequestId(v string) *DeleteAckClusterConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAckClusterConnectorResponseBody) Validate() error {
	return dara.Validate(s)
}
