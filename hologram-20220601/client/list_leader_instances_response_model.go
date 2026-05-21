// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLeaderInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLeaderInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLeaderInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListLeaderInstancesResponseBody) *ListLeaderInstancesResponse
	GetBody() *ListLeaderInstancesResponseBody
}

type ListLeaderInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLeaderInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLeaderInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLeaderInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListLeaderInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLeaderInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLeaderInstancesResponse) GetBody() *ListLeaderInstancesResponseBody {
	return s.Body
}

func (s *ListLeaderInstancesResponse) SetHeaders(v map[string]*string) *ListLeaderInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListLeaderInstancesResponse) SetStatusCode(v int32) *ListLeaderInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLeaderInstancesResponse) SetBody(v *ListLeaderInstancesResponseBody) *ListLeaderInstancesResponse {
	s.Body = v
	return s
}

func (s *ListLeaderInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
