// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAclEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkAclEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkAclEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkAclEntryResponseBody) *CreateNetworkAclEntryResponse
	GetBody() *CreateNetworkAclEntryResponseBody
}

type CreateNetworkAclEntryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkAclEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkAclEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkAclEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkAclEntryResponse) GetBody() *CreateNetworkAclEntryResponseBody {
	return s.Body
}

func (s *CreateNetworkAclEntryResponse) SetHeaders(v map[string]*string) *CreateNetworkAclEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkAclEntryResponse) SetStatusCode(v int32) *CreateNetworkAclEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkAclEntryResponse) SetBody(v *CreateNetworkAclEntryResponseBody) *CreateNetworkAclEntryResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkAclEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
