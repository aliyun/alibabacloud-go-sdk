// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListLogRegionsResponseBody) *ListLogRegionsResponse
	GetBody() *ListLogRegionsResponseBody
}

type ListLogRegionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListLogRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogRegionsResponse) GetBody() *ListLogRegionsResponseBody {
	return s.Body
}

func (s *ListLogRegionsResponse) SetHeaders(v map[string]*string) *ListLogRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListLogRegionsResponse) SetStatusCode(v int32) *ListLogRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogRegionsResponse) SetBody(v *ListLogRegionsResponseBody) *ListLogRegionsResponse {
	s.Body = v
	return s
}

func (s *ListLogRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
