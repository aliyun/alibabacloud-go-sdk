// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePushRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdatePushRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdatePushRuleResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdatePushRuleResponseBody
	GetRequestId() *string
	SetResult(v *UpdatePushRuleResponseBodyResult) *UpdatePushRuleResponseBody
	GetResult() *UpdatePushRuleResponseBodyResult
	SetSuccess(v bool) *UpdatePushRuleResponseBody
	GetSuccess() *bool
}

type UpdatePushRuleResponseBody struct {
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
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdatePushRuleResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePushRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePushRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdatePushRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdatePushRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePushRuleResponseBody) GetResult() *UpdatePushRuleResponseBodyResult {
	return s.Result
}

func (s *UpdatePushRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePushRuleResponseBody) SetErrorCode(v string) *UpdatePushRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePushRuleResponseBody) SetErrorMessage(v string) *UpdatePushRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePushRuleResponseBody) SetRequestId(v string) *UpdatePushRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePushRuleResponseBody) SetResult(v *UpdatePushRuleResponseBodyResult) *UpdatePushRuleResponseBody {
	s.Result = v
	return s
}

func (s *UpdatePushRuleResponseBody) SetSuccess(v bool) *UpdatePushRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePushRuleResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePushRuleResponseBodyResult struct {
	// example:
	//
	// 2023-09-03T18:20:06+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2023-09-03T18:20:06+08:00
	GmtModified *string                                      `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id          *int64                                       `json:"id,omitempty" xml:"id,omitempty"`
	RuleInfos   []*UpdatePushRuleResponseBodyResultRuleInfos `json:"ruleInfos,omitempty" xml:"ruleInfos,omitempty" type:"Repeated"`
}

func (s UpdatePushRuleResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePushRuleResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdatePushRuleResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdatePushRuleResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *UpdatePushRuleResponseBodyResult) GetRuleInfos() []*UpdatePushRuleResponseBodyResultRuleInfos {
	return s.RuleInfos
}

func (s *UpdatePushRuleResponseBodyResult) SetGmtCreate(v string) *UpdatePushRuleResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *UpdatePushRuleResponseBodyResult) SetGmtModified(v string) *UpdatePushRuleResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *UpdatePushRuleResponseBodyResult) SetId(v int64) *UpdatePushRuleResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdatePushRuleResponseBodyResult) SetRuleInfos(v []*UpdatePushRuleResponseBodyResultRuleInfos) *UpdatePushRuleResponseBodyResult {
	s.RuleInfos = v
	return s
}

func (s *UpdatePushRuleResponseBodyResult) Validate() error {
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

type UpdatePushRuleResponseBodyResultRuleInfos struct {
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

func (s UpdatePushRuleResponseBodyResultRuleInfos) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushRuleResponseBodyResultRuleInfos) GoString() string {
	return s.String()
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) GetCheckerName() *string {
	return s.CheckerName
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) GetCheckerType() *string {
	return s.CheckerType
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) GetExtraMessage() *string {
	return s.ExtraMessage
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) GetFileRuleRegexes() []*string {
	return s.FileRuleRegexes
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) SetCheckerName(v string) *UpdatePushRuleResponseBodyResultRuleInfos {
	s.CheckerName = &v
	return s
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) SetCheckerType(v string) *UpdatePushRuleResponseBodyResultRuleInfos {
	s.CheckerType = &v
	return s
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) SetExtraMessage(v string) *UpdatePushRuleResponseBodyResultRuleInfos {
	s.ExtraMessage = &v
	return s
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) SetFileRuleRegexes(v []*string) *UpdatePushRuleResponseBodyResultRuleInfos {
	s.FileRuleRegexes = v
	return s
}

func (s *UpdatePushRuleResponseBodyResultRuleInfos) Validate() error {
	return dara.Validate(s)
}
