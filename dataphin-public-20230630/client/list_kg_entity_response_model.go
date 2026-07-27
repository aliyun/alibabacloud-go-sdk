// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKgEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKgEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKgEntityResponse
	GetStatusCode() *int32
	SetBody(v *ListKgEntityResponseBody) *ListKgEntityResponse
	GetBody() *ListKgEntityResponseBody
}

type ListKgEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKgEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKgEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityResponse) GoString() string {
	return s.String()
}

func (s *ListKgEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKgEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKgEntityResponse) GetBody() *ListKgEntityResponseBody {
	return s.Body
}

func (s *ListKgEntityResponse) SetHeaders(v map[string]*string) *ListKgEntityResponse {
	s.Headers = v
	return s
}

func (s *ListKgEntityResponse) SetStatusCode(v int32) *ListKgEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKgEntityResponse) SetBody(v *ListKgEntityResponseBody) *ListKgEntityResponse {
	s.Body = v
	return s
}

func (s *ListKgEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
