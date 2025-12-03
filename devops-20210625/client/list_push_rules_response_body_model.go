// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPushRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListPushRulesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPushRulesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListPushRulesResponseBody
	GetRequestId() *string
	SetResult(v []*ListPushRulesResponseBodyResult) *ListPushRulesResponseBody
	GetResult() []*ListPushRulesResponseBodyResult
	SetSuccess(v bool) *ListPushRulesResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListPushRulesResponseBody
	GetTotal() *int64
}

type ListPushRulesResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListPushRulesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPushRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPushRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPushRulesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPushRulesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPushRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPushRulesResponseBody) GetResult() []*ListPushRulesResponseBodyResult {
	return s.Result
}

func (s *ListPushRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPushRulesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListPushRulesResponseBody) SetErrorCode(v string) *ListPushRulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPushRulesResponseBody) SetErrorMessage(v string) *ListPushRulesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPushRulesResponseBody) SetRequestId(v string) *ListPushRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPushRulesResponseBody) SetResult(v []*ListPushRulesResponseBodyResult) *ListPushRulesResponseBody {
	s.Result = v
	return s
}

func (s *ListPushRulesResponseBody) SetSuccess(v bool) *ListPushRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPushRulesResponseBody) SetTotal(v int64) *ListPushRulesResponseBody {
	s.Total = &v
	return s
}

func (s *ListPushRulesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPushRulesResponseBodyResult struct {
	// example:
	//
	// 2023-09-03T18:20:06+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2023-09-03T18:20:06+08:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 2077
	Id        *int64                                      `json:"id,omitempty" xml:"id,omitempty"`
	RuleInfos []*ListPushRulesResponseBodyResultRuleInfos `json:"ruleInfos,omitempty" xml:"ruleInfos,omitempty" type:"Repeated"`
}

func (s ListPushRulesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListPushRulesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPushRulesResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListPushRulesResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListPushRulesResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListPushRulesResponseBodyResult) GetRuleInfos() []*ListPushRulesResponseBodyResultRuleInfos {
	return s.RuleInfos
}

func (s *ListPushRulesResponseBodyResult) SetGmtCreate(v string) *ListPushRulesResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListPushRulesResponseBodyResult) SetGmtModified(v string) *ListPushRulesResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListPushRulesResponseBodyResult) SetId(v int64) *ListPushRulesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListPushRulesResponseBodyResult) SetRuleInfos(v []*ListPushRulesResponseBodyResultRuleInfos) *ListPushRulesResponseBodyResult {
	s.RuleInfos = v
	return s
}

func (s *ListPushRulesResponseBodyResult) Validate() error {
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

type ListPushRulesResponseBodyResultRuleInfos struct {
	// example:
	//
	// CommitFilesChecker
	CheckerName *string `json:"checkerName,omitempty" xml:"checkerName,omitempty"`
	// example:
	//
	// warn
	CheckerType *string `json:"checkerType,omitempty" xml:"checkerType,omitempty"`
	// example:
	//
	// ""
	ExtraMessage    *string   `json:"extraMessage,omitempty" xml:"extraMessage,omitempty"`
	FileRuleRegexes []*string `json:"fileRuleRegexes,omitempty" xml:"fileRuleRegexes,omitempty" type:"Repeated"`
}

func (s ListPushRulesResponseBodyResultRuleInfos) String() string {
	return dara.Prettify(s)
}

func (s ListPushRulesResponseBodyResultRuleInfos) GoString() string {
	return s.String()
}

func (s *ListPushRulesResponseBodyResultRuleInfos) GetCheckerName() *string {
	return s.CheckerName
}

func (s *ListPushRulesResponseBodyResultRuleInfos) GetCheckerType() *string {
	return s.CheckerType
}

func (s *ListPushRulesResponseBodyResultRuleInfos) GetExtraMessage() *string {
	return s.ExtraMessage
}

func (s *ListPushRulesResponseBodyResultRuleInfos) GetFileRuleRegexes() []*string {
	return s.FileRuleRegexes
}

func (s *ListPushRulesResponseBodyResultRuleInfos) SetCheckerName(v string) *ListPushRulesResponseBodyResultRuleInfos {
	s.CheckerName = &v
	return s
}

func (s *ListPushRulesResponseBodyResultRuleInfos) SetCheckerType(v string) *ListPushRulesResponseBodyResultRuleInfos {
	s.CheckerType = &v
	return s
}

func (s *ListPushRulesResponseBodyResultRuleInfos) SetExtraMessage(v string) *ListPushRulesResponseBodyResultRuleInfos {
	s.ExtraMessage = &v
	return s
}

func (s *ListPushRulesResponseBodyResultRuleInfos) SetFileRuleRegexes(v []*string) *ListPushRulesResponseBodyResultRuleInfos {
	s.FileRuleRegexes = v
	return s
}

func (s *ListPushRulesResponseBodyResultRuleInfos) Validate() error {
	return dara.Validate(s)
}
