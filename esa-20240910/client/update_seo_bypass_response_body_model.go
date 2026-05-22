// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSeoBypassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSeoBypassResponseBody
	GetRequestId() *string
}

type UpdateSeoBypassResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSeoBypassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSeoBypassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSeoBypassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSeoBypassResponseBody) SetRequestId(v string) *UpdateSeoBypassResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSeoBypassResponseBody) Validate() error {
	return dara.Validate(s)
}
