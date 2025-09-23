// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayer4RuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLayer4RuleResponseBody
	GetRequestId() *string
}

type DeleteLayer4RuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLayer4RuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayer4RuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLayer4RuleResponseBody) SetRequestId(v string) *DeleteLayer4RuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLayer4RuleResponseBody) Validate() error {
	return dara.Validate(s)
}
