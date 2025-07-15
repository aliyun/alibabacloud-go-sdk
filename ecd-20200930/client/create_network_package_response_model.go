// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkPackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkPackageResponseBody) *CreateNetworkPackageResponse
	GetBody() *CreateNetworkPackageResponseBody
}

type CreateNetworkPackageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkPackageResponse) GetBody() *CreateNetworkPackageResponseBody {
	return s.Body
}

func (s *CreateNetworkPackageResponse) SetHeaders(v map[string]*string) *CreateNetworkPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkPackageResponse) SetStatusCode(v int32) *CreateNetworkPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkPackageResponse) SetBody(v *CreateNetworkPackageResponseBody) *CreateNetworkPackageResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkPackageResponse) Validate() error {
	return dara.Validate(s)
}
