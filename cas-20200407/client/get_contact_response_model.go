// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContactResponse
	GetStatusCode() *int32
	SetBody(v *GetContactResponseBody) *GetContactResponse
	GetBody() *GetContactResponseBody
}

type GetContactResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContactResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContactResponse) GoString() string {
	return s.String()
}

func (s *GetContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContactResponse) GetBody() *GetContactResponseBody {
	return s.Body
}

func (s *GetContactResponse) SetHeaders(v map[string]*string) *GetContactResponse {
	s.Headers = v
	return s
}

func (s *GetContactResponse) SetStatusCode(v int32) *GetContactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContactResponse) SetBody(v *GetContactResponseBody) *GetContactResponse {
	s.Body = v
	return s
}

func (s *GetContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
