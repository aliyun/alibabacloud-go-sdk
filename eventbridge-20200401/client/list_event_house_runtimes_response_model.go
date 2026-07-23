// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventHouseRuntimesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventHouseRuntimesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventHouseRuntimesResponse
	GetStatusCode() *int32
	SetBody(v *ListEventHouseRuntimesResponseBody) *ListEventHouseRuntimesResponse
	GetBody() *ListEventHouseRuntimesResponseBody
}

type ListEventHouseRuntimesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventHouseRuntimesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventHouseRuntimesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventHouseRuntimesResponse) GoString() string {
	return s.String()
}

func (s *ListEventHouseRuntimesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventHouseRuntimesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventHouseRuntimesResponse) GetBody() *ListEventHouseRuntimesResponseBody {
	return s.Body
}

func (s *ListEventHouseRuntimesResponse) SetHeaders(v map[string]*string) *ListEventHouseRuntimesResponse {
	s.Headers = v
	return s
}

func (s *ListEventHouseRuntimesResponse) SetStatusCode(v int32) *ListEventHouseRuntimesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventHouseRuntimesResponse) SetBody(v *ListEventHouseRuntimesResponseBody) *ListEventHouseRuntimesResponse {
	s.Body = v
	return s
}

func (s *ListEventHouseRuntimesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
