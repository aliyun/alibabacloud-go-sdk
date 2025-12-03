// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePushRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreatePushRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreatePushRuleResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreatePushRuleResponseBody
	GetRequestId() *string
	SetResult(v *CreatePushRuleResponseBodyResult) *CreatePushRuleResponseBody
	GetResult() *CreatePushRuleResponseBodyResult
	SetSuccess(v bool) *CreatePushRuleResponseBody
	GetSuccess() *bool
}

type CreatePushRuleResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreatePushRuleResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreatePushRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePushRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePushRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreatePushRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreatePushRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePushRuleResponseBody) GetResult() *CreatePushRuleResponseBodyResult {
	return s.Result
}

func (s *CreatePushRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePushRuleResponseBody) SetErrorCode(v string) *CreatePushRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreatePushRuleResponseBody) SetErrorMessage(v string) *CreatePushRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePushRuleResponseBody) SetRequestId(v string) *CreatePushRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePushRuleResponseBody) SetResult(v *CreatePushRuleResponseBodyResult) *CreatePushRuleResponseBody {
	s.Result = v
	return s
}

func (s *CreatePushRuleResponseBody) SetSuccess(v bool) *CreatePushRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePushRuleResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePushRuleResponseBodyResult struct {
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
	Id        *int64                                       `json:"id,omitempty" xml:"id,omitempty"`
	RuleInfos []*CreatePushRuleResponseBodyResultRuleInfos `json:"ruleInfos,omitempty" xml:"ruleInfos,omitempty" type:"Repeated"`
}

func (s CreatePushRuleResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreatePushRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreatePushRuleResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePushRuleResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreatePushRuleResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreatePushRuleResponseBodyResult) GetRuleInfos() []*CreatePushRuleResponseBodyResultRuleInfos {
	return s.RuleInfos
}

func (s *CreatePushRuleResponseBodyResult) SetGmtCreate(v string) *CreatePushRuleResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *CreatePushRuleResponseBodyResult) SetGmtModified(v string) *CreatePushRuleResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *CreatePushRuleResponseBodyResult) SetId(v int64) *CreatePushRuleResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreatePushRuleResponseBodyResult) SetRuleInfos(v []*CreatePushRuleResponseBodyResultRuleInfos) *CreatePushRuleResponseBodyResult {
	s.RuleInfos = v
	return s
}

func (s *CreatePushRuleResponseBodyResult) Validate() error {
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

type CreatePushRuleResponseBodyResultRuleInfos struct {
	// example:
	//
	// ForcePushChecker
	CheckerName *string `json:"checkerName,omitempty" xml:"checkerName,omitempty"`
	// example:
	//
	// warn
	CheckerType *string `json:"checkerType,omitempty" xml:"checkerType,omitempty"`
	// example:
	//
	// disabled
	ExtraMessage    *string   `json:"extraMessage,omitempty" xml:"extraMessage,omitempty"`
	FileRuleRegexes []*string `json:"fileRuleRegexes,omitempty" xml:"fileRuleRegexes,omitempty" type:"Repeated"`
}

func (s CreatePushRuleResponseBodyResultRuleInfos) String() string {
	return dara.Prettify(s)
}

func (s CreatePushRuleResponseBodyResultRuleInfos) GoString() string {
	return s.String()
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) GetCheckerName() *string {
	return s.CheckerName
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) GetCheckerType() *string {
	return s.CheckerType
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) GetExtraMessage() *string {
	return s.ExtraMessage
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) GetFileRuleRegexes() []*string {
	return s.FileRuleRegexes
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) SetCheckerName(v string) *CreatePushRuleResponseBodyResultRuleInfos {
	s.CheckerName = &v
	return s
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) SetCheckerType(v string) *CreatePushRuleResponseBodyResultRuleInfos {
	s.CheckerType = &v
	return s
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) SetExtraMessage(v string) *CreatePushRuleResponseBodyResultRuleInfos {
	s.ExtraMessage = &v
	return s
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) SetFileRuleRegexes(v []*string) *CreatePushRuleResponseBodyResultRuleInfos {
	s.FileRuleRegexes = v
	return s
}

func (s *CreatePushRuleResponseBodyResultRuleInfos) Validate() error {
	return dara.Validate(s)
}
