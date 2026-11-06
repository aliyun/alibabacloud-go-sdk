// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceResponseBody) *ListInstanceResponse
	GetBody() *ListInstanceResponseBody
}

type ListInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceResponse) GetBody() *ListInstanceResponseBody {
	return s.Body
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetStatusCode(v int32) *ListInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

func (s *ListInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
