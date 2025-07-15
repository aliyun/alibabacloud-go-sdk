// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePoolResponseBody
	GetRequestId() *string
}

type UpdatePoolResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePoolResponseBody) SetRequestId(v string) *UpdatePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
