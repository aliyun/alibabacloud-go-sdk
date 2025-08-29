// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTrafficControlTaskExpressionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsValie(v bool) *CheckTrafficControlTaskExpressionResponseBody
	GetIsValie() *bool
	SetReason(v string) *CheckTrafficControlTaskExpressionResponseBody
	GetReason() *string
	SetRequestId(v string) *CheckTrafficControlTaskExpressionResponseBody
	GetRequestId() *string
}

type CheckTrafficControlTaskExpressionResponseBody struct {
	IsValie   *bool   `json:"IsValie,omitempty" xml:"IsValie,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckTrafficControlTaskExpressionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckTrafficControlTaskExpressionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTrafficControlTaskExpressionResponseBody) GetIsValie() *bool {
	return s.IsValie
}

func (s *CheckTrafficControlTaskExpressionResponseBody) GetReason() *string {
	return s.Reason
}

func (s *CheckTrafficControlTaskExpressionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckTrafficControlTaskExpressionResponseBody) SetIsValie(v bool) *CheckTrafficControlTaskExpressionResponseBody {
	s.IsValie = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponseBody) SetReason(v string) *CheckTrafficControlTaskExpressionResponseBody {
	s.Reason = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponseBody) SetRequestId(v string) *CheckTrafficControlTaskExpressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponseBody) Validate() error {
	return dara.Validate(s)
}
