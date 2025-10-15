// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQaLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParseQaResults(v []*UpdateQaLibraryRequestParseQaResults) *UpdateQaLibraryRequest
	GetParseQaResults() []*UpdateQaLibraryRequestParseQaResults
	SetQaLibraryId(v string) *UpdateQaLibraryRequest
	GetQaLibraryId() *string
	SetRequestId(v string) *UpdateQaLibraryRequest
	GetRequestId() *string
}

type UpdateQaLibraryRequest struct {
	// This parameter is required.
	ParseQaResults []*UpdateQaLibraryRequestParseQaResults `json:"parseQaResults,omitempty" xml:"parseQaResults,omitempty" type:"Repeated"`
	// example:
	//
	// 6jh378d
	QaLibraryId *string `json:"qaLibraryId,omitempty" xml:"qaLibraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQaLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQaLibraryRequest) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryRequest) GetParseQaResults() []*UpdateQaLibraryRequestParseQaResults {
	return s.ParseQaResults
}

func (s *UpdateQaLibraryRequest) GetQaLibraryId() *string {
	return s.QaLibraryId
}

func (s *UpdateQaLibraryRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQaLibraryRequest) SetParseQaResults(v []*UpdateQaLibraryRequestParseQaResults) *UpdateQaLibraryRequest {
	s.ParseQaResults = v
	return s
}

func (s *UpdateQaLibraryRequest) SetQaLibraryId(v string) *UpdateQaLibraryRequest {
	s.QaLibraryId = &v
	return s
}

func (s *UpdateQaLibraryRequest) SetRequestId(v string) *UpdateQaLibraryRequest {
	s.RequestId = &v
	return s
}

func (s *UpdateQaLibraryRequest) Validate() error {
	if s.ParseQaResults != nil {
		for _, item := range s.ParseQaResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateQaLibraryRequestParseQaResults struct {
	// This parameter is required.
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// This parameter is required.
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
}

func (s UpdateQaLibraryRequestParseQaResults) String() string {
	return dara.Prettify(s)
}

func (s UpdateQaLibraryRequestParseQaResults) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryRequestParseQaResults) GetAnswer() *string {
	return s.Answer
}

func (s *UpdateQaLibraryRequestParseQaResults) GetQuestion() *string {
	return s.Question
}

func (s *UpdateQaLibraryRequestParseQaResults) SetAnswer(v string) *UpdateQaLibraryRequestParseQaResults {
	s.Answer = &v
	return s
}

func (s *UpdateQaLibraryRequestParseQaResults) SetQuestion(v string) *UpdateQaLibraryRequestParseQaResults {
	s.Question = &v
	return s
}

func (s *UpdateQaLibraryRequestParseQaResults) Validate() error {
	return dara.Validate(s)
}
