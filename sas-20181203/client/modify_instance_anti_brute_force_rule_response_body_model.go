// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAntiBruteForceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceAntiBruteForceRuleResponseBody
	GetRequestId() *string
}

type ModifyInstanceAntiBruteForceRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4BE468CE-47A0-54F0-98A1-E253546E6A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAntiBruteForceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAntiBruteForceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAntiBruteForceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceAntiBruteForceRuleResponseBody) SetRequestId(v string) *ModifyInstanceAntiBruteForceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
