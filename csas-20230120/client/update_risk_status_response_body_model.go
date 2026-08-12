// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRiskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRiskStatusResponseBody
	GetRequestId() *string
}

type UpdateRiskStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRiskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRiskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRiskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRiskStatusResponseBody) SetRequestId(v string) *UpdateRiskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRiskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
