// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeClustersResponseBody) *ListComputeClustersResponse
	GetBody() *ListComputeClustersResponseBody
}

type ListComputeClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeClustersResponse) GoString() string {
	return s.String()
}

func (s *ListComputeClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeClustersResponse) GetBody() *ListComputeClustersResponseBody {
	return s.Body
}

func (s *ListComputeClustersResponse) SetHeaders(v map[string]*string) *ListComputeClustersResponse {
	s.Headers = v
	return s
}

func (s *ListComputeClustersResponse) SetStatusCode(v int32) *ListComputeClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeClustersResponse) SetBody(v *ListComputeClustersResponseBody) *ListComputeClustersResponse {
	s.Body = v
	return s
}

func (s *ListComputeClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
