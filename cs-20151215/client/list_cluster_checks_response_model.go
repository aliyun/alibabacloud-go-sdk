// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterChecksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterChecksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterChecksResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterChecksResponseBody) *ListClusterChecksResponse
	GetBody() *ListClusterChecksResponseBody
}

type ListClusterChecksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterChecksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterChecksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterChecksResponse) GoString() string {
	return s.String()
}

func (s *ListClusterChecksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterChecksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterChecksResponse) GetBody() *ListClusterChecksResponseBody {
	return s.Body
}

func (s *ListClusterChecksResponse) SetHeaders(v map[string]*string) *ListClusterChecksResponse {
	s.Headers = v
	return s
}

func (s *ListClusterChecksResponse) SetStatusCode(v int32) *ListClusterChecksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterChecksResponse) SetBody(v *ListClusterChecksResponseBody) *ListClusterChecksResponse {
	s.Body = v
	return s
}

func (s *ListClusterChecksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
