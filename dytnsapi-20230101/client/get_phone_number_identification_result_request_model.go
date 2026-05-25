// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberIdentificationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GetPhoneNumberIdentificationResultRequest
	GetAuthCode() *string
	SetOutId(v string) *GetPhoneNumberIdentificationResultRequest
	GetOutId() *string
	SetOwnerId(v int64) *GetPhoneNumberIdentificationResultRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetPhoneNumberIdentificationResultRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetPhoneNumberIdentificationResultRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPhoneNumberIdentificationResultRequest
	GetResourceOwnerId() *int64
	SetSessionId(v string) *GetPhoneNumberIdentificationResultRequest
	GetSessionId() *string
	SetSessionPayload(v string) *GetPhoneNumberIdentificationResultRequest
	GetSessionPayload() *string
}

type GetPhoneNumberIdentificationResultRequest struct {
	// The authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// K***9i7CIe
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The external ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 149b03d2-a749-4e6e-8f5b-34******5815
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number of the subscriber. The phone number to be verified must be in the Mobile Station International Subscriber Directory Number (MSISDN) format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 628211****113
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8636b75e2fcb40c53ffecc2b5947115c.149b03d2a7494e6e8f5b34c915245815.707c7f0d93f4409db0761aa5da94ce01.1686******041
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The session payload.
	//
	// This parameter is required.
	//
	// example:
	//
	// uQne0vsuNywXVvI4VP5taHsgDNsd3BwcbmrhjXi58WbxBGFW+e8ufMEi9j89YonphV6NZ1PIeKvboHtU1nsSjZMTcoFPfkjqaORIHdSlPb6vmIzqOnJMsP1KPQ8K1JLXSaAKsB2lQ5A9HCkX2HzDEwje14HYQsnPd/Ka2YWgXuL0N8GE9oYi25d4DdlU0XR52YjSj8GMLSgbW7yNxEPvUCOQG83FZfQqmIWG2+0C/fQ3gdG9WI7AeeHZo4IRKGtQnpjKGtZZl8VoLPNIswDqZeeyjCyZlKUXKrAt4Co9c4I4q8G1jZm53COQJ+DuTiWH7w+tois3WJwFV/HmdlAKt8SqpiVrEv47VQ9V+8FYsdKz3A3CRyBVgNj6wYKKbwaI9BdQoOkbYzzA8CfAKO5w1oYVD2nOcYS/AffbPbE31PJj7SdVvKghwPL56OVdjS9Hd0iW0SMBWD0F1iRNCUNHL3ffHcFjJLdhTrMt8VHSRn0nOlvO1ZaWqMQ0yE0q*************************kXTpoQLo0+0h+CEcf90hTg8XdMhj9B0A3SOINceLlmoZb3czvYl00+CC0075DjOX41YtnuAUfaNYPgLIZkjYyq+JopBQFAkxPUbJHC0oCzB9dQahUthWY38OPBs=
	SessionPayload *string `json:"SessionPayload,omitempty" xml:"SessionPayload,omitempty"`
}

func (s GetPhoneNumberIdentificationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberIdentificationResultRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationResultRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetPhoneNumberIdentificationResultRequest) GetOutId() *string {
	return s.OutId
}

func (s *GetPhoneNumberIdentificationResultRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPhoneNumberIdentificationResultRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPhoneNumberIdentificationResultRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPhoneNumberIdentificationResultRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPhoneNumberIdentificationResultRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetPhoneNumberIdentificationResultRequest) GetSessionPayload() *string {
	return s.SessionPayload
}

func (s *GetPhoneNumberIdentificationResultRequest) SetAuthCode(v string) *GetPhoneNumberIdentificationResultRequest {
	s.AuthCode = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetOutId(v string) *GetPhoneNumberIdentificationResultRequest {
	s.OutId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetOwnerId(v int64) *GetPhoneNumberIdentificationResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetPhoneNumber(v string) *GetPhoneNumberIdentificationResultRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetResourceOwnerAccount(v string) *GetPhoneNumberIdentificationResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetResourceOwnerId(v int64) *GetPhoneNumberIdentificationResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetSessionId(v string) *GetPhoneNumberIdentificationResultRequest {
	s.SessionId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetSessionPayload(v string) *GetPhoneNumberIdentificationResultRequest {
	s.SessionPayload = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) Validate() error {
	return dara.Validate(s)
}
