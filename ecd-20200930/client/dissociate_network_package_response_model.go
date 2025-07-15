// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateNetworkPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateNetworkPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateNetworkPackageResponse
	GetStatusCode() *int32
	SetBody(v *DissociateNetworkPackageResponseBody) *DissociateNetworkPackageResponse
	GetBody() *DissociateNetworkPackageResponseBody
}

type DissociateNetworkPackageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateNetworkPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateNetworkPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *DissociateNetworkPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateNetworkPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateNetworkPackageResponse) GetBody() *DissociateNetworkPackageResponseBody {
	return s.Body
}

func (s *DissociateNetworkPackageResponse) SetHeaders(v map[string]*string) *DissociateNetworkPackageResponse {
	s.Headers = v
	return s
}

func (s *DissociateNetworkPackageResponse) SetStatusCode(v int32) *DissociateNetworkPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateNetworkPackageResponse) SetBody(v *DissociateNetworkPackageResponseBody) *DissociateNetworkPackageResponse {
	s.Body = v
	return s
}

func (s *DissociateNetworkPackageResponse) Validate() error {
	return dara.Validate(s)
}
