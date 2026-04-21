// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalKnowResult interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*MedicalKnowledgeInfo) *MedicalKnowResult
	GetMessages() []*MedicalKnowledgeInfo
	SetQueryContext(v *MultimodalQueryContext) *MedicalKnowResult
	GetQueryContext() *MultimodalQueryContext
	SetRequestId(v string) *MedicalKnowResult
	GetRequestId() *string
	SetSearchInformation(v *UnifiedSearchInformation) *MedicalKnowResult
	GetSearchInformation() *UnifiedSearchInformation
}

type MedicalKnowResult struct {
	Messages          []*MedicalKnowledgeInfo   `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	QueryContext      *MultimodalQueryContext   `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	RequestId         *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SearchInformation *UnifiedSearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
}

func (s MedicalKnowResult) String() string {
	return dara.Prettify(s)
}

func (s MedicalKnowResult) GoString() string {
	return s.String()
}

func (s *MedicalKnowResult) GetMessages() []*MedicalKnowledgeInfo {
	return s.Messages
}

func (s *MedicalKnowResult) GetQueryContext() *MultimodalQueryContext {
	return s.QueryContext
}

func (s *MedicalKnowResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalKnowResult) GetSearchInformation() *UnifiedSearchInformation {
	return s.SearchInformation
}

func (s *MedicalKnowResult) SetMessages(v []*MedicalKnowledgeInfo) *MedicalKnowResult {
	s.Messages = v
	return s
}

func (s *MedicalKnowResult) SetQueryContext(v *MultimodalQueryContext) *MedicalKnowResult {
	s.QueryContext = v
	return s
}

func (s *MedicalKnowResult) SetRequestId(v string) *MedicalKnowResult {
	s.RequestId = &v
	return s
}

func (s *MedicalKnowResult) SetSearchInformation(v *UnifiedSearchInformation) *MedicalKnowResult {
	s.SearchInformation = v
	return s
}

func (s *MedicalKnowResult) Validate() error {
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
