// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAITrafficAnalysisStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAITrafficAnalysisStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAITrafficAnalysisStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAITrafficAnalysisStatusResponseBody) *UpdateAITrafficAnalysisStatusResponse
	GetBody() *UpdateAITrafficAnalysisStatusResponseBody
}

type UpdateAITrafficAnalysisStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAITrafficAnalysisStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAITrafficAnalysisStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAITrafficAnalysisStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAITrafficAnalysisStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAITrafficAnalysisStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAITrafficAnalysisStatusResponse) GetBody() *UpdateAITrafficAnalysisStatusResponseBody {
	return s.Body
}

func (s *UpdateAITrafficAnalysisStatusResponse) SetHeaders(v map[string]*string) *UpdateAITrafficAnalysisStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAITrafficAnalysisStatusResponse) SetStatusCode(v int32) *UpdateAITrafficAnalysisStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAITrafficAnalysisStatusResponse) SetBody(v *UpdateAITrafficAnalysisStatusResponseBody) *UpdateAITrafficAnalysisStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAITrafficAnalysisStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
