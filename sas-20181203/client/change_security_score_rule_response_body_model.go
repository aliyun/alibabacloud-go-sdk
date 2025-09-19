// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSecurityScoreRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeSecurityScoreRuleResponseBody
	GetRequestId() *string
}

type ChangeSecurityScoreRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 23AD0BD2-8771-5647-819E-6BA51E21****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeSecurityScoreRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeSecurityScoreRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeSecurityScoreRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeSecurityScoreRuleResponseBody) SetRequestId(v string) *ChangeSecurityScoreRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeSecurityScoreRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
