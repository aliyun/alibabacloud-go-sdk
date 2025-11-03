// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpcCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateVpcCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateVpcCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *AssociateVpcCidrBlockResponseBody) *AssociateVpcCidrBlockResponse
	GetBody() *AssociateVpcCidrBlockResponseBody
}

type AssociateVpcCidrBlockResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateVpcCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateVpcCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpcCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *AssociateVpcCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateVpcCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateVpcCidrBlockResponse) GetBody() *AssociateVpcCidrBlockResponseBody {
	return s.Body
}

func (s *AssociateVpcCidrBlockResponse) SetHeaders(v map[string]*string) *AssociateVpcCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *AssociateVpcCidrBlockResponse) SetStatusCode(v int32) *AssociateVpcCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateVpcCidrBlockResponse) SetBody(v *AssociateVpcCidrBlockResponseBody) *AssociateVpcCidrBlockResponse {
	s.Body = v
	return s
}

func (s *AssociateVpcCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
