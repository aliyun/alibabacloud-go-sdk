// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTryRunTaskFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TryRunTaskFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TryRunTaskFlowResponse
	GetStatusCode() *int32
	SetBody(v *TryRunTaskFlowResponseBody) *TryRunTaskFlowResponse
	GetBody() *TryRunTaskFlowResponseBody
}

type TryRunTaskFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TryRunTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TryRunTaskFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s TryRunTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *TryRunTaskFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TryRunTaskFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TryRunTaskFlowResponse) GetBody() *TryRunTaskFlowResponseBody {
	return s.Body
}

func (s *TryRunTaskFlowResponse) SetHeaders(v map[string]*string) *TryRunTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *TryRunTaskFlowResponse) SetStatusCode(v int32) *TryRunTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *TryRunTaskFlowResponse) SetBody(v *TryRunTaskFlowResponseBody) *TryRunTaskFlowResponse {
	s.Body = v
	return s
}

func (s *TryRunTaskFlowResponse) Validate() error {
	return dara.Validate(s)
}
