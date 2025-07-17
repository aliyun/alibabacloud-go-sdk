// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenVClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenVClusterResponse
	GetStatusCode() *int32
	SetBody(v *OpenVClusterResponseBody) *OpenVClusterResponse
	GetBody() *OpenVClusterResponseBody
}

type OpenVClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenVClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenVClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenVClusterResponse) GoString() string {
	return s.String()
}

func (s *OpenVClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenVClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenVClusterResponse) GetBody() *OpenVClusterResponseBody {
	return s.Body
}

func (s *OpenVClusterResponse) SetHeaders(v map[string]*string) *OpenVClusterResponse {
	s.Headers = v
	return s
}

func (s *OpenVClusterResponse) SetStatusCode(v int32) *OpenVClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenVClusterResponse) SetBody(v *OpenVClusterResponseBody) *OpenVClusterResponse {
	s.Body = v
	return s
}

func (s *OpenVClusterResponse) Validate() error {
	return dara.Validate(s)
}
