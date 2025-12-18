// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateAccountComplianceByPackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateAccountComplianceByPackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateAccountComplianceByPackResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateAccountComplianceByPackResponseBody) *GetAggregateAccountComplianceByPackResponse
	GetBody() *GetAggregateAccountComplianceByPackResponseBody
}

type GetAggregateAccountComplianceByPackResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateAccountComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateAccountComplianceByPackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateAccountComplianceByPackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateAccountComplianceByPackResponse) GetBody() *GetAggregateAccountComplianceByPackResponseBody {
	return s.Body
}

func (s *GetAggregateAccountComplianceByPackResponse) SetHeaders(v map[string]*string) *GetAggregateAccountComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponse) SetStatusCode(v int32) *GetAggregateAccountComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponse) SetBody(v *GetAggregateAccountComplianceByPackResponseBody) *GetAggregateAccountComplianceByPackResponse {
	s.Body = v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
