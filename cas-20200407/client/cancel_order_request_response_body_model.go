// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelOrderRequestResponseBody
	GetRequestId() *string
}

type CancelOrderRequestResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOrderRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelOrderRequestResponseBody) SetRequestId(v string) *CancelOrderRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOrderRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
