// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionRuleSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInterceptionRuleSwitchResponseBody
	GetRequestId() *string
}

type ModifyInterceptionRuleSwitchResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DE725A60-95F2-50E8-8F5D-81055215E7DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInterceptionRuleSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionRuleSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionRuleSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInterceptionRuleSwitchResponseBody) SetRequestId(v string) *ModifyInterceptionRuleSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInterceptionRuleSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
