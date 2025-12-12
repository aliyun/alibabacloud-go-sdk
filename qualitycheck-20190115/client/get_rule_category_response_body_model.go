// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRuleCategoryResponseBody
	GetCode() *string
	SetData(v *GetRuleCategoryResponseBodyData) *GetRuleCategoryResponseBody
	GetData() *GetRuleCategoryResponseBodyData
	SetMessage(v string) *GetRuleCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRuleCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRuleCategoryResponseBody
	GetSuccess() *bool
}

type GetRuleCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRuleCategoryResponseBody) GetData() *GetRuleCategoryResponseBodyData {
	return s.Data
}

func (s *GetRuleCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuleCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRuleCategoryResponseBody) SetCode(v string) *GetRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetData(v *GetRuleCategoryResponseBodyData) *GetRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleCategoryResponseBody) SetMessage(v string) *GetRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetRequestId(v string) *GetRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetSuccess(v bool) *GetRuleCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetRuleCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleCategoryResponseBodyData struct {
	RuleCountInfo []*GetRuleCategoryResponseBodyDataRuleCountInfo `json:"RuleCountInfo,omitempty" xml:"RuleCountInfo,omitempty" type:"Repeated"`
}

func (s GetRuleCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBodyData) GetRuleCountInfo() []*GetRuleCategoryResponseBodyDataRuleCountInfo {
	return s.RuleCountInfo
}

func (s *GetRuleCategoryResponseBodyData) SetRuleCountInfo(v []*GetRuleCategoryResponseBodyDataRuleCountInfo) *GetRuleCategoryResponseBodyData {
	s.RuleCountInfo = v
	return s
}

func (s *GetRuleCategoryResponseBodyData) Validate() error {
	if s.RuleCountInfo != nil {
		for _, item := range s.RuleCountInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuleCategoryResponseBodyDataRuleCountInfo struct {
	// example:
	//
	// false
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
	// example:
	//
	// 22
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetRuleCategoryResponseBodyDataRuleCountInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRuleCategoryResponseBodyDataRuleCountInfo) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) GetSelect() *bool {
	return s.Select
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) GetType() *int32 {
	return s.Type
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) GetTypeName() *string {
	return s.TypeName
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetSelect(v bool) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.Select = &v
	return s
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetType(v int32) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.Type = &v
	return s
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetTypeName(v string) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.TypeName = &v
	return s
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) Validate() error {
	return dara.Validate(s)
}
