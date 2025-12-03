// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePushRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdatePushRuleRequest
	GetAccessToken() *string
	SetRuleInfos(v []*UpdatePushRuleRequestRuleInfos) *UpdatePushRuleRequest
	GetRuleInfos() []*UpdatePushRuleRequestRuleInfos
	SetOrganizationId(v string) *UpdatePushRuleRequest
	GetOrganizationId() *string
}

type UpdatePushRuleRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string                           `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	RuleInfos   []*UpdatePushRuleRequestRuleInfos `json:"ruleInfos,omitempty" xml:"ruleInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdatePushRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdatePushRuleRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdatePushRuleRequest) GetRuleInfos() []*UpdatePushRuleRequestRuleInfos {
	return s.RuleInfos
}

func (s *UpdatePushRuleRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdatePushRuleRequest) SetAccessToken(v string) *UpdatePushRuleRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdatePushRuleRequest) SetRuleInfos(v []*UpdatePushRuleRequestRuleInfos) *UpdatePushRuleRequest {
	s.RuleInfos = v
	return s
}

func (s *UpdatePushRuleRequest) SetOrganizationId(v string) *UpdatePushRuleRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdatePushRuleRequest) Validate() error {
	if s.RuleInfos != nil {
		for _, item := range s.RuleInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePushRuleRequestRuleInfos struct {
	// example:
	//
	// CommitAuthorChecker
	CheckerName *string `json:"checkerName,omitempty" xml:"checkerName,omitempty"`
	// example:
	//
	// warn
	CheckerType *string `json:"checkerType,omitempty" xml:"checkerType,omitempty"`
	// example:
	//
	// on
	ExtraMessage    *string   `json:"extraMessage,omitempty" xml:"extraMessage,omitempty"`
	FileRuleRegexes []*string `json:"fileRuleRegexes,omitempty" xml:"fileRuleRegexes,omitempty" type:"Repeated"`
}

func (s UpdatePushRuleRequestRuleInfos) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushRuleRequestRuleInfos) GoString() string {
	return s.String()
}

func (s *UpdatePushRuleRequestRuleInfos) GetCheckerName() *string {
	return s.CheckerName
}

func (s *UpdatePushRuleRequestRuleInfos) GetCheckerType() *string {
	return s.CheckerType
}

func (s *UpdatePushRuleRequestRuleInfos) GetExtraMessage() *string {
	return s.ExtraMessage
}

func (s *UpdatePushRuleRequestRuleInfos) GetFileRuleRegexes() []*string {
	return s.FileRuleRegexes
}

func (s *UpdatePushRuleRequestRuleInfos) SetCheckerName(v string) *UpdatePushRuleRequestRuleInfos {
	s.CheckerName = &v
	return s
}

func (s *UpdatePushRuleRequestRuleInfos) SetCheckerType(v string) *UpdatePushRuleRequestRuleInfos {
	s.CheckerType = &v
	return s
}

func (s *UpdatePushRuleRequestRuleInfos) SetExtraMessage(v string) *UpdatePushRuleRequestRuleInfos {
	s.ExtraMessage = &v
	return s
}

func (s *UpdatePushRuleRequestRuleInfos) SetFileRuleRegexes(v []*string) *UpdatePushRuleRequestRuleInfos {
	s.FileRuleRegexes = v
	return s
}

func (s *UpdatePushRuleRequestRuleInfos) Validate() error {
	return dara.Validate(s)
}
