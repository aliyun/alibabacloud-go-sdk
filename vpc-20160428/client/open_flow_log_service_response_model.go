// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenFlowLogServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenFlowLogServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenFlowLogServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenFlowLogServiceResponseBody) *OpenFlowLogServiceResponse
	GetBody() *OpenFlowLogServiceResponseBody
}

type OpenFlowLogServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenFlowLogServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenFlowLogServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenFlowLogServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenFlowLogServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenFlowLogServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenFlowLogServiceResponse) GetBody() *OpenFlowLogServiceResponseBody {
	return s.Body
}

func (s *OpenFlowLogServiceResponse) SetHeaders(v map[string]*string) *OpenFlowLogServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenFlowLogServiceResponse) SetStatusCode(v int32) *OpenFlowLogServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenFlowLogServiceResponse) SetBody(v *OpenFlowLogServiceResponseBody) *OpenFlowLogServiceResponse {
	s.Body = v
	return s
}

func (s *OpenFlowLogServiceResponse) Validate() error {
	return dara.Validate(s)
}
