// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkMeetingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *CreatePersonalDingtalkMeetingRequest
	GetCredentialId() *string
	SetDescription(v string) *CreatePersonalDingtalkMeetingRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkMeetingRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalDingtalkMeetingRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalDingtalkMeetingRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalDingtalkMeetingRequest
	GetOperatingObjectName() *string
	SetRoomCode(v string) *CreatePersonalDingtalkMeetingRequest
	GetRoomCode() *string
	SetTenantId(v string) *CreatePersonalDingtalkMeetingRequest
	GetTenantId() *string
}

type CreatePersonalDingtalkMeetingRequest struct {
	// The credential ID.
	//
	// example:
	//
	// exampleCredentialId
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The pipeline description.
	//
	// example:
	//
	// Watchlist Monitor Layer
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-default-007735a2-58f5-47a5-9e37-ea3fd64e0899
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
	//
	// example:
	//
	// string_value
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The meeting code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 356 776 973
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalDingtalkMeetingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkMeetingRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreatePersonalDingtalkMeetingRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDingtalkMeetingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkMeetingRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDingtalkMeetingRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalDingtalkMeetingRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDingtalkMeetingRequest) GetRoomCode() *string {
	return s.RoomCode
}

func (s *CreatePersonalDingtalkMeetingRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDingtalkMeetingRequest) SetCredentialId(v string) *CreatePersonalDingtalkMeetingRequest {
	s.CredentialId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetDescription(v string) *CreatePersonalDingtalkMeetingRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetDirectoryId(v string) *CreatePersonalDingtalkMeetingRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetName(v string) *CreatePersonalDingtalkMeetingRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetNotes(v string) *CreatePersonalDingtalkMeetingRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetOperatingObjectName(v string) *CreatePersonalDingtalkMeetingRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetRoomCode(v string) *CreatePersonalDingtalkMeetingRequest {
	s.RoomCode = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetTenantId(v string) *CreatePersonalDingtalkMeetingRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) Validate() error {
	return dara.Validate(s)
}
