// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchModifyDcdnWafRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchModifyDcdnWafRulesResponseBody
	GetRequestId() *string
}

type BatchModifyDcdnWafRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-XXXX-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchModifyDcdnWafRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchModifyDcdnWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchModifyDcdnWafRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchModifyDcdnWafRulesResponseBody) SetRequestId(v string) *BatchModifyDcdnWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchModifyDcdnWafRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
