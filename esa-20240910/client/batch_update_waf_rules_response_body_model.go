// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateWafRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUpdateWafRulesResponseBody
	GetRequestId() *string
}

type BatchUpdateWafRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUpdateWafRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateWafRulesResponseBody) SetRequestId(v string) *BatchUpdateWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateWafRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
