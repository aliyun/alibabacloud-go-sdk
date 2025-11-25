// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContainerDefenseRuleDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetContainerDefenseRuleDetailRequest
	GetLang() *string
	SetRuleId(v int64) *GetContainerDefenseRuleDetailRequest
	GetRuleId() *int64
}

type GetContainerDefenseRuleDetailRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The rule ID.
	//
	// >  You can call the [ListContainerDefenseRule](https://help.aliyun.com/document_detail/2590599.html) operation to query the rule ID.
	//
	// example:
	//
	// 156
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetContainerDefenseRuleDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContainerDefenseRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetContainerDefenseRuleDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *GetContainerDefenseRuleDetailRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetContainerDefenseRuleDetailRequest) SetLang(v string) *GetContainerDefenseRuleDetailRequest {
	s.Lang = &v
	return s
}

func (s *GetContainerDefenseRuleDetailRequest) SetRuleId(v int64) *GetContainerDefenseRuleDetailRequest {
	s.RuleId = &v
	return s
}

func (s *GetContainerDefenseRuleDetailRequest) Validate() error {
	return dara.Validate(s)
}
