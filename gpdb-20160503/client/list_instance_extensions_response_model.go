// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceExtensionsResponseBody) *ListInstanceExtensionsResponse
	GetBody() *ListInstanceExtensionsResponseBody
}

type ListInstanceExtensionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceExtensionsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceExtensionsResponse) GetBody() *ListInstanceExtensionsResponseBody {
	return s.Body
}

func (s *ListInstanceExtensionsResponse) SetHeaders(v map[string]*string) *ListInstanceExtensionsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceExtensionsResponse) SetStatusCode(v int32) *ListInstanceExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceExtensionsResponse) SetBody(v *ListInstanceExtensionsResponseBody) *ListInstanceExtensionsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
