// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManuallyStartDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManuallyStartDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManuallyStartDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *ManuallyStartDBClusterResponseBody) *ManuallyStartDBClusterResponse
	GetBody() *ManuallyStartDBClusterResponseBody
}

type ManuallyStartDBClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManuallyStartDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManuallyStartDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ManuallyStartDBClusterResponse) GoString() string {
	return s.String()
}

func (s *ManuallyStartDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManuallyStartDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManuallyStartDBClusterResponse) GetBody() *ManuallyStartDBClusterResponseBody {
	return s.Body
}

func (s *ManuallyStartDBClusterResponse) SetHeaders(v map[string]*string) *ManuallyStartDBClusterResponse {
	s.Headers = v
	return s
}

func (s *ManuallyStartDBClusterResponse) SetStatusCode(v int32) *ManuallyStartDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ManuallyStartDBClusterResponse) SetBody(v *ManuallyStartDBClusterResponseBody) *ManuallyStartDBClusterResponse {
	s.Body = v
	return s
}

func (s *ManuallyStartDBClusterResponse) Validate() error {
	return dara.Validate(s)
}
