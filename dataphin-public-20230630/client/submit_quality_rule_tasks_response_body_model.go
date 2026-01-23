// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityRuleTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitQualityRuleTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SubmitQualityRuleTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitQualityRuleTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitQualityRuleTasksResponseBody
	GetRequestId() *string
	SetSubmitResult(v *SubmitQualityRuleTasksResponseBodySubmitResult) *SubmitQualityRuleTasksResponseBody
	GetSubmitResult() *SubmitQualityRuleTasksResponseBodySubmitResult
	SetSuccess(v bool) *SubmitQualityRuleTasksResponseBody
	GetSuccess() *bool
}

type SubmitQualityRuleTasksResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubmitResult *SubmitQualityRuleTasksResponseBodySubmitResult `json:"SubmitResult,omitempty" xml:"SubmitResult,omitempty" type:"Struct"`
	Success      *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitQualityRuleTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityRuleTasksResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitQualityRuleTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitQualityRuleTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitQualityRuleTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitQualityRuleTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitQualityRuleTasksResponseBody) GetSubmitResult() *SubmitQualityRuleTasksResponseBodySubmitResult {
	return s.SubmitResult
}

func (s *SubmitQualityRuleTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitQualityRuleTasksResponseBody) SetCode(v string) *SubmitQualityRuleTasksResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitQualityRuleTasksResponseBody) SetHttpStatusCode(v int32) *SubmitQualityRuleTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitQualityRuleTasksResponseBody) SetMessage(v string) *SubmitQualityRuleTasksResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitQualityRuleTasksResponseBody) SetRequestId(v string) *SubmitQualityRuleTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitQualityRuleTasksResponseBody) SetSubmitResult(v *SubmitQualityRuleTasksResponseBodySubmitResult) *SubmitQualityRuleTasksResponseBody {
	s.SubmitResult = v
	return s
}

func (s *SubmitQualityRuleTasksResponseBody) SetSuccess(v bool) *SubmitQualityRuleTasksResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitQualityRuleTasksResponseBody) Validate() error {
	if s.SubmitResult != nil {
		if err := s.SubmitResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitQualityRuleTasksResponseBodySubmitResult struct {
	RuleTaskIdList  []*int64 `json:"RuleTaskIdList,omitempty" xml:"RuleTaskIdList,omitempty" type:"Repeated"`
	WatchTaskIdList []*int64 `json:"WatchTaskIdList,omitempty" xml:"WatchTaskIdList,omitempty" type:"Repeated"`
}

func (s SubmitQualityRuleTasksResponseBodySubmitResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityRuleTasksResponseBodySubmitResult) GoString() string {
	return s.String()
}

func (s *SubmitQualityRuleTasksResponseBodySubmitResult) GetRuleTaskIdList() []*int64 {
	return s.RuleTaskIdList
}

func (s *SubmitQualityRuleTasksResponseBodySubmitResult) GetWatchTaskIdList() []*int64 {
	return s.WatchTaskIdList
}

func (s *SubmitQualityRuleTasksResponseBodySubmitResult) SetRuleTaskIdList(v []*int64) *SubmitQualityRuleTasksResponseBodySubmitResult {
	s.RuleTaskIdList = v
	return s
}

func (s *SubmitQualityRuleTasksResponseBodySubmitResult) SetWatchTaskIdList(v []*int64) *SubmitQualityRuleTasksResponseBodySubmitResult {
	s.WatchTaskIdList = v
	return s
}

func (s *SubmitQualityRuleTasksResponseBodySubmitResult) Validate() error {
	return dara.Validate(s)
}
