// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccFlowInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVccFlowInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVccFlowInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListVccFlowInfosResponseBody) *ListVccFlowInfosResponse
	GetBody() *ListVccFlowInfosResponseBody
}

type ListVccFlowInfosResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVccFlowInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVccFlowInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVccFlowInfosResponse) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVccFlowInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVccFlowInfosResponse) GetBody() *ListVccFlowInfosResponseBody {
	return s.Body
}

func (s *ListVccFlowInfosResponse) SetHeaders(v map[string]*string) *ListVccFlowInfosResponse {
	s.Headers = v
	return s
}

func (s *ListVccFlowInfosResponse) SetStatusCode(v int32) *ListVccFlowInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVccFlowInfosResponse) SetBody(v *ListVccFlowInfosResponseBody) *ListVccFlowInfosResponse {
	s.Body = v
	return s
}

func (s *ListVccFlowInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
