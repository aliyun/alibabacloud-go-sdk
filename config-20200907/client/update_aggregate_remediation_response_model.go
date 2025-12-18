// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateRemediationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAggregateRemediationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAggregateRemediationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAggregateRemediationResponseBody) *UpdateAggregateRemediationResponse
	GetBody() *UpdateAggregateRemediationResponseBody
}

type UpdateAggregateRemediationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAggregateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAggregateRemediationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateRemediationResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregateRemediationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAggregateRemediationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAggregateRemediationResponse) GetBody() *UpdateAggregateRemediationResponseBody {
	return s.Body
}

func (s *UpdateAggregateRemediationResponse) SetHeaders(v map[string]*string) *UpdateAggregateRemediationResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregateRemediationResponse) SetStatusCode(v int32) *UpdateAggregateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregateRemediationResponse) SetBody(v *UpdateAggregateRemediationResponseBody) *UpdateAggregateRemediationResponse {
	s.Body = v
	return s
}

func (s *UpdateAggregateRemediationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
