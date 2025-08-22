// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnDeliverTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDcdnDeliverTaskResponseBody
	GetRequestId() *string
}

type UpdateDcdnDeliverTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnDeliverTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDeliverTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDcdnDeliverTaskResponseBody) SetRequestId(v string) *UpdateDcdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDcdnDeliverTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
