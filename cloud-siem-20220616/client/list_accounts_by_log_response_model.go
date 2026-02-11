// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsByLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccountsByLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccountsByLogResponse
	GetStatusCode() *int32
	SetBody(v *ListAccountsByLogResponseBody) *ListAccountsByLogResponse
	GetBody() *ListAccountsByLogResponseBody
}

type ListAccountsByLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsByLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsByLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsByLogResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsByLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccountsByLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccountsByLogResponse) GetBody() *ListAccountsByLogResponseBody {
	return s.Body
}

func (s *ListAccountsByLogResponse) SetHeaders(v map[string]*string) *ListAccountsByLogResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsByLogResponse) SetStatusCode(v int32) *ListAccountsByLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsByLogResponse) SetBody(v *ListAccountsByLogResponseBody) *ListAccountsByLogResponse {
	s.Body = v
	return s
}

func (s *ListAccountsByLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
