// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkAclResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkAclResponseBody) *DeleteNetworkAclResponse
	GetBody() *DeleteNetworkAclResponseBody
}

type DeleteNetworkAclResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkAclResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkAclResponse) GetBody() *DeleteNetworkAclResponseBody {
	return s.Body
}

func (s *DeleteNetworkAclResponse) SetHeaders(v map[string]*string) *DeleteNetworkAclResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkAclResponse) SetStatusCode(v int32) *DeleteNetworkAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkAclResponse) SetBody(v *DeleteNetworkAclResponseBody) *DeleteNetworkAclResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
