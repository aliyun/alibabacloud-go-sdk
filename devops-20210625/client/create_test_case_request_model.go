// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTestCaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignedTo(v string) *CreateTestCaseRequest
	GetAssignedTo() *string
	SetDirectoryIdentifier(v string) *CreateTestCaseRequest
	GetDirectoryIdentifier() *string
	SetFieldValueList(v []*CreateTestCaseRequestFieldValueList) *CreateTestCaseRequest
	GetFieldValueList() []*CreateTestCaseRequestFieldValueList
	SetPriority(v string) *CreateTestCaseRequest
	GetPriority() *string
	SetSpaceIdentifier(v string) *CreateTestCaseRequest
	GetSpaceIdentifier() *string
	SetSubject(v string) *CreateTestCaseRequest
	GetSubject() *string
	SetTags(v []*string) *CreateTestCaseRequest
	GetTags() []*string
	SetTestcaseStepContentInfo(v *CreateTestCaseRequestTestcaseStepContentInfo) *CreateTestCaseRequest
	GetTestcaseStepContentInfo() *CreateTestCaseRequestTestcaseStepContentInfo
}

type CreateTestCaseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 19xxxx31947xxxx
	AssignedTo *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fdd395xxxxx9q9845xxxxx23
	DirectoryIdentifier *string                                `json:"directoryIdentifier,omitempty" xml:"directoryIdentifier,omitempty"`
	FieldValueList      []*CreateTestCaseRequestFieldValueList `json:"fieldValueList,omitempty" xml:"fieldValueList,omitempty" type:"Repeated"`
	// example:
	//
	// ik3dexxxxxfdfds1xxxxx23
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asd345xxxxx9q9845xxxxx34
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// This parameter is required.
	Subject                 *string                                       `json:"subject,omitempty" xml:"subject,omitempty"`
	Tags                    []*string                                     `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TestcaseStepContentInfo *CreateTestCaseRequestTestcaseStepContentInfo `json:"testcaseStepContentInfo,omitempty" xml:"testcaseStepContentInfo,omitempty" type:"Struct"`
}

func (s CreateTestCaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseRequest) GoString() string {
	return s.String()
}

func (s *CreateTestCaseRequest) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *CreateTestCaseRequest) GetDirectoryIdentifier() *string {
	return s.DirectoryIdentifier
}

func (s *CreateTestCaseRequest) GetFieldValueList() []*CreateTestCaseRequestFieldValueList {
	return s.FieldValueList
}

func (s *CreateTestCaseRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateTestCaseRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *CreateTestCaseRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreateTestCaseRequest) GetTags() []*string {
	return s.Tags
}

func (s *CreateTestCaseRequest) GetTestcaseStepContentInfo() *CreateTestCaseRequestTestcaseStepContentInfo {
	return s.TestcaseStepContentInfo
}

func (s *CreateTestCaseRequest) SetAssignedTo(v string) *CreateTestCaseRequest {
	s.AssignedTo = &v
	return s
}

func (s *CreateTestCaseRequest) SetDirectoryIdentifier(v string) *CreateTestCaseRequest {
	s.DirectoryIdentifier = &v
	return s
}

func (s *CreateTestCaseRequest) SetFieldValueList(v []*CreateTestCaseRequestFieldValueList) *CreateTestCaseRequest {
	s.FieldValueList = v
	return s
}

func (s *CreateTestCaseRequest) SetPriority(v string) *CreateTestCaseRequest {
	s.Priority = &v
	return s
}

func (s *CreateTestCaseRequest) SetSpaceIdentifier(v string) *CreateTestCaseRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateTestCaseRequest) SetSubject(v string) *CreateTestCaseRequest {
	s.Subject = &v
	return s
}

func (s *CreateTestCaseRequest) SetTags(v []*string) *CreateTestCaseRequest {
	s.Tags = v
	return s
}

func (s *CreateTestCaseRequest) SetTestcaseStepContentInfo(v *CreateTestCaseRequestTestcaseStepContentInfo) *CreateTestCaseRequest {
	s.TestcaseStepContentInfo = v
	return s
}

func (s *CreateTestCaseRequest) Validate() error {
	if s.FieldValueList != nil {
		for _, item := range s.FieldValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TestcaseStepContentInfo != nil {
		if err := s.TestcaseStepContentInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTestCaseRequestFieldValueList struct {
	// example:
	//
	// 6aexxxxxa1d98c09c60xxxx16
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// example:
	//
	// 77c7fb03c4186c8691d6...
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateTestCaseRequestFieldValueList) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseRequestFieldValueList) GoString() string {
	return s.String()
}

func (s *CreateTestCaseRequestFieldValueList) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *CreateTestCaseRequestFieldValueList) GetValue() *string {
	return s.Value
}

func (s *CreateTestCaseRequestFieldValueList) SetFieldIdentifier(v string) *CreateTestCaseRequestFieldValueList {
	s.FieldIdentifier = &v
	return s
}

func (s *CreateTestCaseRequestFieldValueList) SetValue(v string) *CreateTestCaseRequestFieldValueList {
	s.Value = &v
	return s
}

func (s *CreateTestCaseRequestFieldValueList) Validate() error {
	return dara.Validate(s)
}

type CreateTestCaseRequestTestcaseStepContentInfo struct {
	Precondition   *string                                                       `json:"precondition,omitempty" xml:"precondition,omitempty"`
	StepResultList []*CreateTestCaseRequestTestcaseStepContentInfoStepResultList `json:"stepResultList,omitempty" xml:"stepResultList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TEXT
	StepType *string `json:"stepType,omitempty" xml:"stepType,omitempty"`
}

