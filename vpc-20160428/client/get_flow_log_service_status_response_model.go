// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowLogServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlowLogServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlowLogServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetFlowLogServiceStatusResponseBody) *GetFlowLogServiceStatusResponse
	GetBody() *GetFlowLogServiceStatusResponseBody
}

type GetFlowLogServiceStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowLogServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowLogServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlowLogServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetFlowLogServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlowLogServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlowLogServiceStatusResponse) GetBody() *GetFlowLogServiceStatusResponseBody {
	return s.Body
}

func (s *GetFlowLogServiceStatusResponse) SetHeaders(v map[string]*string) *GetFlowLogServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetFlowLogServiceStatusResponse) SetStatusCode(v int32) *GetFlowLogServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowLogServiceStatusResponse) SetBody(v *GetFlowLogServiceStatusResponseBody) *GetFlowLogServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetFlowLogServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
