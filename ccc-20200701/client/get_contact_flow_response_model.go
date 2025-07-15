// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContactFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContactFlowResponse
	GetStatusCode() *int32
	SetBody(v *GetContactFlowResponseBody) *GetContactFlowResponse
	GetBody() *GetContactFlowResponseBody
}

type GetContactFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContactFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContactFlowResponse) GoString() string {
	return s.String()
}

func (s *GetContactFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContactFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContactFlowResponse) GetBody() *GetContactFlowResponseBody {
	return s.Body
}

func (s *GetContactFlowResponse) SetHeaders(v map[string]*string) *GetContactFlowResponse {
	s.Headers = v
	return s
}

func (s *GetContactFlowResponse) SetStatusCode(v int32) *GetContactFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContactFlowResponse) SetBody(v *GetContactFlowResponseBody) *GetContactFlowResponse {
	s.Body = v
	return s
}

func (s *GetContactFlowResponse) Validate() error {
	return dara.Validate(s)
}
