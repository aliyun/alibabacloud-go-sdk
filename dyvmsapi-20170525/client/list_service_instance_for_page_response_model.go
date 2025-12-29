// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceForPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceInstanceForPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceInstanceForPageResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceInstanceForPageResponseBody) *ListServiceInstanceForPageResponse
	GetBody() *ListServiceInstanceForPageResponseBody
}

type ListServiceInstanceForPageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceForPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceForPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceForPageResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceForPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceInstanceForPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceInstanceForPageResponse) GetBody() *ListServiceInstanceForPageResponseBody {
	return s.Body
}

func (s *ListServiceInstanceForPageResponse) SetHeaders(v map[string]*string) *ListServiceInstanceForPageResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceForPageResponse) SetStatusCode(v int32) *ListServiceInstanceForPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceForPageResponse) SetBody(v *ListServiceInstanceForPageResponseBody) *ListServiceInstanceForPageResponse {
	s.Body = v
	return s
}

func (s *ListServiceInstanceForPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
