// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateNetworkPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateNetworkPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateNetworkPackageResponse
	GetStatusCode() *int32
	SetBody(v *AssociateNetworkPackageResponseBody) *AssociateNetworkPackageResponse
	GetBody() *AssociateNetworkPackageResponseBody
}

type AssociateNetworkPackageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateNetworkPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateNetworkPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *AssociateNetworkPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateNetworkPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateNetworkPackageResponse) GetBody() *AssociateNetworkPackageResponseBody {
	return s.Body
}

func (s *AssociateNetworkPackageResponse) SetHeaders(v map[string]*string) *AssociateNetworkPackageResponse {
	s.Headers = v
	return s
}

func (s *AssociateNetworkPackageResponse) SetStatusCode(v int32) *AssociateNetworkPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateNetworkPackageResponse) SetBody(v *AssociateNetworkPackageResponseBody) *AssociateNetworkPackageResponse {
	s.Body = v
	return s
}

func (s *AssociateNetworkPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
