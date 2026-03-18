// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenedAccessLogInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOpenedAccessLogInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOpenedAccessLogInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListOpenedAccessLogInstancesResponseBody) *ListOpenedAccessLogInstancesResponse
	GetBody() *ListOpenedAccessLogInstancesResponseBody
}

type ListOpenedAccessLogInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpenedAccessLogInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOpenedAccessLogInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOpenedAccessLogInstancesResponse) GetBody() *ListOpenedAccessLogInstancesResponseBody {
	return s.Body
}

func (s *ListOpenedAccessLogInstancesResponse) SetHeaders(v map[string]*string) *ListOpenedAccessLogInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetStatusCode(v int32) *ListOpenedAccessLogInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetBody(v *ListOpenedAccessLogInstancesResponseBody) *ListOpenedAccessLogInstancesResponse {
	s.Body = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
