// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataAgentMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *AddDataAgentMemoryRequest
	GetContent() *string
	SetDMSUnit(v string) *AddDataAgentMemoryRequest
	GetDMSUnit() *string
	SetFromId(v string) *AddDataAgentMemoryRequest
	GetFromId() *string
	SetLabel(v string) *AddDataAgentMemoryRequest
	GetLabel() *string
	SetMemFrom(v string) *AddDataAgentMemoryRequest
	GetMemFrom() *string
	SetSessionUuid(v string) *AddDataAgentMemoryRequest
	GetSessionUuid() *string
}

type AddDataAgentMemoryRequest struct {
	// The memory content.
	//
	// example:
	//
	// Diamond pricing analysis requires examining the skewness and outliers of the distribution of each feature.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The source ID.
	//
	// - If MemFrom is set to session, FromId indicates the session ID.
	//
	// - If MemFrom is set to user, FromId indicates the RAM user ID.
	//
	// example:
	//
	// 2037**********23
	FromId *string `json:"FromId,omitempty" xml:"FromId,omitempty"`
	// The memory label. Valid values:
	//
	// - fact_specifications: fact definitions.
	//
	// - task_constraints: node constraints.
	//
	// - execution_config: execution configuration.
	//
	// example:
	//
	// fact_specifications
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The memory source. Valid values:
	//
	// - session: generated from a session.
	//
	// - user: edited by a user.
	//
	// example:
	//
	// user
	MemFrom *string `json:"MemFrom,omitempty" xml:"MemFrom,omitempty"`
	// The session ID.
	//
	// - Note: This parameter is deprecated.
	//
	// example:
	//
	// fc5ice**********ac6e
	SessionUuid *string `json:"SessionUuid,omitempty" xml:"SessionUuid,omitempty"`
}

func (s AddDataAgentMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataAgentMemoryRequest) GoString() string {
	return s.String()
}

func (s *AddDataAgentMemoryRequest) GetContent() *string {
	return s.Content
}

func (s *AddDataAgentMemoryRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *AddDataAgentMemoryRequest) GetFromId() *string {
	return s.FromId
}

func (s *AddDataAgentMemoryRequest) GetLabel() *string {
	return s.Label
}

func (s *AddDataAgentMemoryRequest) GetMemFrom() *string {
	return s.MemFrom
}

func (s *AddDataAgentMemoryRequest) GetSessionUuid() *string {
	return s.SessionUuid
}

func (s *AddDataAgentMemoryRequest) SetContent(v string) *AddDataAgentMemoryRequest {
	s.Content = &v
	return s
}

func (s *AddDataAgentMemoryRequest) SetDMSUnit(v string) *AddDataAgentMemoryRequest {
	s.DMSUnit = &v
	return s
}

func (s *AddDataAgentMemoryRequest) SetFromId(v string) *AddDataAgentMemoryRequest {
	s.FromId = &v
	return s
}

func (s *AddDataAgentMemoryRequest) SetLabel(v string) *AddDataAgentMemoryRequest {
	s.Label = &v
	return s
}

func (s *AddDataAgentMemoryRequest) SetMemFrom(v string) *AddDataAgentMemoryRequest {
	s.MemFrom = &v
	return s
}

func (s *AddDataAgentMemoryRequest) SetSessionUuid(v string) *AddDataAgentMemoryRequest {
	s.SessionUuid = &v
	return s
}

func (s *AddDataAgentMemoryRequest) Validate() error {
	return dara.Validate(s)
}
