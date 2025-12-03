// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageTranslateTaskResponseBody) *CreateImageTranslateTaskResponse
	GetBody() *CreateImageTranslateTaskResponseBody
}

type CreateImageTranslateTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageTranslateTaskResponse) GetBody() *CreateImageTranslateTaskResponseBody {
	return s.Body
}

func (s *CreateImageTranslateTaskResponse) SetHeaders(v map[string]*string) *CreateImageTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageTranslateTaskResponse) SetStatusCode(v int32) *CreateImageTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageTranslateTaskResponse) SetBody(v *CreateImageTranslateTaskResponseBody) *CreateImageTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageTranslateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
