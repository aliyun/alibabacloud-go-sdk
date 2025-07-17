// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskFlowResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskFlowResponseBody) *ListTaskFlowResponse
	GetBody() *ListTaskFlowResponseBody
}

type ListTaskFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *ListTaskFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskFlowResponse) GetBody() *ListTaskFlowResponseBody {
	return s.Body
}

func (s *ListTaskFlowResponse) SetHeaders(v map[string]*string) *ListTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *ListTaskFlowResponse) SetStatusCode(v int32) *ListTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskFlowResponse) SetBody(v *ListTaskFlowResponseBody) *ListTaskFlowResponse {
	s.Body = v
	return s
}

func (s *ListTaskFlowResponse) Validate() error {
	return dara.Validate(s)
}
