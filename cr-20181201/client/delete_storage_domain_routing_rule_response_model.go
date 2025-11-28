// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageDomainRoutingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStorageDomainRoutingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStorageDomainRoutingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStorageDomainRoutingRuleResponseBody) *DeleteStorageDomainRoutingRuleResponse
	GetBody() *DeleteStorageDomainRoutingRuleResponseBody
}

type DeleteStorageDomainRoutingRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStorageDomainRoutingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStorageDomainRoutingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageDomainRoutingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteStorageDomainRoutingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStorageDomainRoutingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStorageDomainRoutingRuleResponse) GetBody() *DeleteStorageDomainRoutingRuleResponseBody {
	return s.Body
}

func (s *DeleteStorageDomainRoutingRuleResponse) SetHeaders(v map[string]*string) *DeleteStorageDomainRoutingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteStorageDomainRoutingRuleResponse) SetStatusCode(v int32) *DeleteStorageDomainRoutingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStorageDomainRoutingRuleResponse) SetBody(v *DeleteStorageDomainRoutingRuleResponseBody) *DeleteStorageDomainRoutingRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteStorageDomainRoutingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
