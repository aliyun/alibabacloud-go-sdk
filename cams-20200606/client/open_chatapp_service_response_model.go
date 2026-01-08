// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenChatappServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenChatappServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenChatappServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenChatappServiceResponseBody) *OpenChatappServiceResponse
	GetBody() *OpenChatappServiceResponseBody
}

type OpenChatappServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenChatappServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenChatappServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenChatappServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenChatappServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenChatappServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenChatappServiceResponse) GetBody() *OpenChatappServiceResponseBody {
	return s.Body
}

func (s *OpenChatappServiceResponse) SetHeaders(v map[string]*string) *OpenChatappServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenChatappServiceResponse) SetStatusCode(v int32) *OpenChatappServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenChatappServiceResponse) SetBody(v *OpenChatappServiceResponseBody) *OpenChatappServiceResponse {
	s.Body = v
	return s
}

func (s *OpenChatappServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
