// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentResourceStruct interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *IncidentResourceStruct
	GetDescription() *string
	SetIncidentId(v string) *IncidentResourceStruct
	GetIncidentId() *string
	SetIncidentResourceId(v string) *IncidentResourceStruct
	GetIncidentResourceId() *string
	SetResource(v *IncidentResourceDetail) *IncidentResourceStruct
	GetResource() *IncidentResourceDetail
	SetSource(v string) *IncidentResourceStruct
	GetSource() *string
	SetTime(v int64) *IncidentResourceStruct
	GetTime() *int64
	SetUserId(v int64) *IncidentResourceStruct
	GetUserId() *int64
}

type IncidentResourceStruct struct {
	Description        *string                 `json:"description,omitempty" xml:"description,omitempty"`
	IncidentId         *string                 `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentResourceId *string                 `json:"incidentResourceId,omitempty" xml:"incidentResourceId,omitempty"`
	Resource           *IncidentResourceDetail `json:"resource,omitempty" xml:"resource,omitempty"`
	Source             *string                 `json:"source,omitempty" xml:"source,omitempty"`
	Time               *int64                  `json:"time,omitempty" xml:"time,omitempty"`
	UserId             *int64                  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentResourceStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentResourceStruct) GoString() string {
	return s.String()
}

func (s *IncidentResourceStruct) GetDescription() *string {
	return s.Description
}

func (s *IncidentResourceStruct) GetIncidentId() *string {
	return s.IncidentId
}

func (s *IncidentResourceStruct) GetIncidentResourceId() *string {
	return s.IncidentResourceId
}

func (s *IncidentResourceStruct) GetResource() *IncidentResourceDetail {
	return s.Resource
}

func (s *IncidentResourceStruct) GetSource() *string {
	return s.Source
}

func (s *IncidentResourceStruct) GetTime() *int64 {
	return s.Time
}

func (s *IncidentResourceStruct) GetUserId() *int64 {
	return s.UserId
}

func (s *IncidentResourceStruct) SetDescription(v string) *IncidentResourceStruct {
	s.Description = &v
	return s
}

func (s *IncidentResourceStruct) SetIncidentId(v string) *IncidentResourceStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentResourceStruct) SetIncidentResourceId(v string) *IncidentResourceStruct {
	s.IncidentResourceId = &v
	return s
}

func (s *IncidentResourceStruct) SetResource(v *IncidentResourceDetail) *IncidentResourceStruct {
	s.Resource = v
	return s
}

func (s *IncidentResourceStruct) SetSource(v string) *IncidentResourceStruct {
	s.Source = &v
	return s
}

func (s *IncidentResourceStruct) SetTime(v int64) *IncidentResourceStruct {
	s.Time = &v
	return s
}

func (s *IncidentResourceStruct) SetUserId(v int64) *IncidentResourceStruct {
	s.UserId = &v
	return s
}

func (s *IncidentResourceStruct) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}
