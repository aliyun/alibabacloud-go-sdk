// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaConnectFlowOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaConnectFlowOutputResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaConnectFlowOutputResponseBody) *UpdateMediaConnectFlowOutputResponse
	GetBody() *UpdateMediaConnectFlowOutputResponseBody
}

type UpdateMediaConnectFlowOutputResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaConnectFlowOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaConnectFlowOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowOutputResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaConnectFlowOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaConnectFlowOutputResponse) GetBody() *UpdateMediaConnectFlowOutputResponseBody {
	return s.Body
}

func (s *UpdateMediaConnectFlowOutputResponse) SetHeaders(v map[string]*string) *UpdateMediaConnectFlowOutputResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaConnectFlowOutputResponse) SetStatusCode(v int32) *UpdateMediaConnectFlowOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputResponse) SetBody(v *UpdateMediaConnectFlowOutputResponseBody) *UpdateMediaConnectFlowOutputResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaConnectFlowOutputResponse) Validate() error {
	return dara.Validate(s)
}
