// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAggregateCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAggregateCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAggregateCompliancePackResponseBody) *UpdateAggregateCompliancePackResponse
	GetBody() *UpdateAggregateCompliancePackResponseBody
}

type UpdateAggregateCompliancePackResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAggregateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAggregateCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAggregateCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAggregateCompliancePackResponse) GetBody() *UpdateAggregateCompliancePackResponseBody {
	return s.Body
}

func (s *UpdateAggregateCompliancePackResponse) SetHeaders(v map[string]*string) *UpdateAggregateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregateCompliancePackResponse) SetStatusCode(v int32) *UpdateAggregateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregateCompliancePackResponse) SetBody(v *UpdateAggregateCompliancePackResponseBody) *UpdateAggregateCompliancePackResponse {
	s.Body = v
	return s
}

func (s *UpdateAggregateCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
