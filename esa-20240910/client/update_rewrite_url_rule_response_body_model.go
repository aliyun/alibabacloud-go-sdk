// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRewriteUrlRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRewriteUrlRuleResponseBody
	GetRequestId() *string
}

type UpdateRewriteUrlRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRewriteUrlRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRewriteUrlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRewriteUrlRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRewriteUrlRuleResponseBody) SetRequestId(v string) *UpdateRewriteUrlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRewriteUrlRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
