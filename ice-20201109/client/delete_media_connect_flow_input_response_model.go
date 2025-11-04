// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowInputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaConnectFlowInputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaConnectFlowInputResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaConnectFlowInputResponseBody) *DeleteMediaConnectFlowInputResponse
	GetBody() *DeleteMediaConnectFlowInputResponseBody
}

type DeleteMediaConnectFlowInputResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaConnectFlowInputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaConnectFlowInputResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowInputResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowInputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaConnectFlowInputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaConnectFlowInputResponse) GetBody() *DeleteMediaConnectFlowInputResponseBody {
	return s.Body
}

func (s *DeleteMediaConnectFlowInputResponse) SetHeaders(v map[string]*string) *DeleteMediaConnectFlowInputResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaConnectFlowInputResponse) SetStatusCode(v int32) *DeleteMediaConnectFlowInputResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaConnectFlowInputResponse) SetBody(v *DeleteMediaConnectFlowInputResponseBody) *DeleteMediaConnectFlowInputResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaConnectFlowInputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
