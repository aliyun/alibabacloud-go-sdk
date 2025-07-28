// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseRuleStatusResponseBody
	GetRequestId() *string
}

type ModifyDefenseRuleStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BB305BF3-3C71-57A9-9704-E22F567689B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseRuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseRuleStatusResponseBody) SetRequestId(v string) *ModifyDefenseRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseRuleStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
