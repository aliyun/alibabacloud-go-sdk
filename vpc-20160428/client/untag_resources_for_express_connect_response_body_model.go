// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesForExpressConnectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UntagResourcesForExpressConnectResponseBody
	GetRequestId() *string
}

type UntagResourcesForExpressConnectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DE65F6B7-7566-4802-9007-96F2494AC512
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesForExpressConnectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesForExpressConnectResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesForExpressConnectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesForExpressConnectResponseBody) SetRequestId(v string) *UntagResourcesForExpressConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesForExpressConnectResponseBody) Validate() error {
	return dara.Validate(s)
}
