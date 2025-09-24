// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *RenewResourcePackageResponseBody) *RenewResourcePackageResponse
	GetBody() *RenewResourcePackageResponseBody
}

type RenewResourcePackageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *RenewResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewResourcePackageResponse) GetBody() *RenewResourcePackageResponseBody {
	return s.Body
}

func (s *RenewResourcePackageResponse) SetHeaders(v map[string]*string) *RenewResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *RenewResourcePackageResponse) SetStatusCode(v int32) *RenewResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewResourcePackageResponse) SetBody(v *RenewResourcePackageResponseBody) *RenewResourcePackageResponse {
	s.Body = v
	return s
}

func (s *RenewResourcePackageResponse) Validate() error {
	return dara.Validate(s)
}
