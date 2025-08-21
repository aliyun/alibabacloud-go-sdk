// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseRenderingDataPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseRenderingDataPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseRenderingDataPackageResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseRenderingDataPackageResponseBody) *ReleaseRenderingDataPackageResponse
	GetBody() *ReleaseRenderingDataPackageResponseBody
}

type ReleaseRenderingDataPackageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseRenderingDataPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseRenderingDataPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseRenderingDataPackageResponse) GoString() string {
	return s.String()
}

func (s *ReleaseRenderingDataPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseRenderingDataPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseRenderingDataPackageResponse) GetBody() *ReleaseRenderingDataPackageResponseBody {
	return s.Body
}

func (s *ReleaseRenderingDataPackageResponse) SetHeaders(v map[string]*string) *ReleaseRenderingDataPackageResponse {
	s.Headers = v
	return s
}

func (s *ReleaseRenderingDataPackageResponse) SetStatusCode(v int32) *ReleaseRenderingDataPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseRenderingDataPackageResponse) SetBody(v *ReleaseRenderingDataPackageResponseBody) *ReleaseRenderingDataPackageResponse {
	s.Body = v
	return s
}

func (s *ReleaseRenderingDataPackageResponse) Validate() error {
	return dara.Validate(s)
}
