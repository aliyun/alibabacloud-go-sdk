// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDynamicUpdateWaterMarkStreamRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DynamicUpdateWaterMarkStreamRuleResponseBody
	GetRequestId() *string
}

type DynamicUpdateWaterMarkStreamRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BE9407FF-F897-4DBD-338D-98A750AD805F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DynamicUpdateWaterMarkStreamRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DynamicUpdateWaterMarkStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DynamicUpdateWaterMarkStreamRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DynamicUpdateWaterMarkStreamRuleResponseBody) SetRequestId(v string) *DynamicUpdateWaterMarkStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
