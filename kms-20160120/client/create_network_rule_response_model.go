// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkRuleResponseBody) *CreateNetworkRuleResponse
	GetBody() *CreateNetworkRuleResponseBody
}

type CreateNetworkRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkRuleResponse) GetBody() *CreateNetworkRuleResponseBody {
	return s.Body
}

func (s *CreateNetworkRuleResponse) SetHeaders(v map[string]*string) *CreateNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkRuleResponse) SetStatusCode(v int32) *CreateNetworkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkRuleResponse) SetBody(v *CreateNetworkRuleResponseBody) *CreateNetworkRuleResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkRuleResponse) Validate() error {
	return dara.Validate(s)
}
