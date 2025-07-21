// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlowResponse
	GetStatusCode() *int32
	SetBody(v *GetFlowResponseBody) *GetFlowResponse
	GetBody() *GetFlowResponseBody
}

type GetFlowResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlowResponse) GoString() string {
	return s.String()
}

func (s *GetFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlowResponse) GetBody() *GetFlowResponseBody {
	return s.Body
}

func (s *GetFlowResponse) SetHeaders(v map[string]*string) *GetFlowResponse {
	s.Headers = v
	return s
}

func (s *GetFlowResponse) SetStatusCode(v int32) *GetFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowResponse) SetBody(v *GetFlowResponseBody) *GetFlowResponse {
	s.Body = v
	return s
}

func (s *GetFlowResponse) Validate() error {
	return dara.Validate(s)
}
