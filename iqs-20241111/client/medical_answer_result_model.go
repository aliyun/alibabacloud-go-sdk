// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAnswerResult interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*MedicalAnswerMessage) *MedicalAnswerResult
	GetMessages() []*MedicalAnswerMessage
	SetQueryContext(v *MultimodalQueryContext) *MedicalAnswerResult
	GetQueryContext() *MultimodalQueryContext
	SetRequestId(v string) *MedicalAnswerResult
	GetRequestId() *string
	SetSearchInformation(v *UnifiedSearchInformation) *MedicalAnswerResult
	GetSearchInformation() *UnifiedSearchInformation
}

type MedicalAnswerResult struct {
	Messages          []*MedicalAnswerMessage   `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	QueryContext      *MultimodalQueryContext   `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	RequestId         *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SearchInformation *UnifiedSearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
}

func (s MedicalAnswerResult) String() string {
	return dara.Prettify(s)
}

func (s MedicalAnswerResult) GoString() string {
	return s.String()
}

func (s *MedicalAnswerResult) GetMessages() []*MedicalAnswerMessage {
	return s.Messages
}

func (s *MedicalAnswerResult) GetQueryContext() *MultimodalQueryContext {
	return s.QueryContext
}

func (s *MedicalAnswerResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAnswerResult) GetSearchInformation() *UnifiedSearchInformation {
	return s.SearchInformation
}

func (s *MedicalAnswerResult) SetMessages(v []*MedicalAnswerMessage) *MedicalAnswerResult {
	s.Messages = v
	return s
}

func (s *MedicalAnswerResult) SetQueryContext(v *MultimodalQueryContext) *MedicalAnswerResult {
	s.QueryContext = v
	return s
}

func (s *MedicalAnswerResult) SetRequestId(v string) *MedicalAnswerResult {
	s.RequestId = &v
	return s
}

func (s *MedicalAnswerResult) SetSearchInformation(v *UnifiedSearchInformation) *MedicalAnswerResult {
	s.SearchInformation = v
	return s
}

func (s *MedicalAnswerResult) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryContext != nil {
		if err := s.QueryContext.Validate(); err != nil {
			return err
		}
	}
	if s.SearchInformation != nil {
		if err := s.SearchInformation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
