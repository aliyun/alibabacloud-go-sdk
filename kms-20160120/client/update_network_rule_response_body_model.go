// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNetworkRuleResponseBody
	GetRequestId() *string
}

type UpdateNetworkRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f34-be0f-cc043fda2d85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetworkRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetworkRuleResponseBody) SetRequestId(v string) *UpdateNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
