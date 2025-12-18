// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceGroupByRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceComplianceGroupByRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceComplianceGroupByRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceComplianceGroupByRegionResponseBody) *GetAggregateResourceComplianceGroupByRegionResponse
	GetBody() *GetAggregateResourceComplianceGroupByRegionResponseBody
}

type GetAggregateResourceComplianceGroupByRegionResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceComplianceGroupByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) GetBody() *GetAggregateResourceComplianceGroupByRegionResponseBody {
	return s.Body
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceGroupByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceGroupByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) SetBody(v *GetAggregateResourceComplianceGroupByRegionResponseBody) *GetAggregateResourceComplianceGroupByRegionResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
