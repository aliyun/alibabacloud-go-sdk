// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnWafRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteDcdnWafRulesResponseBody
	GetRequestId() *string
}

type BatchDeleteDcdnWafRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-802B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteDcdnWafRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnWafRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteDcdnWafRulesResponseBody) SetRequestId(v string) *BatchDeleteDcdnWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDcdnWafRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
