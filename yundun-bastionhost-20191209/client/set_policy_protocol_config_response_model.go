// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyProtocolConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPolicyProtocolConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPolicyProtocolConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetPolicyProtocolConfigResponseBody) *SetPolicyProtocolConfigResponse
	GetBody() *SetPolicyProtocolConfigResponseBody
}

type SetPolicyProtocolConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPolicyProtocolConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPolicyProtocolConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyProtocolConfigResponse) GoString() string {
	return s.String()
}

func (s *SetPolicyProtocolConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPolicyProtocolConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPolicyProtocolConfigResponse) GetBody() *SetPolicyProtocolConfigResponseBody {
	return s.Body
}

func (s *SetPolicyProtocolConfigResponse) SetHeaders(v map[string]*string) *SetPolicyProtocolConfigResponse {
	s.Headers = v
	return s
}

func (s *SetPolicyProtocolConfigResponse) SetStatusCode(v int32) *SetPolicyProtocolConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPolicyProtocolConfigResponse) SetBody(v *SetPolicyProtocolConfigResponseBody) *SetPolicyProtocolConfigResponse {
	s.Body = v
	return s
}

func (s *SetPolicyProtocolConfigResponse) Validate() error {
	return dara.Validate(s)
}