func (s CreateTestCaseRequestTestcaseStepContentInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseRequestTestcaseStepContentInfo) GoString() string {
	return s.String()
}

func (s *CreateTestCaseRequestTestcaseStepContentInfo) GetPrecondition() *string {
	return s.Precondition
}

func (s *CreateTestCaseRequestTestcaseStepContentInfo) GetStepResultList() []*CreateTestCaseRequestTestcaseStepContentInfoStepResultList {
	return s.StepResultList
}

func (s *CreateTestCaseRequestTestcaseStepContentInfo) GetStepType() *string {
	return s.StepType
}

func (s *CreateTestCaseRequestTestcaseStepContentInfo) SetPrecondition(v string) *CreateTestCaseRequestTestcaseStepContentInfo {
	s.Precondition = &v
	return s
}

func (s *CreateTestCaseRequestTestcaseStepContentInfo) SetStepResultList(v []*CreateTestCaseRequestTestcaseStepContentInfoStepResultList) *CreateTestCaseRequestTestcaseStepContentInfo {
	s.StepResultList = v
	return s
}

func (s *CreateTestCaseRequestTestcaseStepContentInfo) SetStepType(v string) *CreateTestCaseRequestTestcaseStepContentInfo {
	s.StepType = &v
	return s
}

func (s *CreateTestCaseRequestTestcaseStepContentInfo) Validate() error {
	if s.StepResultList != nil {
		for _, item := range s.StepResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTestCaseRequestTestcaseStepContentInfoStepResultList struct {
	Expected *string `json:"expected,omitempty" xml:"expected,omitempty"`
	Step     *string `json:"step,omitempty" xml:"step,omitempty"`
}

func (s CreateTestCaseRequestTestcaseStepContentInfoStepResultList) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseRequestTestcaseStepContentInfoStepResultList) GoString() string {
	return s.String()
}

func (s *CreateTestCaseRequestTestcaseStepContentInfoStepResultList) GetExpected() *string {
	return s.Expected
}

func (s *CreateTestCaseRequestTestcaseStepContentInfoStepResultList) GetStep() *string {
	return s.Step
}

func (s *CreateTestCaseRequestTestcaseStepContentInfoStepResultList) SetExpected(v string) *CreateTestCaseRequestTestcaseStepContentInfoStepResultList {
	s.Expected = &v
	return s
}

func (s *CreateTestCaseRequestTestcaseStepContentInfoStepResultList) SetStep(v string) *CreateTestCaseRequestTestcaseStepContentInfoStepResultList {
	s.Step = &v
	return s
}

func (s *CreateTestCaseRequestTestcaseStepContentInfoStepResultList) Validate() error {
	return dara.Validate(s)
}
