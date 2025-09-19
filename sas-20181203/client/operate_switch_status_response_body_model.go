// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSwitchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateSwitchStatusResponseBody
	GetRequestId() *string
}

type OperateSwitchStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8AE9D3DA-406B-51FA-AA1C-89440C1459BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateSwitchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *OperateSwitchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateSwitchStatusResponseBody) SetRequestId(v string) *OperateSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateSwitchStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
