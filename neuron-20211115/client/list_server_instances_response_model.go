// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServerInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServerInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ServerInstanceResult) *ListServerInstancesResponse
	GetBody() *ServerInstanceResult
}

type ListServerInstancesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ServerInstanceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServerInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServerInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServerInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServerInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServerInstancesResponse) GetBody() *ServerInstanceResult {
	return s.Body
}

func (s *ListServerInstancesResponse) SetHeaders(v map[string]*string) *ListServerInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServerInstancesResponse) SetStatusCode(v int32) *ListServerInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerInstancesResponse) SetBody(v *ServerInstanceResult) *ListServerInstancesResponse {
	s.Body = v
	return s
}

func (s *ListServerInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
