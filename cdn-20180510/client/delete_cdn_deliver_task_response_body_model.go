// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdnDeliverTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCdnDeliverTaskResponseBody
	GetRequestId() *string
}

type DeleteCdnDeliverTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCdnDeliverTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCdnDeliverTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCdnDeliverTaskResponseBody) SetRequestId(v string) *DeleteCdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCdnDeliverTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
