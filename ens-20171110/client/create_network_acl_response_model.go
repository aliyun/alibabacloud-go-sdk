// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkAclResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkAclResponseBody) *CreateNetworkAclResponse
	GetBody() *CreateNetworkAclResponseBody
}

type CreateNetworkAclResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkAclResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkAclResponse) GetBody() *CreateNetworkAclResponseBody {
	return s.Body
}

func (s *CreateNetworkAclResponse) SetHeaders(v map[string]*string) *CreateNetworkAclResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkAclResponse) SetStatusCode(v int32) *CreateNetworkAclResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkAclResponse) SetBody(v *CreateNetworkAclResponseBody) *CreateNetworkAclResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
