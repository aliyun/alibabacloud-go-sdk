// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoLiveStreamRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAutoLiveStreamRuleResponseBody
	GetRequestId() *string
}

type DeleteAutoLiveStreamRuleResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoLiveStreamRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoLiveStreamRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutoLiveStreamRuleResponseBody) SetRequestId(v string) *DeleteAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
