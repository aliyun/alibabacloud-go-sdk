// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvsResponseBody) *ListEnvsResponse
	GetBody() *ListEnvsResponseBody
}

type ListEnvsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvsResponse) GetBody() *ListEnvsResponseBody {
	return s.Body
}

func (s *ListEnvsResponse) SetHeaders(v map[string]*string) *ListEnvsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvsResponse) SetStatusCode(v int32) *ListEnvsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvsResponse) SetBody(v *ListEnvsResponseBody) *ListEnvsResponse {
	s.Body = v
	return s
}

func (s *ListEnvsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
