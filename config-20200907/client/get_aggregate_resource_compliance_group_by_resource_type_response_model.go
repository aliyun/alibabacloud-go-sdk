// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceGroupByResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceComplianceGroupByResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceComplianceGroupByResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) *GetAggregateResourceComplianceGroupByResourceTypeResponse
	GetBody() *GetAggregateResourceComplianceGroupByResourceTypeResponseBody
}

type GetAggregateResourceComplianceGroupByResourceTypeResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceComplianceGroupByResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) GetBody() *GetAggregateResourceComplianceGroupByResourceTypeResponseBody {
	return s.Body
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceGroupByResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceGroupByResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) SetBody(v *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) *GetAggregateResourceComplianceGroupByResourceTypeResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
