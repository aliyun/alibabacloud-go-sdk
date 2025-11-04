// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaConnectFlowOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaConnectFlowOutputResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaConnectFlowOutputResponseBody) *DeleteMediaConnectFlowOutputResponse
	GetBody() *DeleteMediaConnectFlowOutputResponseBody
}

type DeleteMediaConnectFlowOutputResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaConnectFlowOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaConnectFlowOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowOutputResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaConnectFlowOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaConnectFlowOutputResponse) GetBody() *DeleteMediaConnectFlowOutputResponseBody {
	return s.Body
}

func (s *DeleteMediaConnectFlowOutputResponse) SetHeaders(v map[string]*string) *DeleteMediaConnectFlowOutputResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaConnectFlowOutputResponse) SetStatusCode(v int32) *DeleteMediaConnectFlowOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaConnectFlowOutputResponse) SetBody(v *DeleteMediaConnectFlowOutputResponseBody) *DeleteMediaConnectFlowOutputResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaConnectFlowOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
