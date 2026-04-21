// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOmniAnswerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OmniSearchQuery) *OmniAnswerRequest
	GetBody() *OmniSearchQuery
}

type OmniAnswerRequest struct {
	Body *OmniSearchQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OmniAnswerRequest) String() string {
	return dara.Prettify(s)
}

func (s OmniAnswerRequest) GoString() string {
	return s.String()
}

func (s *OmniAnswerRequest) GetBody() *OmniSearchQuery {
	return s.Body
}

func (s *OmniAnswerRequest) SetBody(v *OmniSearchQuery) *OmniAnswerRequest {
	s.Body = v
	return s
}

func (s *OmniAnswerRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
