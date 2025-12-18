// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAggregateRemediationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAggregateRemediationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAggregateRemediationResponse
	GetStatusCode() *int32
	SetBody(v *StartAggregateRemediationResponseBody) *StartAggregateRemediationResponse
	GetBody() *StartAggregateRemediationResponseBody
}

type StartAggregateRemediationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAggregateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAggregateRemediationResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAggregateRemediationResponse) GoString() string {
	return s.String()
}

func (s *StartAggregateRemediationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAggregateRemediationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAggregateRemediationResponse) GetBody() *StartAggregateRemediationResponseBody {
	return s.Body
}

func (s *StartAggregateRemediationResponse) SetHeaders(v map[string]*string) *StartAggregateRemediationResponse {
	s.Headers = v
	return s
}

func (s *StartAggregateRemediationResponse) SetStatusCode(v int32) *StartAggregateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAggregateRemediationResponse) SetBody(v *StartAggregateRemediationResponseBody) *StartAggregateRemediationResponse {
	s.Body = v
	return s
}

func (s *StartAggregateRemediationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
