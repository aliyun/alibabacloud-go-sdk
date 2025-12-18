// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAggregateCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAggregateCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *CreateAggregateCompliancePackResponseBody) *CreateAggregateCompliancePackResponse
	GetBody() *CreateAggregateCompliancePackResponseBody
}

type CreateAggregateCompliancePackResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAggregateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAggregateCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAggregateCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAggregateCompliancePackResponse) GetBody() *CreateAggregateCompliancePackResponseBody {
	return s.Body
}

func (s *CreateAggregateCompliancePackResponse) SetHeaders(v map[string]*string) *CreateAggregateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateCompliancePackResponse) SetStatusCode(v int32) *CreateAggregateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateCompliancePackResponse) SetBody(v *CreateAggregateCompliancePackResponseBody) *CreateAggregateCompliancePackResponse {
	s.Body = v
	return s
}

func (s *CreateAggregateCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
