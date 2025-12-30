// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIFileTemplateDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIFileTemplateDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIFileTemplateDownloadResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIFileTemplateDownloadResponseBody) *ChatBIFileTemplateDownloadResponse
	GetBody() *ChatBIFileTemplateDownloadResponseBody
}

type ChatBIFileTemplateDownloadResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIFileTemplateDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIFileTemplateDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIFileTemplateDownloadResponse) GoString() string {
	return s.String()
}

func (s *ChatBIFileTemplateDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIFileTemplateDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIFileTemplateDownloadResponse) GetBody() *ChatBIFileTemplateDownloadResponseBody {
	return s.Body
}

func (s *ChatBIFileTemplateDownloadResponse) SetHeaders(v map[string]*string) *ChatBIFileTemplateDownloadResponse {
	s.Headers = v
	return s
}

func (s *ChatBIFileTemplateDownloadResponse) SetStatusCode(v int32) *ChatBIFileTemplateDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIFileTemplateDownloadResponse) SetBody(v *ChatBIFileTemplateDownloadResponseBody) *ChatBIFileTemplateDownloadResponse {
	s.Body = v
	return s
}

func (s *ChatBIFileTemplateDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
