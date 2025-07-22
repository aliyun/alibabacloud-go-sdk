// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoLiveStreamRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAutoLiveStreamRuleResponseBody
	GetRequestId() *string
}

type DisableAutoLiveStreamRuleResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAutoLiveStreamRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAutoLiveStreamRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAutoLiveStreamRuleResponseBody) SetRequestId(v string) *DisableAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAutoLiveStreamRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
