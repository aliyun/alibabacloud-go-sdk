// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSagInstanceFromCcnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeSagInstanceFromCcnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeSagInstanceFromCcnResponse
	GetStatusCode() *int32
	SetBody(v *RevokeSagInstanceFromCcnResponseBody) *RevokeSagInstanceFromCcnResponse
	GetBody() *RevokeSagInstanceFromCcnResponseBody
}

type RevokeSagInstanceFromCcnResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeSagInstanceFromCcnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeSagInstanceFromCcnResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeSagInstanceFromCcnResponse) GoString() string {
	return s.String()
}

func (s *RevokeSagInstanceFromCcnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeSagInstanceFromCcnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeSagInstanceFromCcnResponse) GetBody() *RevokeSagInstanceFromCcnResponseBody {
	return s.Body
}

func (s *RevokeSagInstanceFromCcnResponse) SetHeaders(v map[string]*string) *RevokeSagInstanceFromCcnResponse {
	s.Headers = v
	return s
}

func (s *RevokeSagInstanceFromCcnResponse) SetStatusCode(v int32) *RevokeSagInstanceFromCcnResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeSagInstanceFromCcnResponse) SetBody(v *RevokeSagInstanceFromCcnResponseBody) *RevokeSagInstanceFromCcnResponse {
	s.Body = v
	return s
}

func (s *RevokeSagInstanceFromCcnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
