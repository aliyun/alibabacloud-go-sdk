// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaConnectFlowInputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMediaConnectFlowInputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMediaConnectFlowInputResponse
	GetStatusCode() *int32
	SetBody(v *AddMediaConnectFlowInputResponseBody) *AddMediaConnectFlowInputResponse
	GetBody() *AddMediaConnectFlowInputResponseBody
}

type AddMediaConnectFlowInputResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMediaConnectFlowInputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMediaConnectFlowInputResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMediaConnectFlowInputResponse) GoString() string {
	return s.String()
}

func (s *AddMediaConnectFlowInputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMediaConnectFlowInputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMediaConnectFlowInputResponse) GetBody() *AddMediaConnectFlowInputResponseBody {
	return s.Body
}

func (s *AddMediaConnectFlowInputResponse) SetHeaders(v map[string]*string) *AddMediaConnectFlowInputResponse {
	s.Headers = v
	return s
}

func (s *AddMediaConnectFlowInputResponse) SetStatusCode(v int32) *AddMediaConnectFlowInputResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMediaConnectFlowInputResponse) SetBody(v *AddMediaConnectFlowInputResponseBody) *AddMediaConnectFlowInputResponse {
	s.Body = v
	return s
}

func (s *AddMediaConnectFlowInputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
