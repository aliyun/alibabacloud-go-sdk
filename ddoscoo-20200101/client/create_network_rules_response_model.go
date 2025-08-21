// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkRulesResponseBody) *CreateNetworkRulesResponse
	GetBody() *CreateNetworkRulesResponseBody
}

type CreateNetworkRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkRulesResponse) GetBody() *CreateNetworkRulesResponseBody {
	return s.Body
}

func (s *CreateNetworkRulesResponse) SetHeaders(v map[string]*string) *CreateNetworkRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkRulesResponse) SetStatusCode(v int32) *CreateNetworkRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkRulesResponse) SetBody(v *CreateNetworkRulesResponseBody) *CreateNetworkRulesResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkRulesResponse) Validate() error {
	return dara.Validate(s)
}
