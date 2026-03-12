// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryContactResponse
	GetStatusCode() *int32
	SetBody(v *QueryContactResponseBody) *QueryContactResponse
	GetBody() *QueryContactResponseBody
}

type QueryContactResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContactResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryContactResponse) GoString() string {
	return s.String()
}

func (s *QueryContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryContactResponse) GetBody() *QueryContactResponseBody {
	return s.Body
}

func (s *QueryContactResponse) SetHeaders(v map[string]*string) *QueryContactResponse {
	s.Headers = v
	return s
}

func (s *QueryContactResponse) SetStatusCode(v int32) *QueryContactResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContactResponse) SetBody(v *QueryContactResponseBody) *QueryContactResponse {
	s.Body = v
	return s
}

func (s *QueryContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
