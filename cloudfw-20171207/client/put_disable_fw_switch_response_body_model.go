// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDisableFwSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *PutDisableFwSwitchResponseBody
	GetDryRun() *bool
	SetRequestId(v string) *PutDisableFwSwitchResponseBody
	GetRequestId() *string
}

type PutDisableFwSwitchResponseBody struct {
	// Indicates whether only a dry run was performed. If this parameter is true, the system performed pre-checks such as parameter validity, identity permissions, resource existence, quota limits, and dependency relationships without creating, updating, or deleting actual resources, triggering asynchronous traffic diversion tasks, or generating downstream side effects such as billing, notifications, or callbacks. If the dry run succeeded, DryRun=true is returned in the response, which can be distinguished from an actual call response. If the dry run failed, a machine-readable error code is returned (such as ErrorParamsInvalid for parameter errors, ErrorAuthentication for insufficient permissions, or ErrorInstanceOpenIpNumExceed for insufficient quota). A value of false (default) indicates that the request was initiated and the enable operation was performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDisableFwSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDisableFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchResponseBody) GetDryRun() *bool {
	return s.DryRun
}

func (s *PutDisableFwSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutDisableFwSwitchResponseBody) SetDryRun(v bool) *PutDisableFwSwitchResponseBody {
	s.DryRun = &v
	return s
}

func (s *PutDisableFwSwitchResponseBody) SetRequestId(v string) *PutDisableFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutDisableFwSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
