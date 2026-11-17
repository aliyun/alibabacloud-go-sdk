// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommandInvocationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCommandInvocationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCommandInvocationResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListCommandInvocationResultsResponseBody) *ListCommandInvocationResultsResponse
	GetBody() *ListCommandInvocationResultsResponseBody
}

type ListCommandInvocationResultsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommandInvocationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommandInvocationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCommandInvocationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListCommandInvocationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCommandInvocationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCommandInvocationResultsResponse) GetBody() *ListCommandInvocationResultsResponseBody {
	return s.Body
}

func (s *ListCommandInvocationResultsResponse) SetHeaders(v map[string]*string) *ListCommandInvocationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListCommandInvocationResultsResponse) SetStatusCode(v int32) *ListCommandInvocationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommandInvocationResultsResponse) SetBody(v *ListCommandInvocationResultsResponseBody) *ListCommandInvocationResultsResponse {
	s.Body = v
	return s
}

func (s *ListCommandInvocationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
