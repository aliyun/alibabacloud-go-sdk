// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserPoolResponse
	GetStatusCode() *int32
	SetBody(v *GetUserPoolResponseBody) *GetUserPoolResponse
	GetBody() *GetUserPoolResponseBody
}

type GetUserPoolResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolResponse) GoString() string {
	return s.String()
}

func (s *GetUserPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserPoolResponse) GetBody() *GetUserPoolResponseBody {
	return s.Body
}

func (s *GetUserPoolResponse) SetHeaders(v map[string]*string) *GetUserPoolResponse {
	s.Headers = v
	return s
}

func (s *GetUserPoolResponse) SetStatusCode(v int32) *GetUserPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserPoolResponse) SetBody(v *GetUserPoolResponseBody) *GetUserPoolResponse {
	s.Body = v
	return s
}

func (s *GetUserPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
