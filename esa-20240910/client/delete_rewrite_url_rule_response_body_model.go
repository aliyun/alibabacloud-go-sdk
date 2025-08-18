// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRewriteUrlRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRewriteUrlRuleResponseBody
	GetRequestId() *string
}

type DeleteRewriteUrlRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 35C66C7B-671H-4297-9187-2C4477247A78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRewriteUrlRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRewriteUrlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRewriteUrlRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRewriteUrlRuleResponseBody) SetRequestId(v string) *DeleteRewriteUrlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRewriteUrlRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
