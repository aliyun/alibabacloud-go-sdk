// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnDeliverTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnDeliverTaskResponseBody
	GetRequestId() *string
}

type DeleteDcdnDeliverTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnDeliverTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDeliverTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnDeliverTaskResponseBody) SetRequestId(v string) *DeleteDcdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnDeliverTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
