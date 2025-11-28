// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeDeviceResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeDeviceResponseBody) *ListEdgeDeviceResponse
	GetBody() *ListEdgeDeviceResponseBody
}

type ListEdgeDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeDeviceResponse) GetBody() *ListEdgeDeviceResponseBody {
	return s.Body
}

func (s *ListEdgeDeviceResponse) SetHeaders(v map[string]*string) *ListEdgeDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeDeviceResponse) SetStatusCode(v int32) *ListEdgeDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeDeviceResponse) SetBody(v *ListEdgeDeviceResponseBody) *ListEdgeDeviceResponse {
	s.Body = v
	return s
}

func (s *ListEdgeDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
