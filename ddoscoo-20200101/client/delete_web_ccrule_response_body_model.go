// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCCRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWebCCRuleResponseBody
	GetRequestId() *string
}

type DeleteWebCCRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebCCRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebCCRuleResponseBody) SetRequestId(v string) *DeleteWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebCCRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
