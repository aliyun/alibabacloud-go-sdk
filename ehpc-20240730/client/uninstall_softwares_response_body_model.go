// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallSoftwaresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallSoftwaresResponseBody
	GetRequestId() *string
}

type UninstallSoftwaresResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallSoftwaresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallSoftwaresResponseBody) SetRequestId(v string) *UninstallSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallSoftwaresResponseBody) Validate() error {
	return dara.Validate(s)
}
