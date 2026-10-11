// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyRuleGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEgressProxyRuleGroupResponseBody
	GetRequestId() *string
}

type DeleteEgressProxyRuleGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 019F3A67-2D1B-537B-B08D-B2DEB69E8E8B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEgressProxyRuleGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyRuleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyRuleGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEgressProxyRuleGroupResponseBody) SetRequestId(v string) *DeleteEgressProxyRuleGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEgressProxyRuleGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
