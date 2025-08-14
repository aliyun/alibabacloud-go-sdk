// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaConnectFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMediaConnectFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMediaConnectFlowResponse
	GetStatusCode() *int32
	SetBody(v *CreateMediaConnectFlowResponseBody) *CreateMediaConnectFlowResponse
	GetBody() *CreateMediaConnectFlowResponseBody
}

type CreateMediaConnectFlowResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMediaConnectFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMediaConnectFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConnectFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaConnectFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMediaConnectFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMediaConnectFlowResponse) GetBody() *CreateMediaConnectFlowResponseBody {
	return s.Body
}

func (s *CreateMediaConnectFlowResponse) SetHeaders(v map[string]*string) *CreateMediaConnectFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaConnectFlowResponse) SetStatusCode(v int32) *CreateMediaConnectFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMediaConnectFlowResponse) SetBody(v *CreateMediaConnectFlowResponseBody) *CreateMediaConnectFlowResponse {
	s.Body = v
	return s
}

func (s *CreateMediaConnectFlowResponse) Validate() error {
	return dara.Validate(s)
}
