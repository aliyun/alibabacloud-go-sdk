// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpIncomingResponseHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type DeleteHttpIncomingResponseHeaderModificationRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHttpIncomingResponseHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpIncomingResponseHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody) SetRequestId(v string) *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
