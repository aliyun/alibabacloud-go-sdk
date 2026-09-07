// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ListNetBandwidthResponseBody) *ListNetBandwidthResponse
	GetBody() *ListNetBandwidthResponseBody
}

type ListNetBandwidthResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ListNetBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetBandwidthResponse) GetBody() *ListNetBandwidthResponseBody {
	return s.Body
}

func (s *ListNetBandwidthResponse) SetHeaders(v map[string]*string) *ListNetBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ListNetBandwidthResponse) SetStatusCode(v int32) *ListNetBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetBandwidthResponse) SetBody(v *ListNetBandwidthResponseBody) *ListNetBandwidthResponse {
	s.Body = v
	return s
}

func (s *ListNetBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
