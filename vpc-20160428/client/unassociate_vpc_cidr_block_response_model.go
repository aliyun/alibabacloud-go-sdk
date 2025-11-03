// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateVpcCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateVpcCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateVpcCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateVpcCidrBlockResponseBody) *UnassociateVpcCidrBlockResponse
	GetBody() *UnassociateVpcCidrBlockResponseBody
}

type UnassociateVpcCidrBlockResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateVpcCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateVpcCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateVpcCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *UnassociateVpcCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateVpcCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateVpcCidrBlockResponse) GetBody() *UnassociateVpcCidrBlockResponseBody {
	return s.Body
}

func (s *UnassociateVpcCidrBlockResponse) SetHeaders(v map[string]*string) *UnassociateVpcCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *UnassociateVpcCidrBlockResponse) SetStatusCode(v int32) *UnassociateVpcCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateVpcCidrBlockResponse) SetBody(v *UnassociateVpcCidrBlockResponseBody) *UnassociateVpcCidrBlockResponse {
	s.Body = v
	return s
}

func (s *UnassociateVpcCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
