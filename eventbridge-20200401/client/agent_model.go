// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgent interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Agent
	GetCreatedAt() *int64
	SetDescription(v string) *Agent
	GetDescription() *string
	SetMetadata(v *Metadata) *Agent
	GetMetadata() *Metadata
	SetName(v string) *Agent
	GetName() *string
	SetPrompt(v string) *Agent
	GetPrompt() *string
	SetUpdatedAt(v int64) *Agent
	GetUpdatedAt() *int64
}

type Agent struct {
	CreatedAt   *int64    `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Metadata    *Metadata `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Prompt      *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	UpdatedAt   *int64    `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s Agent) String() string {
	return dara.Prettify(s)
}

func (s Agent) GoString() string {
	return s.String()
}

func (s *Agent) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Agent) GetDescription() *string {
	return s.Description
}

func (s *Agent) GetMetadata() *Metadata {
	return s.Metadata
}

func (s *Agent) GetName() *string {
	return s.Name
}

func (s *Agent) GetPrompt() *string {
	return s.Prompt
}

func (s *Agent) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Agent) SetCreatedAt(v int64) *Agent {
	s.CreatedAt = &v
	return s
}

func (s *Agent) SetDescription(v string) *Agent {
	s.Description = &v
	return s
}

func (s *Agent) SetMetadata(v *Metadata) *Agent {
	s.Metadata = v
	return s
}

func (s *Agent) SetName(v string) *Agent {
	s.Name = &v
	return s
}

func (s *Agent) SetPrompt(v string) *Agent {
	s.Prompt = &v
	return s
}

func (s *Agent) SetUpdatedAt(v int64) *Agent {
	s.UpdatedAt = &v
	return s
}

func (s *Agent) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
