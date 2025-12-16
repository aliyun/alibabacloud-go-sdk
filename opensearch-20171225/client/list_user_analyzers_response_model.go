// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAnalyzersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserAnalyzersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserAnalyzersResponse
	GetStatusCode() *int32
	SetBody(v *ListUserAnalyzersResponseBody) *ListUserAnalyzersResponse
	GetBody() *ListUserAnalyzersResponseBody
}

type ListUserAnalyzersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserAnalyzersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserAnalyzersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserAnalyzersResponse) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserAnalyzersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserAnalyzersResponse) GetBody() *ListUserAnalyzersResponseBody {
	return s.Body
}

func (s *ListUserAnalyzersResponse) SetHeaders(v map[string]*string) *ListUserAnalyzersResponse {
	s.Headers = v
	return s
}

func (s *ListUserAnalyzersResponse) SetStatusCode(v int32) *ListUserAnalyzersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserAnalyzersResponse) SetBody(v *ListUserAnalyzersResponseBody) *ListUserAnalyzersResponse {
	s.Body = v
	return s
}

func (s *ListUserAnalyzersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
