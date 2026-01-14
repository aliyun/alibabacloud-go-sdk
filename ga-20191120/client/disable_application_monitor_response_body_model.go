// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableApplicationMonitorResponseBody
	GetRequestId() *string
}

type DisableApplicationMonitorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationMonitorResponseBody) SetRequestId(v string) *DisableApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
