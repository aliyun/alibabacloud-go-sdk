// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpResponseHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHttpResponseHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type DeleteHttpResponseHeaderModificationRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHttpResponseHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpResponseHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpResponseHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpResponseHeaderModificationRuleResponseBody) SetRequestId(v string) *DeleteHttpResponseHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpResponseHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
