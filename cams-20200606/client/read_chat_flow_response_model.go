// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadChatFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadChatFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadChatFlowResponse
	GetStatusCode() *int32
	SetBody(v *ReadChatFlowResponseBody) *ReadChatFlowResponse
	GetBody() *ReadChatFlowResponseBody
}

type ReadChatFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadChatFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadChatFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadChatFlowResponse) GoString() string {
	return s.String()
}

func (s *ReadChatFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadChatFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadChatFlowResponse) GetBody() *ReadChatFlowResponseBody {
	return s.Body
}

func (s *ReadChatFlowResponse) SetHeaders(v map[string]*string) *ReadChatFlowResponse {
	s.Headers = v
	return s
}

func (s *ReadChatFlowResponse) SetStatusCode(v int32) *ReadChatFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadChatFlowResponse) SetBody(v *ReadChatFlowResponseBody) *ReadChatFlowResponse {
	s.Body = v
	return s
}

func (s *ReadChatFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
