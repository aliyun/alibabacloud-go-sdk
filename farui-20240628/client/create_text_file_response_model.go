// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTextFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTextFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTextFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateTextFileResponseBody) *CreateTextFileResponse
	GetBody() *CreateTextFileResponseBody
}

type CreateTextFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTextFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTextFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTextFileResponse) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTextFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTextFileResponse) GetBody() *CreateTextFileResponseBody {
	return s.Body
}

func (s *CreateTextFileResponse) SetHeaders(v map[string]*string) *CreateTextFileResponse {
	s.Headers = v
	return s
}

func (s *CreateTextFileResponse) SetStatusCode(v int32) *CreateTextFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTextFileResponse) SetBody(v *CreateTextFileResponseBody) *CreateTextFileResponse {
	s.Body = v
	return s
}

func (s *CreateTextFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
