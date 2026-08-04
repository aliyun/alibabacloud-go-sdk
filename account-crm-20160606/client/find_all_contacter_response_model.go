// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindAllContacterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindAllContacterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindAllContacterResponse
	GetStatusCode() *int32
	SetBody(v *FindAllContacterResponseBody) *FindAllContacterResponse
	GetBody() *FindAllContacterResponseBody
}

type FindAllContacterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindAllContacterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindAllContacterResponse) String() string {
	return dara.Prettify(s)
}

func (s FindAllContacterResponse) GoString() string {
	return s.String()
}

func (s *FindAllContacterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindAllContacterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindAllContacterResponse) GetBody() *FindAllContacterResponseBody {
	return s.Body
}

func (s *FindAllContacterResponse) SetHeaders(v map[string]*string) *FindAllContacterResponse {
	s.Headers = v
	return s
}

func (s *FindAllContacterResponse) SetStatusCode(v int32) *FindAllContacterResponse {
	s.StatusCode = &v
	return s
}

func (s *FindAllContacterResponse) SetBody(v *FindAllContacterResponseBody) *FindAllContacterResponse {
	s.Body = v
	return s
}

func (s *FindAllContacterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
