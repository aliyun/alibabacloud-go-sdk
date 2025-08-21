// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAckOperatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallAckOperatorResponseBody
	GetRequestId() *string
	SetResult(v bool) *InstallAckOperatorResponseBody
	GetResult() *bool
}

type InstallAckOperatorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EFA88951-7A6F-4A8E-AB8F-2BB7132BA751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether ES-operator is installed. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InstallAckOperatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAckOperatorResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAckOperatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAckOperatorResponseBody) GetResult() *bool {
	return s.Result
}

func (s *InstallAckOperatorResponseBody) SetRequestId(v string) *InstallAckOperatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAckOperatorResponseBody) SetResult(v bool) *InstallAckOperatorResponseBody {
	s.Result = &v
	return s
}

func (s *InstallAckOperatorResponseBody) Validate() error {
	return dara.Validate(s)
}
