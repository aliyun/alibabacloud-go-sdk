// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPackageResponse
	GetStatusCode() *int32
	SetBody(v *GetPackageResponseBody) *GetPackageResponse
	GetBody() *GetPackageResponseBody
}

type GetPackageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPackageResponse) GoString() string {
	return s.String()
}

func (s *GetPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPackageResponse) GetBody() *GetPackageResponseBody {
	return s.Body
}

func (s *GetPackageResponse) SetHeaders(v map[string]*string) *GetPackageResponse {
	s.Headers = v
	return s
}

func (s *GetPackageResponse) SetStatusCode(v int32) *GetPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPackageResponse) SetBody(v *GetPackageResponseBody) *GetPackageResponse {
	s.Body = v
	return s
}

func (s *GetPackageResponse) Validate() error {
	return dara.Validate(s)
}
