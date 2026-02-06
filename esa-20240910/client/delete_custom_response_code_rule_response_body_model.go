// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomResponseCodeRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomResponseCodeRuleResponseBody
	GetRequestId() *string
}

type DeleteCustomResponseCodeRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomResponseCodeRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomResponseCodeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomResponseCodeRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomResponseCodeRuleResponseBody) SetRequestId(v string) *DeleteCustomResponseCodeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomResponseCodeRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
