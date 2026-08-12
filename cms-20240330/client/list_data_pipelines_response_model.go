// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataPipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataPipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataPipelinesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataPipelinesResponseBody) *ListDataPipelinesResponse
	GetBody() *ListDataPipelinesResponseBody
}

type ListDataPipelinesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataPipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataPipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataPipelinesResponse) GetBody() *ListDataPipelinesResponseBody {
	return s.Body
}

func (s *ListDataPipelinesResponse) SetHeaders(v map[string]*string) *ListDataPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListDataPipelinesResponse) SetStatusCode(v int32) *ListDataPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataPipelinesResponse) SetBody(v *ListDataPipelinesResponseBody) *ListDataPipelinesResponse {
	s.Body = v
	return s
}

func (s *ListDataPipelinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
