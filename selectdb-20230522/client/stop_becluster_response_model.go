// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBEClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopBEClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopBEClusterResponse
	GetStatusCode() *int32
	SetBody(v *StopBEClusterResponseBody) *StopBEClusterResponse
	GetBody() *StopBEClusterResponseBody
}

type StopBEClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopBEClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopBEClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StopBEClusterResponse) GoString() string {
	return s.String()
}

func (s *StopBEClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopBEClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopBEClusterResponse) GetBody() *StopBEClusterResponseBody {
	return s.Body
}

func (s *StopBEClusterResponse) SetHeaders(v map[string]*string) *StopBEClusterResponse {
	s.Headers = v
	return s
}

func (s *StopBEClusterResponse) SetStatusCode(v int32) *StopBEClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopBEClusterResponse) SetBody(v *StopBEClusterResponseBody) *StopBEClusterResponse {
	s.Body = v
	return s
}

func (s *StopBEClusterResponse) Validate() error {
	return dara.Validate(s)
}
