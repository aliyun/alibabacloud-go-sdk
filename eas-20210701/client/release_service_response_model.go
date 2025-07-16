// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseServiceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseServiceResponseBody) *ReleaseServiceResponse
	GetBody() *ReleaseServiceResponseBody
}

type ReleaseServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseServiceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseServiceResponse) GetBody() *ReleaseServiceResponseBody {
	return s.Body
}

func (s *ReleaseServiceResponse) SetHeaders(v map[string]*string) *ReleaseServiceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseServiceResponse) SetStatusCode(v int32) *ReleaseServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseServiceResponse) SetBody(v *ReleaseServiceResponseBody) *ReleaseServiceResponse {
	s.Body = v
	return s
}

func (s *ReleaseServiceResponse) Validate() error {
	return dara.Validate(s)
}
