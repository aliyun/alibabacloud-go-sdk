// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayer7RuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLayer7RuleResponseBody
	GetRequestId() *string
}

type CreateLayer7RuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLayer7RuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLayer7RuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLayer7RuleResponseBody) SetRequestId(v string) *CreateLayer7RuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLayer7RuleResponseBody) Validate() error {
	return dara.Validate(s)
}
