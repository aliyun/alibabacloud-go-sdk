// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceIndicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceIndicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceIndicesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceIndicesResponseBody) *ListInstanceIndicesResponse
	GetBody() *ListInstanceIndicesResponseBody
}

type ListInstanceIndicesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceIndicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceIndicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceIndicesResponse) GetBody() *ListInstanceIndicesResponseBody {
	return s.Body
}

func (s *ListInstanceIndicesResponse) SetHeaders(v map[string]*string) *ListInstanceIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceIndicesResponse) SetStatusCode(v int32) *ListInstanceIndicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceIndicesResponse) SetBody(v *ListInstanceIndicesResponseBody) *ListInstanceIndicesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceIndicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
