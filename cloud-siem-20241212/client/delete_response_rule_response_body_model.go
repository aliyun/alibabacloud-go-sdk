// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResponseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResponseRuleResponseBody
	GetRequestId() *string
}

type DeleteResponseRuleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9BCE0BF1-4F0C-5860-87D2-C391799AF4F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResponseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResponseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResponseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResponseRuleResponseBody) SetRequestId(v string) *DeleteResponseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResponseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
