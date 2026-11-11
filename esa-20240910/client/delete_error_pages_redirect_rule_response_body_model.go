// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErrorPagesRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteErrorPagesRedirectRuleResponseBody
	GetRequestId() *string
}

type DeleteErrorPagesRedirectRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 39572251-FB3A-5A40-A6E1-38A7AF7E2D5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErrorPagesRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteErrorPagesRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErrorPagesRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteErrorPagesRedirectRuleResponseBody) SetRequestId(v string) *DeleteErrorPagesRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteErrorPagesRedirectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
