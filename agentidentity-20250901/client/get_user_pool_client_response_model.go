// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserPoolClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserPoolClientResponse
	GetStatusCode() *int32
	SetBody(v *GetUserPoolClientResponseBody) *GetUserPoolClientResponse
	GetBody() *GetUserPoolClientResponseBody
}

type GetUserPoolClientResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserPoolClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserPoolClientResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolClientResponse) GoString() string {
	return s.String()
}

func (s *GetUserPoolClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserPoolClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserPoolClientResponse) GetBody() *GetUserPoolClientResponseBody {
	return s.Body
}

func (s *GetUserPoolClientResponse) SetHeaders(v map[string]*string) *GetUserPoolClientResponse {
	s.Headers = v
	return s
}

func (s *GetUserPoolClientResponse) SetStatusCode(v int32) *GetUserPoolClientResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserPoolClientResponse) SetBody(v *GetUserPoolClientResponseBody) *GetUserPoolClientResponse {
	s.Body = v
	return s
}

func (s *GetUserPoolClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
