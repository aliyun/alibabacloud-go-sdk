// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebCCRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableWebCCRuleResponseBody
	GetRequestId() *string
}

type DisableWebCCRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableWebCCRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWebCCRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableWebCCRuleResponseBody) SetRequestId(v string) *DisableWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableWebCCRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
