// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExpressConnectTrafficQosQueueResponseBody
	GetRequestId() *string
}

type DeleteExpressConnectTrafficQosQueueResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B9300FE-11E2-4E3B-949C-BED3B44DD26D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExpressConnectTrafficQosQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExpressConnectTrafficQosQueueResponseBody) SetRequestId(v string) *DeleteExpressConnectTrafficQosQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
