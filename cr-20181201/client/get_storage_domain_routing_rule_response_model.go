// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageDomainRoutingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageDomainRoutingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageDomainRoutingRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageDomainRoutingRuleResponseBody) *GetStorageDomainRoutingRuleResponse
	GetBody() *GetStorageDomainRoutingRuleResponseBody
}

type GetStorageDomainRoutingRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageDomainRoutingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageDomainRoutingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageDomainRoutingRuleResponse) GoString() string {
	return s.String()
}

func (s *GetStorageDomainRoutingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageDomainRoutingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageDomainRoutingRuleResponse) GetBody() *GetStorageDomainRoutingRuleResponseBody {
	return s.Body
}

func (s *GetStorageDomainRoutingRuleResponse) SetHeaders(v map[string]*string) *GetStorageDomainRoutingRuleResponse {
	s.Headers = v
	return s
}

func (s *GetStorageDomainRoutingRuleResponse) SetStatusCode(v int32) *GetStorageDomainRoutingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageDomainRoutingRuleResponse) SetBody(v *GetStorageDomainRoutingRuleResponseBody) *GetStorageDomainRoutingRuleResponse {
	s.Body = v
	return s
}

func (s *GetStorageDomainRoutingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
