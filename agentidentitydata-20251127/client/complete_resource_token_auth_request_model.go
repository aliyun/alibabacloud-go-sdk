// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteResourceTokenAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionURI(v string) *CompleteResourceTokenAuthRequest
	GetSessionURI() *string
	SetUserIdentifier(v *CompleteResourceTokenAuthRequestUserIdentifier) *CompleteResourceTokenAuthRequest
	GetUserIdentifier() *CompleteResourceTokenAuthRequestUserIdentifier
}

type CompleteResourceTokenAuthRequest struct {
	// example:
	//
	// urn:ietf:params:oauth:request_uri:3030cabc-****-****-****-d0054944102a
	SessionURI     *string                                         `json:"SessionURI,omitempty" xml:"SessionURI,omitempty"`
	UserIdentifier *CompleteResourceTokenAuthRequestUserIdentifier `json:"UserIdentifier,omitempty" xml:"UserIdentifier,omitempty" type:"Struct"`
}

func (s CompleteResourceTokenAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteResourceTokenAuthRequest) GoString() string {
	return s.String()
}

func (s *CompleteResourceTokenAuthRequest) GetSessionURI() *string {
	return s.SessionURI
}

func (s *CompleteResourceTokenAuthRequest) GetUserIdentifier() *CompleteResourceTokenAuthRequestUserIdentifier {
	return s.UserIdentifier
}

func (s *CompleteResourceTokenAuthRequest) SetSessionURI(v string) *CompleteResourceTokenAuthRequest {
	s.SessionURI = &v
	return s
}

func (s *CompleteResourceTokenAuthRequest) SetUserIdentifier(v *CompleteResourceTokenAuthRequestUserIdentifier) *CompleteResourceTokenAuthRequest {
	s.UserIdentifier = v
	return s
}

func (s *CompleteResourceTokenAuthRequest) Validate() error {
	if s.UserIdentifier != nil {
		if err := s.UserIdentifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompleteResourceTokenAuthRequestUserIdentifier struct {
	// example:
	//
	// externalUser
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// eyAgImFsZyI6ICJSU****
	UserJWT *string `json:"UserJWT,omitempty" xml:"UserJWT,omitempty"`
}

func (s CompleteResourceTokenAuthRequestUserIdentifier) String() string {
	return dara.Prettify(s)
}

func (s CompleteResourceTokenAuthRequestUserIdentifier) GoString() string {
	return s.String()
}

func (s *CompleteResourceTokenAuthRequestUserIdentifier) GetUserId() *string {
	return s.UserId
}

func (s *CompleteResourceTokenAuthRequestUserIdentifier) GetUserJWT() *string {
	return s.UserJWT
}

func (s *CompleteResourceTokenAuthRequestUserIdentifier) SetUserId(v string) *CompleteResourceTokenAuthRequestUserIdentifier {
	s.UserId = &v
	return s
}

func (s *CompleteResourceTokenAuthRequestUserIdentifier) SetUserJWT(v string) *CompleteResourceTokenAuthRequestUserIdentifier {
	s.UserJWT = &v
	return s
}

func (s *CompleteResourceTokenAuthRequestUserIdentifier) Validate() error {
	return dara.Validate(s)
}
