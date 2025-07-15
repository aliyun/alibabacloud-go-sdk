// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveEdgeTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveEdgeTransferResponseBody
	GetRequestId() *string
}

type DeleteLiveEdgeTransferResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveEdgeTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveEdgeTransferResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveEdgeTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveEdgeTransferResponseBody) SetRequestId(v string) *DeleteLiveEdgeTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveEdgeTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
