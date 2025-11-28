// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStorageDomainRoutingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStorageDomainRoutingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStorageDomainRoutingRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStorageDomainRoutingRuleResponseBody) *UpdateStorageDomainRoutingRuleResponse
	GetBody() *UpdateStorageDomainRoutingRuleResponseBody
}

type UpdateStorageDomainRoutingRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStorageDomainRoutingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStorageDomainRoutingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStorageDomainRoutingRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateStorageDomainRoutingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStorageDomainRoutingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStorageDomainRoutingRuleResponse) GetBody() *UpdateStorageDomainRoutingRuleResponseBody {
	return s.Body
}

func (s *UpdateStorageDomainRoutingRuleResponse) SetHeaders(v map[string]*string) *UpdateStorageDomainRoutingRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateStorageDomainRoutingRuleResponse) SetStatusCode(v int32) *UpdateStorageDomainRoutingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleResponse) SetBody(v *UpdateStorageDomainRoutingRuleResponseBody) *UpdateStorageDomainRoutingRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateStorageDomainRoutingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
