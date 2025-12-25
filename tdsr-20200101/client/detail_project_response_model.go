// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetailProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetailProjectResponse
	GetStatusCode() *int32
	SetBody(v *DetailProjectResponseBody) *DetailProjectResponse
	GetBody() *DetailProjectResponseBody
}

type DetailProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetailProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetailProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DetailProjectResponse) GoString() string {
	return s.String()
}

func (s *DetailProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetailProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetailProjectResponse) GetBody() *DetailProjectResponseBody {
	return s.Body
}

func (s *DetailProjectResponse) SetHeaders(v map[string]*string) *DetailProjectResponse {
	s.Headers = v
	return s
}

func (s *DetailProjectResponse) SetStatusCode(v int32) *DetailProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailProjectResponse) SetBody(v *DetailProjectResponseBody) *DetailProjectResponse {
	s.Body = v
	return s
}

func (s *DetailProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
