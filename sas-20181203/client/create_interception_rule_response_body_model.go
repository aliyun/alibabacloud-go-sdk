// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterceptionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInterceptionRuleResponseBody
	GetRequestId() *string
}

type CreateInterceptionRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A01810A0-xxx5E2676
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInterceptionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInterceptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInterceptionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInterceptionRuleResponseBody) SetRequestId(v string) *CreateInterceptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInterceptionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
