// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDcdnWafRuleResponseBody
	GetRequestId() *string
}

type ModifyDcdnWafRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C2B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDcdnWafRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDcdnWafRuleResponseBody) SetRequestId(v string) *ModifyDcdnWafRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDcdnWafRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
