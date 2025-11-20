// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssistantCapabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantDescription(v string) *GetAssistantCapabilityResponseBody
	GetAssistantDescription() *string
	SetCanHandle(v bool) *GetAssistantCapabilityResponseBody
	GetCanHandle() *bool
	SetCapabilityAssessment(v *GetAssistantCapabilityResponseBodyCapabilityAssessment) *GetAssistantCapabilityResponseBody
	GetCapabilityAssessment() *GetAssistantCapabilityResponseBodyCapabilityAssessment
	SetRequestId(v string) *GetAssistantCapabilityResponseBody
	GetRequestId() *string
}

type GetAssistantCapabilityResponseBody struct {
	// example:
	//
	// 助理描述
	AssistantDescription *string `json:"assistantDescription,omitempty" xml:"assistantDescription,omitempty"`
	// example:
	//
	// true
	CanHandle *bool `json:"canHandle,omitempty" xml:"canHandle,omitempty"`
	// example:
	//
	// {}
	CapabilityAssessment *GetAssistantCapabilityResponseBodyCapabilityAssessment `json:"capabilityAssessment,omitempty" xml:"capabilityAssessment,omitempty" type:"Struct"`
	// example:
	//
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAssistantCapabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityResponseBody) GetAssistantDescription() *string {
	return s.AssistantDescription
}

func (s *GetAssistantCapabilityResponseBody) GetCanHandle() *bool {
	return s.CanHandle
}

func (s *GetAssistantCapabilityResponseBody) GetCapabilityAssessment() *GetAssistantCapabilityResponseBodyCapabilityAssessment {
	return s.CapabilityAssessment
}

func (s *GetAssistantCapabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssistantCapabilityResponseBody) SetAssistantDescription(v string) *GetAssistantCapabilityResponseBody {
	s.AssistantDescription = &v
	return s
}

func (s *GetAssistantCapabilityResponseBody) SetCanHandle(v bool) *GetAssistantCapabilityResponseBody {
	s.CanHandle = &v
	return s
}

func (s *GetAssistantCapabilityResponseBody) SetCapabilityAssessment(v *GetAssistantCapabilityResponseBodyCapabilityAssessment) *GetAssistantCapabilityResponseBody {
	s.CapabilityAssessment = v
	return s
}

func (s *GetAssistantCapabilityResponseBody) SetRequestId(v string) *GetAssistantCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssistantCapabilityResponseBody) Validate() error {
	if s.CapabilityAssessment != nil {
		if err := s.CapabilityAssessment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAssistantCapabilityResponseBodyCapabilityAssessment struct {
	// example:
	//
	// 能力概览
	BriefCapability *string `json:"briefCapability,omitempty" xml:"briefCapability,omitempty"`
	// example:
	//
	// []
	CapabilityList []*GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList `json:"capabilityList,omitempty" xml:"capabilityList,omitempty" type:"Repeated"`
}

func (s GetAssistantCapabilityResponseBodyCapabilityAssessment) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityResponseBodyCapabilityAssessment) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessment) GetBriefCapability() *string {
	return s.BriefCapability
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessment) GetCapabilityList() []*GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList {
	return s.CapabilityList
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessment) SetBriefCapability(v string) *GetAssistantCapabilityResponseBodyCapabilityAssessment {
	s.BriefCapability = &v
	return s
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessment) SetCapabilityList(v []*GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) *GetAssistantCapabilityResponseBodyCapabilityAssessment {
	s.CapabilityList = v
	return s
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessment) Validate() error {
	if s.CapabilityList != nil {
		for _, item := range s.CapabilityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList struct {
	// example:
	//
	// 能力概览
	CapabilityOverview *string `json:"capabilityOverview,omitempty" xml:"capabilityOverview,omitempty"`
	// example:
	//
	// 能力描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 能力名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) GetCapabilityOverview() *string {
	return s.CapabilityOverview
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) GetDescription() *string {
	return s.Description
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) GetName() *string {
	return s.Name
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) SetCapabilityOverview(v string) *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList {
	s.CapabilityOverview = &v
	return s
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) SetDescription(v string) *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList {
	s.Description = &v
	return s
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) SetName(v string) *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList {
	s.Name = &v
	return s
}

func (s *GetAssistantCapabilityResponseBodyCapabilityAssessmentCapabilityList) Validate() error {
	return dara.Validate(s)
}
