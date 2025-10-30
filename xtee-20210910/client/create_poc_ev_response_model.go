// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocEvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePocEvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePocEvResponse
	GetStatusCode() *int32
	SetBody(v *CreatePocEvResponseBody) *CreatePocEvResponse
	GetBody() *CreatePocEvResponseBody
}

type CreatePocEvResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePocEvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePocEvResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePocEvResponse) GoString() string {
	return s.String()
}

func (s *CreatePocEvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePocEvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePocEvResponse) GetBody() *CreatePocEvResponseBody {
	return s.Body
}

func (s *CreatePocEvResponse) SetHeaders(v map[string]*string) *CreatePocEvResponse {
	s.Headers = v
	return s
}

func (s *CreatePocEvResponse) SetStatusCode(v int32) *CreatePocEvResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePocEvResponse) SetBody(v *CreatePocEvResponseBody) *CreatePocEvResponse {
	s.Body = v
	return s
}

func (s *CreatePocEvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
