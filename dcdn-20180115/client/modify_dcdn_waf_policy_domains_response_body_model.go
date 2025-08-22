// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafPolicyDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDcdnWafPolicyDomainsResponseBody
	GetRequestId() *string
}

type ModifyDcdnWafPolicyDomainsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-2B35-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDcdnWafPolicyDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafPolicyDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafPolicyDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDcdnWafPolicyDomainsResponseBody) SetRequestId(v string) *ModifyDcdnWafPolicyDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDcdnWafPolicyDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
