// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateSessionInput) *UpdateSessionRequest
	GetBody() *UpdateSessionInput
	SetQualifier(v string) *UpdateSessionRequest
	GetQualifier() *string
}

type UpdateSessionRequest struct {
	Body *UpdateSessionInput `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// aliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s UpdateSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSessionRequest) GetBody() *UpdateSessionInput {
	return s.Body
}

func (s *UpdateSessionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *UpdateSessionRequest) SetBody(v *UpdateSessionInput) *UpdateSessionRequest {
	s.Body = v
	return s
}

func (s *UpdateSessionRequest) SetQualifier(v string) *UpdateSessionRequest {
	s.Qualifier = &v
	return s
}

func (s *UpdateSessionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
