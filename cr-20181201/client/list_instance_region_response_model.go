// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceRegionResponseBody) *ListInstanceRegionResponse
	GetBody() *ListInstanceRegionResponseBody
}

type ListInstanceRegionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRegionResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceRegionResponse) GetBody() *ListInstanceRegionResponseBody {
	return s.Body
}

func (s *ListInstanceRegionResponse) SetHeaders(v map[string]*string) *ListInstanceRegionResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceRegionResponse) SetStatusCode(v int32) *ListInstanceRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceRegionResponse) SetBody(v *ListInstanceRegionResponseBody) *ListInstanceRegionResponse {
	s.Body = v
	return s
}

func (s *ListInstanceRegionResponse) Validate() error {
	return dara.Validate(s)
}
