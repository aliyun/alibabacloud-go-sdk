// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateNetworkCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateNetworkCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateNetworkCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *AssociateNetworkCidrBlockResponseBody) *AssociateNetworkCidrBlockResponse
	GetBody() *AssociateNetworkCidrBlockResponseBody
}

type AssociateNetworkCidrBlockResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateNetworkCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateNetworkCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *AssociateNetworkCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateNetworkCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateNetworkCidrBlockResponse) GetBody() *AssociateNetworkCidrBlockResponseBody {
	return s.Body
}

func (s *AssociateNetworkCidrBlockResponse) SetHeaders(v map[string]*string) *AssociateNetworkCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *AssociateNetworkCidrBlockResponse) SetStatusCode(v int32) *AssociateNetworkCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateNetworkCidrBlockResponse) SetBody(v *AssociateNetworkCidrBlockResponseBody) *AssociateNetworkCidrBlockResponse {
	s.Body = v
	return s
}

func (s *AssociateNetworkCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
