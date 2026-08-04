// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindContacterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindContacterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindContacterResponse
	GetStatusCode() *int32
	SetBody(v *FindContacterResponseBody) *FindContacterResponse
	GetBody() *FindContacterResponseBody
}

type FindContacterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindContacterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindContacterResponse) String() string {
	return dara.Prettify(s)
}

func (s FindContacterResponse) GoString() string {
	return s.String()
}

func (s *FindContacterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindContacterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindContacterResponse) GetBody() *FindContacterResponseBody {
	return s.Body
}

func (s *FindContacterResponse) SetHeaders(v map[string]*string) *FindContacterResponse {
	s.Headers = v
	return s
}

func (s *FindContacterResponse) SetStatusCode(v int32) *FindContacterResponse {
	s.StatusCode = &v
	return s
}

func (s *FindContacterResponse) SetBody(v *FindContacterResponseBody) *FindContacterResponse {
	s.Body = v
	return s
}

func (s *FindContacterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
