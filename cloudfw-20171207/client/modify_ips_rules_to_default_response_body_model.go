// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpsRulesToDefaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpsRulesToDefaultResponseBody
	GetRequestId() *string
}

type ModifyIpsRulesToDefaultResponseBody struct {
	// example:
	//
	// B713361D-62E2-5FF0-9D29-BBFAAF40****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpsRulesToDefaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpsRulesToDefaultResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpsRulesToDefaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpsRulesToDefaultResponseBody) SetRequestId(v string) *ModifyIpsRulesToDefaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpsRulesToDefaultResponseBody) Validate() error {
	return dara.Validate(s)
}
