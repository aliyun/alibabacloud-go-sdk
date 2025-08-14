// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowInputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaConnectFlowInputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaConnectFlowInputResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaConnectFlowInputResponseBody) *UpdateMediaConnectFlowInputResponse
	GetBody() *UpdateMediaConnectFlowInputResponseBody
}

type UpdateMediaConnectFlowInputResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaConnectFlowInputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaConnectFlowInputResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowInputResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowInputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaConnectFlowInputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaConnectFlowInputResponse) GetBody() *UpdateMediaConnectFlowInputResponseBody {
	return s.Body
}

func (s *UpdateMediaConnectFlowInputResponse) SetHeaders(v map[string]*string) *UpdateMediaConnectFlowInputResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaConnectFlowInputResponse) SetStatusCode(v int32) *UpdateMediaConnectFlowInputResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaConnectFlowInputResponse) SetBody(v *UpdateMediaConnectFlowInputResponseBody) *UpdateMediaConnectFlowInputResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaConnectFlowInputResponse) Validate() error {
	return dara.Validate(s)
}
