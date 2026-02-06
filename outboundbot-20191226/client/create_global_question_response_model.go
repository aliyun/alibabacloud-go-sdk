// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGlobalQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGlobalQuestionResponse
	GetStatusCode() *int32
	SetBody(v *CreateGlobalQuestionResponseBody) *CreateGlobalQuestionResponse
	GetBody() *CreateGlobalQuestionResponseBody
}

type CreateGlobalQuestionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalQuestionResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGlobalQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGlobalQuestionResponse) GetBody() *CreateGlobalQuestionResponseBody {
	return s.Body
}

func (s *CreateGlobalQuestionResponse) SetHeaders(v map[string]*string) *CreateGlobalQuestionResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalQuestionResponse) SetStatusCode(v int32) *CreateGlobalQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalQuestionResponse) SetBody(v *CreateGlobalQuestionResponseBody) *CreateGlobalQuestionResponse {
	s.Body = v
	return s
}

func (s *CreateGlobalQuestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
