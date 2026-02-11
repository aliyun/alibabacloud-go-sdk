// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutomateResponseConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutomateResponseConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutomateResponseConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListAutomateResponseConfigsResponseBody) *ListAutomateResponseConfigsResponse
	GetBody() *ListAutomateResponseConfigsResponseBody
}

type ListAutomateResponseConfigsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutomateResponseConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutomateResponseConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutomateResponseConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutomateResponseConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutomateResponseConfigsResponse) GetBody() *ListAutomateResponseConfigsResponseBody {
	return s.Body
}

func (s *ListAutomateResponseConfigsResponse) SetHeaders(v map[string]*string) *ListAutomateResponseConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAutomateResponseConfigsResponse) SetStatusCode(v int32) *ListAutomateResponseConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutomateResponseConfigsResponse) SetBody(v *ListAutomateResponseConfigsResponseBody) *ListAutomateResponseConfigsResponse {
	s.Body = v
	return s
}

func (s *ListAutomateResponseConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
