// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteACLRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteACLRuleResponseBody
	GetRequestId() *string
}

type DeleteACLRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 39E71162-699A-4E02-AF0F-17621BA6AEF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteACLRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteACLRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteACLRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteACLRuleResponseBody) SetRequestId(v string) *DeleteACLRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteACLRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
