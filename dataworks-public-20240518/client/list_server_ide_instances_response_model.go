// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServerIdeInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServerIdeInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListServerIdeInstancesResponseBody) *ListServerIdeInstancesResponse
	GetBody() *ListServerIdeInstancesResponseBody
}

type ListServerIdeInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServerIdeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServerIdeInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServerIdeInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServerIdeInstancesResponse) GetBody() *ListServerIdeInstancesResponseBody {
	return s.Body
}

func (s *ListServerIdeInstancesResponse) SetHeaders(v map[string]*string) *ListServerIdeInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServerIdeInstancesResponse) SetStatusCode(v int32) *ListServerIdeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerIdeInstancesResponse) SetBody(v *ListServerIdeInstancesResponseBody) *ListServerIdeInstancesResponse {
	s.Body = v
	return s
}

func (s *ListServerIdeInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
