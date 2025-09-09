// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRuleResponseBody
	GetRequestId() *string
}

type DeleteRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
