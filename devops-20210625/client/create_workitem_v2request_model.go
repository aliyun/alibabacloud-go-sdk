// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAssignedTo(v string) *CreateWorkitemV2Request
	GetAssignedTo() *string
	SetCategory(v string) *CreateWorkitemV2Request
	GetCategory() *string
	SetDescription(v string) *CreateWorkitemV2Request
	GetDescription() *string
	SetFieldValueList(v []*CreateWorkitemV2RequestFieldValueList) *CreateWorkitemV2Request
	GetFieldValueList() []*CreateWorkitemV2RequestFieldValueList
	SetParentIdentifier(v string) *CreateWorkitemV2Request
	GetParentIdentifier() *string
	SetParticipants(v []*string) *CreateWorkitemV2Request
	GetParticipants() []*string
	SetSpaceIdentifier(v string) *CreateWorkitemV2Request
	GetSpaceIdentifier() *string
	SetSprintIdentifier(v string) *CreateWorkitemV2Request
	GetSprintIdentifier() *string
	SetSubject(v string) *CreateWorkitemV2Request
	GetSubject() *string
	SetTags(v []*string) *CreateWorkitemV2Request
	GetTags() []*string
	SetTrackers(v []*string) *CreateWorkitemV2Request
	GetTrackers() []*string
	SetVerifier(v string) *CreateWorkitemV2Request
	GetVerifier() *string
	SetVersions(v []*string) *CreateWorkitemV2Request
	GetVersions() []*string
	SetWorkitemTypeIdentifier(v string) *CreateWorkitemV2Request
	GetWorkitemTypeIdentifier() *string
}

type CreateWorkitemV2Request struct {
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
	Category       *string                                  `json:"category,omitempty" xml:"category,omitempty"`
	Description    *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	FieldValueList []*CreateWorkitemV2RequestFieldValueList `json:"fieldValueList,omitempty" xml:"fieldValueList,omitempty" type:"Repeated"`
	// example:
	//
	// 11223331122
	ParentIdentifier *string   `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	Participants     []*string `json:"participants,omitempty" xml:"participants,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// asd345xxxxx9q9845xxxxx34
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// example:
	//
	// 455532323455
	SprintIdentifier *string `json:"sprintIdentifier,omitempty" xml:"sprintIdentifier,omitempty"`
	// This parameter is required.
	Subject  *string   `json:"subject,omitempty" xml:"subject,omitempty"`
	Tags     []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Trackers []*string `json:"trackers,omitempty" xml:"trackers,omitempty" type:"Repeated"`
	// example:
	//
	// 1561159309......
	Verifier *string   `json:"verifier,omitempty" xml:"verifier,omitempty"`
	Versions []*string `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 9uy29901re573f561d69jn40
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s CreateWorkitemV2Request) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemV2Request) GoString() string {
	return s.String()
}

func (s *CreateWorkitemV2Request) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *CreateWorkitemV2Request) GetCategory() *string {
	return s.Category
}

func (s *CreateWorkitemV2Request) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkitemV2Request) GetFieldValueList() []*CreateWorkitemV2RequestFieldValueList {
	return s.FieldValueList
}

func (s *CreateWorkitemV2Request) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateWorkitemV2Request) GetParticipants() []*string {
	return s.Participants
}

func (s *CreateWorkitemV2Request) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *CreateWorkitemV2Request) GetSprintIdentifier() *string {
	return s.SprintIdentifier
}

func (s *CreateWorkitemV2Request) GetSubject() *string {
	return s.Subject
}

func (s *CreateWorkitemV2Request) GetTags() []*string {
	return s.Tags
}

func (s *CreateWorkitemV2Request) GetTrackers() []*string {
	return s.Trackers
}

func (s *CreateWorkitemV2Request) GetVerifier() *string {
	return s.Verifier
}

func (s *CreateWorkitemV2Request) GetVersions() []*string {
	return s.Versions
}

func (s *CreateWorkitemV2Request) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *CreateWorkitemV2Request) SetAssignedTo(v string) *CreateWorkitemV2Request {
	s.AssignedTo = &v
	return s
}

func (s *CreateWorkitemV2Request) SetCategory(v string) *CreateWorkitemV2Request {
	s.Category = &v
	return s
}

func (s *CreateWorkitemV2Request) SetDescription(v string) *CreateWorkitemV2Request {
	s.Description = &v
	return s
}

func (s *CreateWorkitemV2Request) SetFieldValueList(v []*CreateWorkitemV2RequestFieldValueList) *CreateWorkitemV2Request {
	s.FieldValueList = v
	return s
}

func (s *CreateWorkitemV2Request) SetParentIdentifier(v string) *CreateWorkitemV2Request {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateWorkitemV2Request) SetParticipants(v []*string) *CreateWorkitemV2Request {
	s.Participants = v
	return s
}

func (s *CreateWorkitemV2Request) SetSpaceIdentifier(v string) *CreateWorkitemV2Request {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateWorkitemV2Request) SetSprintIdentifier(v string) *CreateWorkitemV2Request {
	s.SprintIdentifier = &v
	return s
}

func (s *CreateWorkitemV2Request) SetSubject(v string) *CreateWorkitemV2Request {
	s.Subject = &v
	return s
}

func (s *CreateWorkitemV2Request) SetTags(v []*string) *CreateWorkitemV2Request {
	s.Tags = v
	return s
}

func (s *CreateWorkitemV2Request) SetTrackers(v []*string) *CreateWorkitemV2Request {
	s.Trackers = v
	return s
}

func (s *CreateWorkitemV2Request) SetVerifier(v string) *CreateWorkitemV2Request {
	s.Verifier = &v
	return s
}

func (s *CreateWorkitemV2Request) SetVersions(v []*string) *CreateWorkitemV2Request {
	s.Versions = v
	return s
}

func (s *CreateWorkitemV2Request) SetWorkitemTypeIdentifier(v string) *CreateWorkitemV2Request {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *CreateWorkitemV2Request) Validate() error {
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

type CreateWorkitemV2RequestFieldValueList struct {
	// example:
	//
	// 6aexxxxxa1d98c09c60xxxx16
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// example:
	//
	// 10
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateWorkitemV2RequestFieldValueList) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemV2RequestFieldValueList) GoString() string {
	return s.String()
}

func (s *CreateWorkitemV2RequestFieldValueList) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *CreateWorkitemV2RequestFieldValueList) GetValue() *string {
	return s.Value
}

func (s *CreateWorkitemV2RequestFieldValueList) SetFieldIdentifier(v string) *CreateWorkitemV2RequestFieldValueList {
	s.FieldIdentifier = &v
	return s
}

func (s *CreateWorkitemV2RequestFieldValueList) SetValue(v string) *CreateWorkitemV2RequestFieldValueList {
	s.Value = &v
	return s
}

func (s *CreateWorkitemV2RequestFieldValueList) Validate() error {
	return dara.Validate(s)
}
