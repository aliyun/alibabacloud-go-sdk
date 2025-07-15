// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkPackagesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkPackagesResponseBody) *DeleteNetworkPackagesResponse
	GetBody() *DeleteNetworkPackagesResponseBody
}

type DeleteNetworkPackagesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkPackagesResponse) GetBody() *DeleteNetworkPackagesResponseBody {
	return s.Body
}

func (s *DeleteNetworkPackagesResponse) SetHeaders(v map[string]*string) *DeleteNetworkPackagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkPackagesResponse) SetStatusCode(v int32) *DeleteNetworkPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkPackagesResponse) SetBody(v *DeleteNetworkPackagesResponseBody) *DeleteNetworkPackagesResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkPackagesResponse) Validate() error {
	return dara.Validate(s)
}
