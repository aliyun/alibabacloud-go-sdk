// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseRuleResponseBody
	GetRequestId() *string
}

type ModifyDefenseRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1EEA9C98-F166-54FE-ADE3-08D8****BDFA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseRuleResponseBody) SetRequestId(v string) *ModifyDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
