// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDcdnWafPolicyResponseBody
	GetRequestId() *string
}

type ModifyDcdnWafPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-C730-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDcdnWafPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDcdnWafPolicyResponseBody) SetRequestId(v string) *ModifyDcdnWafPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDcdnWafPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
