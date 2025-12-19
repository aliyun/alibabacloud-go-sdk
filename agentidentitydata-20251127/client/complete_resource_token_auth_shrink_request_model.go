// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteResourceTokenAuthShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionURI(v string) *CompleteResourceTokenAuthShrinkRequest
	GetSessionURI() *string
	SetUserIdentifierShrink(v string) *CompleteResourceTokenAuthShrinkRequest
	GetUserIdentifierShrink() *string
}

type CompleteResourceTokenAuthShrinkRequest struct {
	// example:
	//
	// urn:ietf:params:oauth:request_uri:3030cabc-****-****-****-d0054944102a
	SessionURI           *string `json:"SessionURI,omitempty" xml:"SessionURI,omitempty"`
	UserIdentifierShrink *string `json:"UserIdentifier,omitempty" xml:"UserIdentifier,omitempty"`
}

func (s CompleteResourceTokenAuthShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteResourceTokenAuthShrinkRequest) GoString() string {
	return s.String()
}

func (s *CompleteResourceTokenAuthShrinkRequest) GetSessionURI() *string {
	return s.SessionURI
}

func (s *CompleteResourceTokenAuthShrinkRequest) GetUserIdentifierShrink() *string {
	return s.UserIdentifierShrink
}

func (s *CompleteResourceTokenAuthShrinkRequest) SetSessionURI(v string) *CompleteResourceTokenAuthShrinkRequest {
	s.SessionURI = &v
	return s
}

func (s *CompleteResourceTokenAuthShrinkRequest) SetUserIdentifierShrink(v string) *CompleteResourceTokenAuthShrinkRequest {
	s.UserIdentifierShrink = &v
	return s
}

func (s *CompleteResourceTokenAuthShrinkRequest) Validate() error {
	return dara.Validate(s)
}
