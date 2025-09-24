// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourcePackageResponseBody) *CreateResourcePackageResponse
	GetBody() *CreateResourcePackageResponseBody
}

type CreateResourcePackageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *CreateResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourcePackageResponse) GetBody() *CreateResourcePackageResponseBody {
	return s.Body
}

func (s *CreateResourcePackageResponse) SetHeaders(v map[string]*string) *CreateResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *CreateResourcePackageResponse) SetStatusCode(v int32) *CreateResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourcePackageResponse) SetBody(v *CreateResourcePackageResponseBody) *CreateResourcePackageResponse {
	s.Body = v
	return s
}

func (s *CreateResourcePackageResponse) Validate() error {
	return dara.Validate(s)
}
