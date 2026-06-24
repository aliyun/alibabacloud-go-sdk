// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLogstashChargeTypeResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateLogstashChargeTypeResponseBody
	GetResult() *bool
}

type UpdateLogstashChargeTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// - true: The conversion succeeded.
	//
	// - false: The conversion failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateLogstashChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogstashChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLogstashChargeTypeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateLogstashChargeTypeResponseBody) SetRequestId(v string) *UpdateLogstashChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLogstashChargeTypeResponseBody) SetResult(v bool) *UpdateLogstashChargeTypeResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateLogstashChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
