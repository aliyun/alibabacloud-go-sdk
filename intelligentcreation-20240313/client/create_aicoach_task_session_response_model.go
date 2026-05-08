// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAICoachTaskSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAICoachTaskSessionResponse
	GetStatusCode() *int32
	SetBody(v *CreateAICoachTaskSessionResponseBody) *CreateAICoachTaskSessionResponse
	GetBody() *CreateAICoachTaskSessionResponseBody
}

type CreateAICoachTaskSessionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAICoachTaskSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAICoachTaskSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAICoachTaskSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAICoachTaskSessionResponse) GetBody() *CreateAICoachTaskSessionResponseBody {
	return s.Body
}

func (s *CreateAICoachTaskSessionResponse) SetHeaders(v map[string]*string) *CreateAICoachTaskSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateAICoachTaskSessionResponse) SetStatusCode(v int32) *CreateAICoachTaskSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAICoachTaskSessionResponse) SetBody(v *CreateAICoachTaskSessionResponseBody) *CreateAICoachTaskSessionResponse {
	s.Body = v
	return s
}

func (s *CreateAICoachTaskSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
