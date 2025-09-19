// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCheckResultWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *AddCheckResultWhiteListResponseBody
	GetData() map[string]interface{}
	SetRequestId(v string) *AddCheckResultWhiteListResponseBody
	GetRequestId() *string
	SetRuleIds(v []*int64) *AddCheckResultWhiteListResponseBody
	GetRuleIds() []*int64
}

type AddCheckResultWhiteListResponseBody struct {
	// The data returned. This parameter is deprecated.
	//
	// example:
	//
	// 1
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C699E4E4-F2F4-58FC-A949-457FFE59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the whitelist rules that are generated.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s AddCheckResultWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCheckResultWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddCheckResultWhiteListResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *AddCheckResultWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCheckResultWhiteListResponseBody) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *AddCheckResultWhiteListResponseBody) SetData(v map[string]interface{}) *AddCheckResultWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *AddCheckResultWhiteListResponseBody) SetRequestId(v string) *AddCheckResultWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCheckResultWhiteListResponseBody) SetRuleIds(v []*int64) *AddCheckResultWhiteListResponseBody {
	s.RuleIds = v
	return s
}

func (s *AddCheckResultWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
