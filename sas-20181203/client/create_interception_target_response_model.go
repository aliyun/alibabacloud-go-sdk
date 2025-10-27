// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterceptionTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInterceptionTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInterceptionTargetResponse
	GetStatusCode() *int32
	SetBody(v *CreateInterceptionTargetResponseBody) *CreateInterceptionTargetResponse
	GetBody() *CreateInterceptionTargetResponseBody
}

type CreateInterceptionTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInterceptionTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInterceptionTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInterceptionTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateInterceptionTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInterceptionTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInterceptionTargetResponse) GetBody() *CreateInterceptionTargetResponseBody {
	return s.Body
}

func (s *CreateInterceptionTargetResponse) SetHeaders(v map[string]*string) *CreateInterceptionTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateInterceptionTargetResponse) SetStatusCode(v int32) *CreateInterceptionTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInterceptionTargetResponse) SetBody(v *CreateInterceptionTargetResponseBody) *CreateInterceptionTargetResponse {
	s.Body = v
	return s
}

func (s *CreateInterceptionTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
