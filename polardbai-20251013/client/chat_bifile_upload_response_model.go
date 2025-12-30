// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIFileUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIFileUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIFileUploadResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIFileUploadResponseBody) *ChatBIFileUploadResponse
	GetBody() *ChatBIFileUploadResponseBody
}

type ChatBIFileUploadResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIFileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIFileUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIFileUploadResponse) GoString() string {
	return s.String()
}

func (s *ChatBIFileUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIFileUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIFileUploadResponse) GetBody() *ChatBIFileUploadResponseBody {
	return s.Body
}

func (s *ChatBIFileUploadResponse) SetHeaders(v map[string]*string) *ChatBIFileUploadResponse {
	s.Headers = v
	return s
}

func (s *ChatBIFileUploadResponse) SetStatusCode(v int32) *ChatBIFileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIFileUploadResponse) SetBody(v *ChatBIFileUploadResponseBody) *ChatBIFileUploadResponse {
	s.Body = v
	return s
}

func (s *ChatBIFileUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
