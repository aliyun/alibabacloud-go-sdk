// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCodeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCodeSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateCodeSourceResponseBody) *CreateCodeSourceResponse
	GetBody() *CreateCodeSourceResponseBody
}

type CreateCodeSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCodeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCodeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCodeSourceResponse) GetBody() *CreateCodeSourceResponseBody {
	return s.Body
}

func (s *CreateCodeSourceResponse) SetHeaders(v map[string]*string) *CreateCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateCodeSourceResponse) SetStatusCode(v int32) *CreateCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCodeSourceResponse) SetBody(v *CreateCodeSourceResponseBody) *CreateCodeSourceResponse {
	s.Body = v
	return s
}

func (s *CreateCodeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
