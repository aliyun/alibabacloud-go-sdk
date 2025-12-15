// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSoftwaresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallSoftwaresResponseBody
	GetRequestId() *string
}

type InstallSoftwaresResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallSoftwaresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallSoftwaresResponseBody) SetRequestId(v string) *InstallSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallSoftwaresResponseBody) Validate() error {
	return dara.Validate(s)
}
