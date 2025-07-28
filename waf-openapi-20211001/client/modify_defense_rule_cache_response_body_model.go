// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseRuleCacheResponseBody
	GetRequestId() *string
}

type ModifyDefenseRuleCacheResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A0F2B994-8645-5270-A05D-9DAD8C****B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseRuleCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleCacheResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseRuleCacheResponseBody) SetRequestId(v string) *ModifyDefenseRuleCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseRuleCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
