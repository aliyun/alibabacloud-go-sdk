// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaConnectFlowOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMediaConnectFlowOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMediaConnectFlowOutputResponse
	GetStatusCode() *int32
	SetBody(v *AddMediaConnectFlowOutputResponseBody) *AddMediaConnectFlowOutputResponse
	GetBody() *AddMediaConnectFlowOutputResponseBody
}

type AddMediaConnectFlowOutputResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMediaConnectFlowOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMediaConnectFlowOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMediaConnectFlowOutputResponse) GoString() string {
	return s.String()
}

func (s *AddMediaConnectFlowOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMediaConnectFlowOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMediaConnectFlowOutputResponse) GetBody() *AddMediaConnectFlowOutputResponseBody {
	return s.Body
}

func (s *AddMediaConnectFlowOutputResponse) SetHeaders(v map[string]*string) *AddMediaConnectFlowOutputResponse {
	s.Headers = v
	return s
}

func (s *AddMediaConnectFlowOutputResponse) SetStatusCode(v int32) *AddMediaConnectFlowOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMediaConnectFlowOutputResponse) SetBody(v *AddMediaConnectFlowOutputResponseBody) *AddMediaConnectFlowOutputResponse {
	s.Body = v
	return s
}

func (s *AddMediaConnectFlowOutputResponse) Validate() error {
	return dara.Validate(s)
}
