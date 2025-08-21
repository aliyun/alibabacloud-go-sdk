// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAclEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkAclEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkAclEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkAclEntryResponseBody) *DeleteNetworkAclEntryResponse
	GetBody() *DeleteNetworkAclEntryResponseBody
}

type DeleteNetworkAclEntryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkAclEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkAclEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAclEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAclEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkAclEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkAclEntryResponse) GetBody() *DeleteNetworkAclEntryResponseBody {
	return s.Body
}

func (s *DeleteNetworkAclEntryResponse) SetHeaders(v map[string]*string) *DeleteNetworkAclEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkAclEntryResponse) SetStatusCode(v int32) *DeleteNetworkAclEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkAclEntryResponse) SetBody(v *DeleteNetworkAclEntryResponseBody) *DeleteNetworkAclEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkAclEntryResponse) Validate() error {
	return dara.Validate(s)
}
