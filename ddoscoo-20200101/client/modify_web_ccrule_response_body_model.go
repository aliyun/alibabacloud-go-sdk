// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCCRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebCCRuleResponseBody
	GetRequestId() *string
}

type ModifyWebCCRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCCRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCCRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebCCRuleResponseBody) SetRequestId(v string) *ModifyWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebCCRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
