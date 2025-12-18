// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateCompliancePackResponseBody) *GetAggregateCompliancePackResponse
	GetBody() *GetAggregateCompliancePackResponseBody
}

type GetAggregateCompliancePackResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateCompliancePackResponse) GetBody() *GetAggregateCompliancePackResponseBody {
	return s.Body
}

func (s *GetAggregateCompliancePackResponse) SetHeaders(v map[string]*string) *GetAggregateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateCompliancePackResponse) SetStatusCode(v int32) *GetAggregateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateCompliancePackResponse) SetBody(v *GetAggregateCompliancePackResponseBody) *GetAggregateCompliancePackResponse {
	s.Body = v
	return s
}

func (s *GetAggregateCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
