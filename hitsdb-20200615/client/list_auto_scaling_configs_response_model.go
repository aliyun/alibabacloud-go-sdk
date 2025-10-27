// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoScalingConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoScalingConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoScalingConfigsResponseBody) *ListAutoScalingConfigsResponse
	GetBody() *ListAutoScalingConfigsResponseBody
}

type ListAutoScalingConfigsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoScalingConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoScalingConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoScalingConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoScalingConfigsResponse) GetBody() *ListAutoScalingConfigsResponseBody {
	return s.Body
}

func (s *ListAutoScalingConfigsResponse) SetHeaders(v map[string]*string) *ListAutoScalingConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAutoScalingConfigsResponse) SetStatusCode(v int32) *ListAutoScalingConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoScalingConfigsResponse) SetBody(v *ListAutoScalingConfigsResponseBody) *ListAutoScalingConfigsResponse {
	s.Body = v
	return s
}

func (s *ListAutoScalingConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
