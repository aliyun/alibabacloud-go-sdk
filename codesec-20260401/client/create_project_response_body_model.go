// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRevision(v int64) *CreateProjectResponseBody
	GetConfigRevision() *int64
	SetCreatedAt(v string) *CreateProjectResponseBody
	GetCreatedAt() *string
	SetCreatedBy(v string) *CreateProjectResponseBody
	GetCreatedBy() *string
	SetDescription(v string) *CreateProjectResponseBody
	GetDescription() *string
	SetEngines(v *CreateProjectResponseBodyEngines) *CreateProjectResponseBody
	GetEngines() *CreateProjectResponseBodyEngines
	SetId(v int64) *CreateProjectResponseBody
	GetId() *int64
	SetInstructionPrompt(v string) *CreateProjectResponseBody
	GetInstructionPrompt() *string
	SetName(v string) *CreateProjectResponseBody
	GetName() *string
	SetRequestId(v string) *CreateProjectResponseBody
	GetRequestId() *string
	SetSource(v *CreateProjectResponseBodySource) *CreateProjectResponseBody
	GetSource() *CreateProjectResponseBodySource
	SetUpdatedAt(v string) *CreateProjectResponseBody
	GetUpdatedAt() *string
}

type CreateProjectResponseBody struct {
	// The project configuration version number.
	//
	// example:
	//
	// 1
	ConfigRevision *int64 `json:"configRevision,omitempty" xml:"configRevision,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The user ID of the project creator.
	//
	// example:
	//
	// 3221
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The description.
	//
	// example:
	//
	// This is default function description by fc-deploy component
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The engine switches for the project or scan snapshot. Only SAST and SCA are supported.
	Engines *CreateProjectResponseBodyEngines `json:"engines,omitempty" xml:"engines,omitempty" type:"Struct"`
	// The project ID.
	//
	// example:
	//
	// 111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The natural language prompt that describes scanning or result processing preferences, such as ignoring low-risk vulnerabilities.
	//
	// example:
	//
	// such as ignoring low-severity vulnerabilities, etc.
	InstructionPrompt *string `json:"instructionPrompt,omitempty" xml:"instructionPrompt,omitempty"`
	// The project name.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The project source.
	Source *CreateProjectResponseBodySource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The update time.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) GetConfigRevision() *int64 {
	return s.ConfigRevision
}

func (s *CreateProjectResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateProjectResponseBody) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreateProjectResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectResponseBody) GetEngines() *CreateProjectResponseBodyEngines {
	return s.Engines
}

func (s *CreateProjectResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateProjectResponseBody) GetInstructionPrompt() *string {
	return s.InstructionPrompt
}

func (s *CreateProjectResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectResponseBody) GetSource() *CreateProjectResponseBodySource {
	return s.Source
}

func (s *CreateProjectResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateProjectResponseBody) SetConfigRevision(v int64) *CreateProjectResponseBody {
	s.ConfigRevision = &v
	return s
}

func (s *CreateProjectResponseBody) SetCreatedAt(v string) *CreateProjectResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateProjectResponseBody) SetCreatedBy(v string) *CreateProjectResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *CreateProjectResponseBody) SetDescription(v string) *CreateProjectResponseBody {
	s.Description = &v
	return s
}

func (s *CreateProjectResponseBody) SetEngines(v *CreateProjectResponseBodyEngines) *CreateProjectResponseBody {
	s.Engines = v
	return s
}

func (s *CreateProjectResponseBody) SetId(v int64) *CreateProjectResponseBody {
	s.Id = &v
	return s
}

func (s *CreateProjectResponseBody) SetInstructionPrompt(v string) *CreateProjectResponseBody {
	s.InstructionPrompt = &v
	return s
}

func (s *CreateProjectResponseBody) SetName(v string) *CreateProjectResponseBody {
	s.Name = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSource(v *CreateProjectResponseBodySource) *CreateProjectResponseBody {
	s.Source = v
	return s
}

func (s *CreateProjectResponseBody) SetUpdatedAt(v string) *CreateProjectResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CreateProjectResponseBody) Validate() error {
	if s.Engines != nil {
		if err := s.Engines.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectResponseBodyEngines struct {
	// Specifies whether SAST is supported.
	//
	// example:
	//
	// true
	Sast *bool `json:"sast,omitempty" xml:"sast,omitempty"`
	// Specifies whether SCA is supported.
	//
	// example:
	//
	// true
	Sca *bool `json:"sca,omitempty" xml:"sca,omitempty"`
}

func (s CreateProjectResponseBodyEngines) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyEngines) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyEngines) GetSast() *bool {
	return s.Sast
}

func (s *CreateProjectResponseBodyEngines) GetSca() *bool {
	return s.Sca
}

func (s *CreateProjectResponseBodyEngines) SetSast(v bool) *CreateProjectResponseBodyEngines {
	s.Sast = &v
	return s
}

func (s *CreateProjectResponseBodyEngines) SetSca(v bool) *CreateProjectResponseBodyEngines {
	s.Sca = &v
	return s
}

func (s *CreateProjectResponseBodyEngines) Validate() error {
	return dara.Validate(s)
}

type CreateProjectResponseBodySource struct {
	// The project type.
	//
	// example:
	//
	// api
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateProjectResponseBodySource) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodySource) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodySource) GetType() *string {
	return s.Type
}

func (s *CreateProjectResponseBodySource) SetType(v string) *CreateProjectResponseBodySource {
	s.Type = &v
	return s
}

func (s *CreateProjectResponseBodySource) Validate() error {
	return dara.Validate(s)
}
