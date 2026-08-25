// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetUserIdRequest
	GetDirectoryId() *string
	SetExternalId(v *GetUserIdRequestExternalId) *GetUserIdRequest
	GetExternalId() *GetUserIdRequestExternalId
}

type GetUserIdRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The identifier information about the user that is synchronized from an external identity provider (IdP).
	ExternalId *GetUserIdRequestExternalId `json:"ExternalId,omitempty" xml:"ExternalId,omitempty" type:"Struct"`
}

func (s GetUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserIdRequest) GetExternalId() *GetUserIdRequestExternalId {
	return s.ExternalId
}

func (s *GetUserIdRequest) SetDirectoryId(v string) *GetUserIdRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserIdRequest) SetExternalId(v *GetUserIdRequestExternalId) *GetUserIdRequest {
	s.ExternalId = v
	return s
}

func (s *GetUserIdRequest) Validate() error {
	if s.ExternalId != nil {
		if err := s.ExternalId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserIdRequestExternalId struct {
	// The identifier of the user that is synchronized from an external IdP.
	//
	// example:
	//
	// c73******a5fdd5
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The method for external identity synchronization. Only System for Cross-domain Identity Management (SCIM) synchronization is supported.
	//
	// example:
	//
	// SCIM
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s GetUserIdRequestExternalId) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdRequestExternalId) GoString() string {
	return s.String()
}

func (s *GetUserIdRequestExternalId) GetId() *string {
	return s.Id
}

func (s *GetUserIdRequestExternalId) GetIssuer() *string {
	return s.Issuer
}

func (s *GetUserIdRequestExternalId) SetId(v string) *GetUserIdRequestExternalId {
	s.Id = &v
	return s
}

func (s *GetUserIdRequestExternalId) SetIssuer(v string) *GetUserIdRequestExternalId {
	s.Issuer = &v
	return s
}

func (s *GetUserIdRequestExternalId) Validate() error {
	return dara.Validate(s)
}
