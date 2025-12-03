// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePushRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreatePushRuleRequest
	GetAccessToken() *string
	SetRuleInfos(v []*CreatePushRuleRequestRuleInfos) *CreatePushRuleRequest
	GetRuleInfos() []*CreatePushRuleRequestRuleInfos
	SetOrganizationId(v string) *CreatePushRuleRequest
	GetOrganizationId() *string
}

type CreatePushRuleRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	RuleInfos []*CreatePushRuleRequestRuleInfos `json:"ruleInfos,omitempty" xml:"ruleInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreatePushRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePushRuleRequest) GoString() string {
	return s.String()
}

func (s *CreatePushRuleRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreatePushRuleRequest) GetRuleInfos() []*CreatePushRuleRequestRuleInfos {
	return s.RuleInfos
}

func (s *CreatePushRuleRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreatePushRuleRequest) SetAccessToken(v string) *CreatePushRuleRequest {
	s.AccessToken = &v
	return s
}

func (s *CreatePushRuleRequest) SetRuleInfos(v []*CreatePushRuleRequestRuleInfos) *CreatePushRuleRequest {
	s.RuleInfos = v
	return s
}

func (s *CreatePushRuleRequest) SetOrganizationId(v string) *CreatePushRuleRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreatePushRuleRequest) Validate() error {
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

type CreatePushRuleRequestRuleInfos struct {
	// example:
	//
	// CommitMessageChecker
	CheckerName *string `json:"checkerName,omitempty" xml:"checkerName,omitempty"`
	// example:
	//
	// warn
	CheckerType *string `json:"checkerType,omitempty" xml:"checkerType,omitempty"`
	// example:
	//
	// user@example.com
	ExtraMessage    *string   `json:"extraMessage,omitempty" xml:"extraMessage,omitempty"`
	FileRuleRegexes []*string `json:"fileRuleRegexes,omitempty" xml:"fileRuleRegexes,omitempty" type:"Repeated"`
}

func (s CreatePushRuleRequestRuleInfos) String() string {
	return dara.Prettify(s)
}

func (s CreatePushRuleRequestRuleInfos) GoString() string {
	return s.String()
}

func (s *CreatePushRuleRequestRuleInfos) GetCheckerName() *string {
	return s.CheckerName
}

func (s *CreatePushRuleRequestRuleInfos) GetCheckerType() *string {
	return s.CheckerType
}

func (s *CreatePushRuleRequestRuleInfos) GetExtraMessage() *string {
	return s.ExtraMessage
}

func (s *CreatePushRuleRequestRuleInfos) GetFileRuleRegexes() []*string {
	return s.FileRuleRegexes
}

func (s *CreatePushRuleRequestRuleInfos) SetCheckerName(v string) *CreatePushRuleRequestRuleInfos {
	s.CheckerName = &v
	return s
}

func (s *CreatePushRuleRequestRuleInfos) SetCheckerType(v string) *CreatePushRuleRequestRuleInfos {
	s.CheckerType = &v
	return s
}

func (s *CreatePushRuleRequestRuleInfos) SetExtraMessage(v string) *CreatePushRuleRequestRuleInfos {
	s.ExtraMessage = &v
	return s
}

func (s *CreatePushRuleRequestRuleInfos) SetFileRuleRegexes(v []*string) *CreatePushRuleRequestRuleInfos {
	s.FileRuleRegexes = v
	return s
}

func (s *CreatePushRuleRequestRuleInfos) Validate() error {
	return dara.Validate(s)
}
