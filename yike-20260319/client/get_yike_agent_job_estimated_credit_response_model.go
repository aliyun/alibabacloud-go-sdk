// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAgentJobEstimatedCreditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeAgentJobEstimatedCreditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeAgentJobEstimatedCreditResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeAgentJobEstimatedCreditResponseBody) *GetYikeAgentJobEstimatedCreditResponse
	GetBody() *GetYikeAgentJobEstimatedCreditResponseBody
}

type GetYikeAgentJobEstimatedCreditResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeAgentJobEstimatedCreditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeAgentJobEstimatedCreditResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAgentJobEstimatedCreditResponse) GoString() string {
	return s.String()
}

func (s *GetYikeAgentJobEstimatedCreditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeAgentJobEstimatedCreditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeAgentJobEstimatedCreditResponse) GetBody() *GetYikeAgentJobEstimatedCreditResponseBody {
	return s.Body
}

func (s *GetYikeAgentJobEstimatedCreditResponse) SetHeaders(v map[string]*string) *GetYikeAgentJobEstimatedCreditResponse {
	s.Headers = v
	return s
}

func (s *GetYikeAgentJobEstimatedCreditResponse) SetStatusCode(v int32) *GetYikeAgentJobEstimatedCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeAgentJobEstimatedCreditResponse) SetBody(v *GetYikeAgentJobEstimatedCreditResponseBody) *GetYikeAgentJobEstimatedCreditResponse {
	s.Body = v
	return s
}

func (s *GetYikeAgentJobEstimatedCreditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
