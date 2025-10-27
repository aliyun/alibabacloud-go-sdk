// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientUserDefineRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddClientUserDefineRuleResponseBody
	GetRequestId() *string
	SetUserDefineRuleAddResult(v *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) *AddClientUserDefineRuleResponseBody
	GetUserDefineRuleAddResult() *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult
}

type AddClientUserDefineRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 379a9b8f-107b-4630-9e95-2299a1ea****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The identifier of the custom defense rule.
	UserDefineRuleAddResult *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult `json:"UserDefineRuleAddResult,omitempty" xml:"UserDefineRuleAddResult,omitempty" type:"Struct"`
}

func (s AddClientUserDefineRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddClientUserDefineRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddClientUserDefineRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddClientUserDefineRuleResponseBody) GetUserDefineRuleAddResult() *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult {
	return s.UserDefineRuleAddResult
}

func (s *AddClientUserDefineRuleResponseBody) SetRequestId(v string) *AddClientUserDefineRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddClientUserDefineRuleResponseBody) SetUserDefineRuleAddResult(v *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) *AddClientUserDefineRuleResponseBody {
	s.UserDefineRuleAddResult = v
	return s
}

func (s *AddClientUserDefineRuleResponseBody) Validate() error {
	if s.UserDefineRuleAddResult != nil {
		if err := s.UserDefineRuleAddResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult struct {
	// The ID of the rule.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **windows**: Windows
	//
	// 	- **linux**: Linux
	//
	// 	- **all**: all types
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The switch ID of the custom defense rule.
	//
	// example:
	//
	// USER-DEFINE-RULE-SWITCH-TYPE_200****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) String() string {
	return dara.Prettify(s)
}

func (s AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) GoString() string {
	return s.String()
}

func (s *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) GetId() *int64 {
	return s.Id
}

func (s *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) GetPlatform() *string {
	return s.Platform
}

func (s *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) GetSwitchId() *string {
	return s.SwitchId
}

func (s *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) SetId(v int64) *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult {
	s.Id = &v
	return s
}

func (s *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) SetPlatform(v string) *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult {
	s.Platform = &v
	return s
}

func (s *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) SetSwitchId(v string) *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult {
	s.SwitchId = &v
	return s
}

func (s *AddClientUserDefineRuleResponseBodyUserDefineRuleAddResult) Validate() error {
	return dara.Validate(s)
}
