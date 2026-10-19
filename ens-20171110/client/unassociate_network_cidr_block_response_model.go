// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateNetworkCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateNetworkCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateNetworkCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateNetworkCidrBlockResponseBody) *UnassociateNetworkCidrBlockResponse
	GetBody() *UnassociateNetworkCidrBlockResponseBody
}

type UnassociateNetworkCidrBlockResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateNetworkCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateNetworkCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateNetworkCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateNetworkCidrBlockResponse) GetBody() *UnassociateNetworkCidrBlockResponseBody {
	return s.Body
}

func (s *UnassociateNetworkCidrBlockResponse) SetHeaders(v map[string]*string) *UnassociateNetworkCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *UnassociateNetworkCidrBlockResponse) SetStatusCode(v int32) *UnassociateNetworkCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateNetworkCidrBlockResponse) SetBody(v *UnassociateNetworkCidrBlockResponseBody) *UnassociateNetworkCidrBlockResponse {
	s.Body = v
	return s
}

func (s *UnassociateNetworkCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
