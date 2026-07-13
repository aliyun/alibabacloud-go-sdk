// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelInvocationSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelInvocationSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelInvocationSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetModelInvocationSummaryResponseBody) *GetModelInvocationSummaryResponse
	GetBody() *GetModelInvocationSummaryResponseBody
}

type GetModelInvocationSummaryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelInvocationSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelInvocationSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelInvocationSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetModelInvocationSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelInvocationSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelInvocationSummaryResponse) GetBody() *GetModelInvocationSummaryResponseBody {
	return s.Body
}

func (s *GetModelInvocationSummaryResponse) SetHeaders(v map[string]*string) *GetModelInvocationSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetModelInvocationSummaryResponse) SetStatusCode(v int32) *GetModelInvocationSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelInvocationSummaryResponse) SetBody(v *GetModelInvocationSummaryResponseBody) *GetModelInvocationSummaryResponse {
	s.Body = v
	return s
}

func (s *GetModelInvocationSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
