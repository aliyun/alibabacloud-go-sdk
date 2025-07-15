// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAclEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetworkAclEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetworkAclEntriesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetworkAclEntriesResponseBody) *UpdateNetworkAclEntriesResponse
	GetBody() *UpdateNetworkAclEntriesResponseBody
}

type UpdateNetworkAclEntriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkAclEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkAclEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAclEntriesResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAclEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetworkAclEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetworkAclEntriesResponse) GetBody() *UpdateNetworkAclEntriesResponseBody {
	return s.Body
}

func (s *UpdateNetworkAclEntriesResponse) SetHeaders(v map[string]*string) *UpdateNetworkAclEntriesResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkAclEntriesResponse) SetStatusCode(v int32) *UpdateNetworkAclEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkAclEntriesResponse) SetBody(v *UpdateNetworkAclEntriesResponseBody) *UpdateNetworkAclEntriesResponse {
	s.Body = v
	return s
}

func (s *UpdateNetworkAclEntriesResponse) Validate() error {
	return dara.Validate(s)
}
