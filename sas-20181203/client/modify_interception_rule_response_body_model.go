// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInterceptionRuleResponseBody
	GetRequestId() *string
}

type ModifyInterceptionRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6961B151-B43C-533B-8B2E-1D3151D7F5B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInterceptionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInterceptionRuleResponseBody) SetRequestId(v string) *ModifyInterceptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInterceptionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
