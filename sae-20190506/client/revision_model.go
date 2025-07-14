// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevision interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Revision
	GetCreatedTime() *string
	SetDescription(v string) *Revision
	GetDescription() *string
	SetRevisionConfig(v *RevisionConfig) *Revision
	GetRevisionConfig() *RevisionConfig
	SetRevisionId(v string) *Revision
	GetRevisionId() *string
	SetWeight(v float32) *Revision
	GetWeight() *float32
}

type Revision struct {
	CreatedTime    *string         `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description    *string         `json:"Description,omitempty" xml:"Description,omitempty"`
	RevisionConfig *RevisionConfig `json:"RevisionConfig,omitempty" xml:"RevisionConfig,omitempty"`
	RevisionId     *string         `json:"RevisionId,omitempty" xml:"RevisionId,omitempty"`
	Weight         *float32        `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s Revision) String() string {
	return dara.Prettify(s)
}

func (s Revision) GoString() string {
	return s.String()
}

func (s *Revision) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Revision) GetDescription() *string {
	return s.Description
}

func (s *Revision) GetRevisionConfig() *RevisionConfig {
	return s.RevisionConfig
}

func (s *Revision) GetRevisionId() *string {
	return s.RevisionId
}

func (s *Revision) GetWeight() *float32 {
	return s.Weight
}

func (s *Revision) SetCreatedTime(v string) *Revision {
	s.CreatedTime = &v
	return s
}

func (s *Revision) SetDescription(v string) *Revision {
	s.Description = &v
	return s
}

func (s *Revision) SetRevisionConfig(v *RevisionConfig) *Revision {
	s.RevisionConfig = v
	return s
}

func (s *Revision) SetRevisionId(v string) *Revision {
	s.RevisionId = &v
	return s
}

func (s *Revision) SetWeight(v float32) *Revision {
	s.Weight = &v
	return s
}

func (s *Revision) Validate() error {
	return dara.Validate(s)
}
