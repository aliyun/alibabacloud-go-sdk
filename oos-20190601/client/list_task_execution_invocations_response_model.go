// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskExecutionInvocationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskExecutionInvocationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskExecutionInvocationsResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskExecutionInvocationsResponseBody) *ListTaskExecutionInvocationsResponse
	GetBody() *ListTaskExecutionInvocationsResponseBody
}

type ListTaskExecutionInvocationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskExecutionInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskExecutionInvocationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskExecutionInvocationsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionInvocationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskExecutionInvocationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskExecutionInvocationsResponse) GetBody() *ListTaskExecutionInvocationsResponseBody {
	return s.Body
}

func (s *ListTaskExecutionInvocationsResponse) SetHeaders(v map[string]*string) *ListTaskExecutionInvocationsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskExecutionInvocationsResponse) SetStatusCode(v int32) *ListTaskExecutionInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskExecutionInvocationsResponse) SetBody(v *ListTaskExecutionInvocationsResponseBody) *ListTaskExecutionInvocationsResponse {
	s.Body = v
	return s
}

func (s *ListTaskExecutionInvocationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
