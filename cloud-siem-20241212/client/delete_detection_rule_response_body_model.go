// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDetectionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDetectionRuleResponseBody
	GetRequestId() *string
}

type DeleteDetectionRuleResponseBody struct {
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDetectionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDetectionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDetectionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDetectionRuleResponseBody) SetRequestId(v string) *DeleteDetectionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDetectionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
