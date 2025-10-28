// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateSessionInput) *CreateSessionRequest
	GetBody() *CreateSessionInput
	SetQualifier(v string) *CreateSessionRequest
	GetQualifier() *string
}

type CreateSessionRequest struct {
	Body *CreateSessionInput `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// aliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) GetBody() *CreateSessionInput {
	return s.Body
}

func (s *CreateSessionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *CreateSessionRequest) SetBody(v *CreateSessionInput) *CreateSessionRequest {
	s.Body = v
	return s
}

func (s *CreateSessionRequest) SetQualifier(v string) *CreateSessionRequest {
	s.Qualifier = &v
	return s
}

func (s *CreateSessionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
