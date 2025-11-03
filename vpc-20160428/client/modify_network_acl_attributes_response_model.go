// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkAclAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNetworkAclAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNetworkAclAttributesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNetworkAclAttributesResponseBody) *ModifyNetworkAclAttributesResponse
	GetBody() *ModifyNetworkAclAttributesResponseBody
}

type ModifyNetworkAclAttributesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNetworkAclAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNetworkAclAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkAclAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkAclAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNetworkAclAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNetworkAclAttributesResponse) GetBody() *ModifyNetworkAclAttributesResponseBody {
	return s.Body
}

func (s *ModifyNetworkAclAttributesResponse) SetHeaders(v map[string]*string) *ModifyNetworkAclAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkAclAttributesResponse) SetStatusCode(v int32) *ModifyNetworkAclAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNetworkAclAttributesResponse) SetBody(v *ModifyNetworkAclAttributesResponseBody) *ModifyNetworkAclAttributesResponse {
	s.Body = v
	return s
}

func (s *ModifyNetworkAclAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
