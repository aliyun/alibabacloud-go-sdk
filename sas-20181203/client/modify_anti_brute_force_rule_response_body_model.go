// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAntiBruteForceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAntiBruteForceRuleResponseBody
	GetRequestId() *string
}

type ModifyAntiBruteForceRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D0760E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAntiBruteForceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAntiBruteForceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAntiBruteForceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAntiBruteForceRuleResponseBody) SetRequestId(v string) *ModifyAntiBruteForceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAntiBruteForceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
