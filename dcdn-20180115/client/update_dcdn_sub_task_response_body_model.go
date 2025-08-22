// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnSubTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDcdnSubTaskResponseBody
	GetRequestId() *string
}

type UpdateDcdnSubTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnSubTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSubTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDcdnSubTaskResponseBody) SetRequestId(v string) *UpdateDcdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDcdnSubTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
