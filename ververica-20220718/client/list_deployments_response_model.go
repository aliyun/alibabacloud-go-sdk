// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeploymentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeploymentsResponse
	GetStatusCode() *int32
	SetBody(v *ListDeploymentsResponseBody) *ListDeploymentsResponse
	GetBody() *ListDeploymentsResponseBody
}

type ListDeploymentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeploymentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeploymentsResponse) GetBody() *ListDeploymentsResponseBody {
	return s.Body
}

func (s *ListDeploymentsResponse) SetHeaders(v map[string]*string) *ListDeploymentsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentsResponse) SetStatusCode(v int32) *ListDeploymentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentsResponse) SetBody(v *ListDeploymentsResponseBody) *ListDeploymentsResponse {
	s.Body = v
	return s
}

func (s *ListDeploymentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
