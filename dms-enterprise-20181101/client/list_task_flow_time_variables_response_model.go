// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowTimeVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskFlowTimeVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskFlowTimeVariablesResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskFlowTimeVariablesResponseBody) *ListTaskFlowTimeVariablesResponse
	GetBody() *ListTaskFlowTimeVariablesResponseBody
}

type ListTaskFlowTimeVariablesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskFlowTimeVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskFlowTimeVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowTimeVariablesResponse) GoString() string {
	return s.String()
}

func (s *ListTaskFlowTimeVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskFlowTimeVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskFlowTimeVariablesResponse) GetBody() *ListTaskFlowTimeVariablesResponseBody {
	return s.Body
}

func (s *ListTaskFlowTimeVariablesResponse) SetHeaders(v map[string]*string) *ListTaskFlowTimeVariablesResponse {
	s.Headers = v
	return s
}

func (s *ListTaskFlowTimeVariablesResponse) SetStatusCode(v int32) *ListTaskFlowTimeVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskFlowTimeVariablesResponse) SetBody(v *ListTaskFlowTimeVariablesResponseBody) *ListTaskFlowTimeVariablesResponse {
	s.Body = v
	return s
}

func (s *ListTaskFlowTimeVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
