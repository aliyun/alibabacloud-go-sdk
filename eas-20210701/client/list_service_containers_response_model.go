// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceContainersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceContainersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceContainersResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceContainersResponseBody) *ListServiceContainersResponse
	GetBody() *ListServiceContainersResponseBody
}

type ListServiceContainersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceContainersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceContainersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceContainersResponse) GoString() string {
	return s.String()
}

func (s *ListServiceContainersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceContainersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceContainersResponse) GetBody() *ListServiceContainersResponseBody {
	return s.Body
}

func (s *ListServiceContainersResponse) SetHeaders(v map[string]*string) *ListServiceContainersResponse {
	s.Headers = v
	return s
}

func (s *ListServiceContainersResponse) SetStatusCode(v int32) *ListServiceContainersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceContainersResponse) SetBody(v *ListServiceContainersResponseBody) *ListServiceContainersResponse {
	s.Body = v
	return s
}

func (s *ListServiceContainersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
