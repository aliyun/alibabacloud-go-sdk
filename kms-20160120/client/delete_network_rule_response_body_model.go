// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkRuleResponseBody
	GetRequestId() *string
}

type DeleteNetworkRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f93-be0f-cc043fda2d4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkRuleResponseBody) SetRequestId(v string) *DeleteNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
