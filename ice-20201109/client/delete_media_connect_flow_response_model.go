// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaConnectFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaConnectFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaConnectFlowResponseBody) *DeleteMediaConnectFlowResponse
	GetBody() *DeleteMediaConnectFlowResponseBody
}

type DeleteMediaConnectFlowResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaConnectFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaConnectFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaConnectFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaConnectFlowResponse) GetBody() *DeleteMediaConnectFlowResponseBody {
	return s.Body
}

func (s *DeleteMediaConnectFlowResponse) SetHeaders(v map[string]*string) *DeleteMediaConnectFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaConnectFlowResponse) SetStatusCode(v int32) *DeleteMediaConnectFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaConnectFlowResponse) SetBody(v *DeleteMediaConnectFlowResponseBody) *DeleteMediaConnectFlowResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaConnectFlowResponse) Validate() error {
	return dara.Validate(s)
}
