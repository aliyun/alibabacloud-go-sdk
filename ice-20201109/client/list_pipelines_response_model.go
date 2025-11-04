// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelinesResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelinesResponseBody) *ListPipelinesResponse
	GetBody() *ListPipelinesResponseBody
}

type ListPipelinesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelinesResponse) GetBody() *ListPipelinesResponseBody {
	return s.Body
}

func (s *ListPipelinesResponse) SetHeaders(v map[string]*string) *ListPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelinesResponse) SetStatusCode(v int32) *ListPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelinesResponse) SetBody(v *ListPipelinesResponseBody) *ListPipelinesResponse {
	s.Body = v
	return s
}

func (s *ListPipelinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
