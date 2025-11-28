// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageDomainRoutingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStorageDomainRoutingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStorageDomainRoutingRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateStorageDomainRoutingRuleResponseBody) *CreateStorageDomainRoutingRuleResponse
	GetBody() *CreateStorageDomainRoutingRuleResponseBody
}

type CreateStorageDomainRoutingRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStorageDomainRoutingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStorageDomainRoutingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageDomainRoutingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateStorageDomainRoutingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStorageDomainRoutingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStorageDomainRoutingRuleResponse) GetBody() *CreateStorageDomainRoutingRuleResponseBody {
	return s.Body
}

func (s *CreateStorageDomainRoutingRuleResponse) SetHeaders(v map[string]*string) *CreateStorageDomainRoutingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateStorageDomainRoutingRuleResponse) SetStatusCode(v int32) *CreateStorageDomainRoutingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStorageDomainRoutingRuleResponse) SetBody(v *CreateStorageDomainRoutingRuleResponseBody) *CreateStorageDomainRoutingRuleResponse {
	s.Body = v
	return s
}

func (s *CreateStorageDomainRoutingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
