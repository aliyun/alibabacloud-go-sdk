// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAiAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiAnalysisResultList(v *GetMediaAiAnalysisResponseBodyAiAnalysisResultList) *GetMediaAiAnalysisResponseBody
	GetAiAnalysisResultList() *GetMediaAiAnalysisResponseBodyAiAnalysisResultList
	SetRequestId(v string) *GetMediaAiAnalysisResponseBody
	GetRequestId() *string
}

type GetMediaAiAnalysisResponseBody struct {
	AiAnalysisResultList *GetMediaAiAnalysisResponseBodyAiAnalysisResultList `json:"AiAnalysisResultList,omitempty" xml:"AiAnalysisResultList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 746FFA07-8BBB-46B1-3E94E3B2915E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAiAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAiAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAiAnalysisResponseBody) GetAiAnalysisResultList() *GetMediaAiAnalysisResponseBodyAiAnalysisResultList {
	return s.AiAnalysisResultList
}

func (s *GetMediaAiAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaAiAnalysisResponseBody) SetAiAnalysisResultList(v *GetMediaAiAnalysisResponseBodyAiAnalysisResultList) *GetMediaAiAnalysisResponseBody {
	s.AiAnalysisResultList = v
	return s
}

func (s *GetMediaAiAnalysisResponseBody) SetRequestId(v string) *GetMediaAiAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaAiAnalysisResponseBody) Validate() error {
	if s.AiAnalysisResultList != nil {
		if err := s.AiAnalysisResultList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaAiAnalysisResponseBodyAiAnalysisResultList struct {
	AiAnalysisResult []*GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult `json:"AiAnalysisResult,omitempty" xml:"AiAnalysisResult,omitempty" type:"Repeated"`
}

func (s GetMediaAiAnalysisResponseBodyAiAnalysisResultList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAiAnalysisResponseBodyAiAnalysisResultList) GoString() string {
	return s.String()
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultList) GetAiAnalysisResult() []*GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult {
	return s.AiAnalysisResult
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultList) SetAiAnalysisResult(v []*GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) *GetMediaAiAnalysisResponseBodyAiAnalysisResultList {
	s.AiAnalysisResult = v
	return s
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultList) Validate() error {
	if s.AiAnalysisResult != nil {
		for _, item := range s.AiAnalysisResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Extra      *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	Summary    *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) GoString() string {
	return s.String()
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) GetContent() *string {
	return s.Content
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) GetExtra() *string {
	return s.Extra
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) GetResultType() *string {
	return s.ResultType
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) GetSummary() *string {
	return s.Summary
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) GetTitle() *string {
	return s.Title
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) SetContent(v string) *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult {
	s.Content = &v
	return s
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) SetExtra(v string) *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult {
	s.Extra = &v
	return s
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) SetResultType(v string) *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult {
	s.ResultType = &v
	return s
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) SetSummary(v string) *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult {
	s.Summary = &v
	return s
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) SetTitle(v string) *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult {
	s.Title = &v
	return s
}

func (s *GetMediaAiAnalysisResponseBodyAiAnalysisResultListAiAnalysisResult) Validate() error {
	return dara.Validate(s)
}
