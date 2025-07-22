// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoLiveStreamRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAutoLiveStreamRuleResponseBody
	GetRequestId() *string
}

type UpdateAutoLiveStreamRuleResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAutoLiveStreamRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoLiveStreamRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoLiveStreamRuleResponseBody) SetRequestId(v string) *UpdateAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
