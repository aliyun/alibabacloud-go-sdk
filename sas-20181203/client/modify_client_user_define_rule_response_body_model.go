// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientUserDefineRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClientUserDefineRuleResponseBody
	GetRequestId() *string
}

type ModifyClientUserDefineRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5cbb3c39-88ec-429a-be26-5d0f62cc****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClientUserDefineRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientUserDefineRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClientUserDefineRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClientUserDefineRuleResponseBody) SetRequestId(v string) *ModifyClientUserDefineRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClientUserDefineRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
