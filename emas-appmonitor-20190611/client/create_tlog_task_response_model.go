// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTlogTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTlogTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTlogTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTlogTaskResponseBody) *CreateTlogTaskResponse
	GetBody() *CreateTlogTaskResponseBody
}

type CreateTlogTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTlogTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTlogTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTlogTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTlogTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTlogTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTlogTaskResponse) GetBody() *CreateTlogTaskResponseBody {
	return s.Body
}

func (s *CreateTlogTaskResponse) SetHeaders(v map[string]*string) *CreateTlogTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTlogTaskResponse) SetStatusCode(v int32) *CreateTlogTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTlogTaskResponse) SetBody(v *CreateTlogTaskResponseBody) *CreateTlogTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTlogTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
