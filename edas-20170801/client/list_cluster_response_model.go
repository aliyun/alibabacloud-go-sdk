// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterResponseBody) *ListClusterResponse
	GetBody() *ListClusterResponseBody
}

type ListClusterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterResponse) GoString() string {
	return s.String()
}

func (s *ListClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterResponse) GetBody() *ListClusterResponseBody {
	return s.Body
}

func (s *ListClusterResponse) SetHeaders(v map[string]*string) *ListClusterResponse {
	s.Headers = v
	return s
}

func (s *ListClusterResponse) SetStatusCode(v int32) *ListClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterResponse) SetBody(v *ListClusterResponseBody) *ListClusterResponse {
	s.Body = v
	return s
}

func (s *ListClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
