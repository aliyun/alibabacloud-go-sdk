// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpsRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpsRulesResponseBody
	GetRequestId() *string
}

type ModifyIpsRulesResponseBody struct {
	// example:
	//
	// 30FB7F84-1FC5-5A3D-BBBE-5779FC74****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpsRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpsRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpsRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpsRulesResponseBody) SetRequestId(v string) *ModifyIpsRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpsRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
