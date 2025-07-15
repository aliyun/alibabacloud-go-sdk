// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveEdgeTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveEdgeTransferResponseBody
	GetRequestId() *string
}

type SetLiveEdgeTransferResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveEdgeTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveEdgeTransferResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveEdgeTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveEdgeTransferResponseBody) SetRequestId(v string) *SetLiveEdgeTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveEdgeTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
