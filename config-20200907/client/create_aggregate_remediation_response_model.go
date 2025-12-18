// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateRemediationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAggregateRemediationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAggregateRemediationResponse
	GetStatusCode() *int32
	SetBody(v *CreateAggregateRemediationResponseBody) *CreateAggregateRemediationResponse
	GetBody() *CreateAggregateRemediationResponseBody
}

type CreateAggregateRemediationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAggregateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAggregateRemediationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateRemediationResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateRemediationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAggregateRemediationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAggregateRemediationResponse) GetBody() *CreateAggregateRemediationResponseBody {
	return s.Body
}

func (s *CreateAggregateRemediationResponse) SetHeaders(v map[string]*string) *CreateAggregateRemediationResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateRemediationResponse) SetStatusCode(v int32) *CreateAggregateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateRemediationResponse) SetBody(v *CreateAggregateRemediationResponseBody) *CreateAggregateRemediationResponse {
	s.Body = v
	return s
}

func (s *CreateAggregateRemediationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
