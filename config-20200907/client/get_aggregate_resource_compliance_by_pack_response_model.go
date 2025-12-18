// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceByPackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceComplianceByPackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceComplianceByPackResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceComplianceByPackResponseBody) *GetAggregateResourceComplianceByPackResponse
	GetBody() *GetAggregateResourceComplianceByPackResponseBody
}

type GetAggregateResourceComplianceByPackResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceComplianceByPackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByPackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceComplianceByPackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceComplianceByPackResponse) GetBody() *GetAggregateResourceComplianceByPackResponseBody {
	return s.Body
}

func (s *GetAggregateResourceComplianceByPackResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponse) SetBody(v *GetAggregateResourceComplianceByPackResponseBody) *GetAggregateResourceComplianceByPackResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
