// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowVersionsResponseBody) *ListFlowVersionsResponse
	GetBody() *ListFlowVersionsResponseBody
}

type ListFlowVersionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowVersionsResponse) GetBody() *ListFlowVersionsResponseBody {
	return s.Body
}

func (s *ListFlowVersionsResponse) SetHeaders(v map[string]*string) *ListFlowVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowVersionsResponse) SetStatusCode(v int32) *ListFlowVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowVersionsResponse) SetBody(v *ListFlowVersionsResponseBody) *ListFlowVersionsResponse {
	s.Body = v
	return s
}

func (s *ListFlowVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
