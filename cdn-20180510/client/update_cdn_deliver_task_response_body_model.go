// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCdnDeliverTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCdnDeliverTaskResponseBody
	GetRequestId() *string
}

type UpdateCdnDeliverTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCdnDeliverTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCdnDeliverTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCdnDeliverTaskResponseBody) SetRequestId(v string) *UpdateCdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCdnDeliverTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
