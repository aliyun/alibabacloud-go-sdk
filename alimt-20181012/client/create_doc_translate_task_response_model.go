// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDocTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDocTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDocTranslateTaskResponseBody) *CreateDocTranslateTaskResponse
	GetBody() *CreateDocTranslateTaskResponseBody
}

type CreateDocTranslateTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDocTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDocTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDocTranslateTaskResponse) GetBody() *CreateDocTranslateTaskResponseBody {
	return s.Body
}

func (s *CreateDocTranslateTaskResponse) SetHeaders(v map[string]*string) *CreateDocTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDocTranslateTaskResponse) SetStatusCode(v int32) *CreateDocTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocTranslateTaskResponse) SetBody(v *CreateDocTranslateTaskResponseBody) *CreateDocTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDocTranslateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
