// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpIncomingRequestHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type DeleteHttpIncomingRequestHeaderModificationRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHttpIncomingRequestHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpIncomingRequestHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody) SetRequestId(v string) *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
