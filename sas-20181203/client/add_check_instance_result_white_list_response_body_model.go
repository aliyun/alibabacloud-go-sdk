// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCheckInstanceResultWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *AddCheckInstanceResultWhiteListResponseBody
	GetData() map[string]interface{}
	SetRequestId(v string) *AddCheckInstanceResultWhiteListResponseBody
	GetRequestId() *string
	SetRuleId(v string) *AddCheckInstanceResultWhiteListResponseBody
	GetRuleId() *string
}

type AddCheckInstanceResultWhiteListResponseBody struct {
	// The data returned. This parameter is deprecated.
	//
	// example:
	//
	// xxx
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ADE57832-9666-511C-9A80-B87DE2E8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the whitelist rule.
	//
	// example:
	//
	// 381049
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s AddCheckInstanceResultWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCheckInstanceResultWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddCheckInstanceResultWhiteListResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *AddCheckInstanceResultWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCheckInstanceResultWhiteListResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *AddCheckInstanceResultWhiteListResponseBody) SetData(v map[string]interface{}) *AddCheckInstanceResultWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *AddCheckInstanceResultWhiteListResponseBody) SetRequestId(v string) *AddCheckInstanceResultWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListResponseBody) SetRuleId(v string) *AddCheckInstanceResultWhiteListResponseBody {
	s.RuleId = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
