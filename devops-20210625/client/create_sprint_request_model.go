// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSprintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *CreateSprintRequest
	GetEndDate() *string
	SetName(v string) *CreateSprintRequest
	GetName() *string
	SetSpaceIdentifier(v string) *CreateSprintRequest
	GetSpaceIdentifier() *string
	SetStaffIds(v []*string) *CreateSprintRequest
	GetStaffIds() []*string
	SetStartDate(v string) *CreateSprintRequest
	GetStartDate() *string
}

type CreateSprintRequest struct {
	// example:
	//
	// 2021-12-02
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asd345xxxxx9q9845xxxxx34
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// This parameter is required.
	StaffIds []*string `json:"staffIds,omitempty" xml:"staffIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s CreateSprintRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSprintRequest) GoString() string {
	return s.String()
}

func (s *CreateSprintRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateSprintRequest) GetName() *string {
	return s.Name
}

func (s *CreateSprintRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *CreateSprintRequest) GetStaffIds() []*string {
	return s.StaffIds
}

func (s *CreateSprintRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateSprintRequest) SetEndDate(v string) *CreateSprintRequest {
	s.EndDate = &v
	return s
}

func (s *CreateSprintRequest) SetName(v string) *CreateSprintRequest {
	s.Name = &v
	return s
}

func (s *CreateSprintRequest) SetSpaceIdentifier(v string) *CreateSprintRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateSprintRequest) SetStaffIds(v []*string) *CreateSprintRequest {
	s.StaffIds = v
	return s
}

func (s *CreateSprintRequest) SetStartDate(v string) *CreateSprintRequest {
	s.StartDate = &v
	return s
}

func (s *CreateSprintRequest) Validate() error {
	return dara.Validate(s)
}
