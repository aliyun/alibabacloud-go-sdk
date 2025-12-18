// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceGroupByResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceComplianceGroupByResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceComplianceGroupByResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceComplianceGroupByResourceTypeResponseBody) *GetResourceComplianceGroupByResourceTypeResponse
	GetBody() *GetResourceComplianceGroupByResourceTypeResponseBody
}

type GetResourceComplianceGroupByResourceTypeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceComplianceGroupByResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceComplianceGroupByResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) GetBody() *GetResourceComplianceGroupByResourceTypeResponseBody {
	return s.Body
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) SetHeaders(v map[string]*string) *GetResourceComplianceGroupByResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) SetStatusCode(v int32) *GetResourceComplianceGroupByResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) SetBody(v *GetResourceComplianceGroupByResourceTypeResponseBody) *GetResourceComplianceGroupByResourceTypeResponse {
	s.Body = v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
