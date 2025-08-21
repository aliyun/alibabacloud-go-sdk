// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingDataPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRenderingDataPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRenderingDataPackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateRenderingDataPackageResponseBody) *CreateRenderingDataPackageResponse
	GetBody() *CreateRenderingDataPackageResponseBody
}

type CreateRenderingDataPackageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRenderingDataPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRenderingDataPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingDataPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateRenderingDataPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRenderingDataPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRenderingDataPackageResponse) GetBody() *CreateRenderingDataPackageResponseBody {
	return s.Body
}

func (s *CreateRenderingDataPackageResponse) SetHeaders(v map[string]*string) *CreateRenderingDataPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateRenderingDataPackageResponse) SetStatusCode(v int32) *CreateRenderingDataPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRenderingDataPackageResponse) SetBody(v *CreateRenderingDataPackageResponseBody) *CreateRenderingDataPackageResponse {
	s.Body = v
	return s
}

func (s *CreateRenderingDataPackageResponse) Validate() error {
	return dara.Validate(s)
}
