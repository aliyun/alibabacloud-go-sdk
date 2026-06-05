// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeWorkersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeWorkersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeWorkersResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeWorkersResponseBody) *ListEdgeWorkersResponse
	GetBody() *ListEdgeWorkersResponseBody
}

type ListEdgeWorkersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeWorkersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeWorkersResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeWorkersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeWorkersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeWorkersResponse) GetBody() *ListEdgeWorkersResponseBody {
	return s.Body
}

func (s *ListEdgeWorkersResponse) SetHeaders(v map[string]*string) *ListEdgeWorkersResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeWorkersResponse) SetStatusCode(v int32) *ListEdgeWorkersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeWorkersResponse) SetBody(v *ListEdgeWorkersResponseBody) *ListEdgeWorkersResponse {
	s.Body = v
	return s
}

func (s *ListEdgeWorkersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
