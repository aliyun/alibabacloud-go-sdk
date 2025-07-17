// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskFlowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskFlowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskFlowInstanceResponseBody) *ListTaskFlowInstanceResponse
	GetBody() *ListTaskFlowInstanceResponseBody
}

type ListTaskFlowInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskFlowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListTaskFlowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskFlowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskFlowInstanceResponse) GetBody() *ListTaskFlowInstanceResponseBody {
	return s.Body
}

func (s *ListTaskFlowInstanceResponse) SetHeaders(v map[string]*string) *ListTaskFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListTaskFlowInstanceResponse) SetStatusCode(v int32) *ListTaskFlowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskFlowInstanceResponse) SetBody(v *ListTaskFlowInstanceResponseBody) *ListTaskFlowInstanceResponse {
	s.Body = v
	return s
}

func (s *ListTaskFlowInstanceResponse) Validate() error {
	return dara.Validate(s)
}
