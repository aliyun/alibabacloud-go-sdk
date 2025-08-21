// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCCRuleV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWebCCRuleV2ResponseBody
	GetRequestId() *string
}

type DeleteWebCCRuleV2ResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6D48AED0-41DB-5D9B-B484-3B6AAD312AD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebCCRuleV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCCRuleV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebCCRuleV2ResponseBody) SetRequestId(v string) *DeleteWebCCRuleV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebCCRuleV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
