// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetRuleResponseBody
	GetRequestId() *string
}

type SetRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SetRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRuleResponseBody) SetRequestId(v string) *SetRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
