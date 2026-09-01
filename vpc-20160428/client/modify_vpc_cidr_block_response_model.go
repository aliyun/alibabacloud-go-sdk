// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcCidrBlockResponseBody) *ModifyVpcCidrBlockResponse
	GetBody() *ModifyVpcCidrBlockResponseBody
}

type ModifyVpcCidrBlockResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcCidrBlockResponse) GetBody() *ModifyVpcCidrBlockResponseBody {
	return s.Body
}

func (s *ModifyVpcCidrBlockResponse) SetHeaders(v map[string]*string) *ModifyVpcCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcCidrBlockResponse) SetStatusCode(v int32) *ModifyVpcCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcCidrBlockResponse) SetBody(v *ModifyVpcCidrBlockResponseBody) *ModifyVpcCidrBlockResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
