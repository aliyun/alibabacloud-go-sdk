// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRuleResponseBody
	GetRequestId() *string
}

type ModifyRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRuleResponseBody) SetRequestId(v string) *ModifyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
