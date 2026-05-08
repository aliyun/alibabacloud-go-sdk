// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIllustrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *IllustrationTaskCreateCmd) *CreateIllustrationTaskRequest
	GetBody() *IllustrationTaskCreateCmd
}

type CreateIllustrationTaskRequest struct {
	Body *IllustrationTaskCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIllustrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIllustrationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateIllustrationTaskRequest) GetBody() *IllustrationTaskCreateCmd {
	return s.Body
}

func (s *CreateIllustrationTaskRequest) SetBody(v *IllustrationTaskCreateCmd) *CreateIllustrationTaskRequest {
	s.Body = v
	return s
}

func (s *CreateIllustrationTaskRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
