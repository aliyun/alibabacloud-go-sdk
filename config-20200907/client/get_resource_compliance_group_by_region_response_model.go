// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceGroupByRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceComplianceGroupByRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceComplianceGroupByRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceComplianceGroupByRegionResponseBody) *GetResourceComplianceGroupByRegionResponse
	GetBody() *GetResourceComplianceGroupByRegionResponseBody
}

type GetResourceComplianceGroupByRegionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceComplianceGroupByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceComplianceGroupByRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceComplianceGroupByRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceComplianceGroupByRegionResponse) GetBody() *GetResourceComplianceGroupByRegionResponseBody {
	return s.Body
}

func (s *GetResourceComplianceGroupByRegionResponse) SetHeaders(v map[string]*string) *GetResourceComplianceGroupByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponse) SetStatusCode(v int32) *GetResourceComplianceGroupByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponse) SetBody(v *GetResourceComplianceGroupByRegionResponseBody) *GetResourceComplianceGroupByRegionResponse {
	s.Body = v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
