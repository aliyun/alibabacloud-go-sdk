// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateDataAgentMemoryRequest
	GetContent() *string
	SetDMSUnit(v string) *UpdateDataAgentMemoryRequest
	GetDMSUnit() *string
	SetFromId(v string) *UpdateDataAgentMemoryRequest
	GetFromId() *string
	SetMemFrom(v string) *UpdateDataAgentMemoryRequest
	GetMemFrom() *string
	SetUuid(v string) *UpdateDataAgentMemoryRequest
	GetUuid() *string
}

type UpdateDataAgentMemoryRequest struct {
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
	// 8zm3**********g3yxa1
	FromId *string `json:"FromId,omitempty" xml:"FromId,omitempty"`
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
	// The memory UUID.
	//
	// example:
	//
	// ed3f67***********ed
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UpdateDataAgentMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentMemoryRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateDataAgentMemoryRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *UpdateDataAgentMemoryRequest) GetFromId() *string {
	return s.FromId
}

func (s *UpdateDataAgentMemoryRequest) GetMemFrom() *string {
	return s.MemFrom
}

func (s *UpdateDataAgentMemoryRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateDataAgentMemoryRequest) SetContent(v string) *UpdateDataAgentMemoryRequest {
	s.Content = &v
	return s
}

func (s *UpdateDataAgentMemoryRequest) SetDMSUnit(v string) *UpdateDataAgentMemoryRequest {
	s.DMSUnit = &v
	return s
}

func (s *UpdateDataAgentMemoryRequest) SetFromId(v string) *UpdateDataAgentMemoryRequest {
	s.FromId = &v
	return s
}

func (s *UpdateDataAgentMemoryRequest) SetMemFrom(v string) *UpdateDataAgentMemoryRequest {
	s.MemFrom = &v
	return s
}

func (s *UpdateDataAgentMemoryRequest) SetUuid(v string) *UpdateDataAgentMemoryRequest {
	s.Uuid = &v
	return s
}

func (s *UpdateDataAgentMemoryRequest) Validate() error {
	return dara.Validate(s)
}
