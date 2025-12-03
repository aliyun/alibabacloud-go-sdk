// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignedTo(v string) *CreateWorkitemRequest
	GetAssignedTo() *string
	SetCategory(v string) *CreateWorkitemRequest
	GetCategory() *string
	SetDescription(v string) *CreateWorkitemRequest
	GetDescription() *string
	SetDescriptionFormat(v string) *CreateWorkitemRequest
	GetDescriptionFormat() *string
	SetFieldValueList(v []*CreateWorkitemRequestFieldValueList) *CreateWorkitemRequest
	GetFieldValueList() []*CreateWorkitemRequestFieldValueList
	SetParent(v string) *CreateWorkitemRequest
	GetParent() *string
	SetParticipant(v []*string) *CreateWorkitemRequest
	GetParticipant() []*string
	SetSpace(v string) *CreateWorkitemRequest
	GetSpace() *string
	SetSpaceIdentifier(v string) *CreateWorkitemRequest
	GetSpaceIdentifier() *string
	SetSpaceType(v string) *CreateWorkitemRequest
	GetSpaceType() *string
	SetSprint(v []*string) *CreateWorkitemRequest
	GetSprint() []*string
	SetSubject(v string) *CreateWorkitemRequest
	GetSubject() *string
	SetTracker(v []*string) *CreateWorkitemRequest
	GetTracker() []*string
	SetVerifier(v []*string) *CreateWorkitemRequest
	GetVerifier() []*string
	SetWorkitemType(v string) *CreateWorkitemRequest
	GetWorkitemType() *string
}

type CreateWorkitemRequest struct {
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
	// Req
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 测试内容
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// RICHTEXT
	DescriptionFormat *string                                `json:"descriptionFormat,omitempty" xml:"descriptionFormat,omitempty"`
	FieldValueList    []*CreateWorkitemRequestFieldValueList `json:"fieldValueList,omitempty" xml:"fieldValueList,omitempty" type:"Repeated"`
	// example:
	//
	// 3a0c9cdd24ae1e1995b8...
	Parent      *string   `json:"parent,omitempty" xml:"parent,omitempty"`
	Participant []*string `json:"participant,omitempty" xml:"participant,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// asd345xxxxx9q9845xxxxx34
	Space *string `json:"space,omitempty" xml:"space,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asd345xxxxx9q9845xxxxx34
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Project
	SpaceType *string   `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Sprint    []*string `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 测试工作项
	Subject  *string   `json:"subject,omitempty" xml:"subject,omitempty"`
	Tracker  []*string `json:"tracker,omitempty" xml:"tracker,omitempty" type:"Repeated"`
	Verifier []*string `json:"verifier,omitempty" xml:"verifier,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 9uyxxxxxre573f561dxxn40
	WorkitemType *string `json:"workitemType,omitempty" xml:"workitemType,omitempty"`
}

func (s CreateWorkitemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRequest) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *CreateWorkitemRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateWorkitemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkitemRequest) GetDescriptionFormat() *string {
	return s.DescriptionFormat
}

func (s *CreateWorkitemRequest) GetFieldValueList() []*CreateWorkitemRequestFieldValueList {
	return s.FieldValueList
}

func (s *CreateWorkitemRequest) GetParent() *string {
	return s.Parent
}

func (s *CreateWorkitemRequest) GetParticipant() []*string {
	return s.Participant
}

func (s *CreateWorkitemRequest) GetSpace() *string {
	return s.Space
}

func (s *CreateWorkitemRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *CreateWorkitemRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *CreateWorkitemRequest) GetSprint() []*string {
	return s.Sprint
}

func (s *CreateWorkitemRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreateWorkitemRequest) GetTracker() []*string {
	return s.Tracker
}

func (s *CreateWorkitemRequest) GetVerifier() []*string {
	return s.Verifier
}

func (s *CreateWorkitemRequest) GetWorkitemType() *string {
	return s.WorkitemType
}

func (s *CreateWorkitemRequest) SetAssignedTo(v string) *CreateWorkitemRequest {
	s.AssignedTo = &v
	return s
}

func (s *CreateWorkitemRequest) SetCategory(v string) *CreateWorkitemRequest {
	s.Category = &v
	return s
}

func (s *CreateWorkitemRequest) SetDescription(v string) *CreateWorkitemRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkitemRequest) SetDescriptionFormat(v string) *CreateWorkitemRequest {
	s.DescriptionFormat = &v
	return s
}

func (s *CreateWorkitemRequest) SetFieldValueList(v []*CreateWorkitemRequestFieldValueList) *CreateWorkitemRequest {
	s.FieldValueList = v
	return s
}

func (s *CreateWorkitemRequest) SetParent(v string) *CreateWorkitemRequest {
	s.Parent = &v
	return s
}

func (s *CreateWorkitemRequest) SetParticipant(v []*string) *CreateWorkitemRequest {
	s.Participant = v
	return s
}

func (s *CreateWorkitemRequest) SetSpace(v string) *CreateWorkitemRequest {
	s.Space = &v
	return s
}

func (s *CreateWorkitemRequest) SetSpaceIdentifier(v string) *CreateWorkitemRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateWorkitemRequest) SetSpaceType(v string) *CreateWorkitemRequest {
	s.SpaceType = &v
	return s
}

func (s *CreateWorkitemRequest) SetSprint(v []*string) *CreateWorkitemRequest {
	s.Sprint = v
	return s
}

func (s *CreateWorkitemRequest) SetSubject(v string) *CreateWorkitemRequest {
	s.Subject = &v
	return s
}

func (s *CreateWorkitemRequest) SetTracker(v []*string) *CreateWorkitemRequest {
	s.Tracker = v
	return s
}

func (s *CreateWorkitemRequest) SetVerifier(v []*string) *CreateWorkitemRequest {
	s.Verifier = v
	return s
}

func (s *CreateWorkitemRequest) SetWorkitemType(v string) *CreateWorkitemRequest {
	s.WorkitemType = &v
	return s
}

func (s *CreateWorkitemRequest) Validate() error {
	if s.FieldValueList != nil {
		for _, item := range s.FieldValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkitemRequestFieldValueList struct {
	// example:
	//
	// 6aexxxxxa1d98c09c60xxxx16
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// example:
	//
	// 77c7fb03c4186c8691d6...
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// null
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemRequestFieldValueList) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemRequestFieldValueList) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRequestFieldValueList) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *CreateWorkitemRequestFieldValueList) GetValue() *string {
	return s.Value
}

func (s *CreateWorkitemRequestFieldValueList) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *CreateWorkitemRequestFieldValueList) SetFieldIdentifier(v string) *CreateWorkitemRequestFieldValueList {
	s.FieldIdentifier = &v
	return s
}

func (s *CreateWorkitemRequestFieldValueList) SetValue(v string) *CreateWorkitemRequestFieldValueList {
	s.Value = &v
	return s
}

func (s *CreateWorkitemRequestFieldValueList) SetWorkitemIdentifier(v string) *CreateWorkitemRequestFieldValueList {
	s.WorkitemIdentifier = &v
	return s
}

func (s *CreateWorkitemRequestFieldValueList) Validate() error {
	return dara.Validate(s)
}
