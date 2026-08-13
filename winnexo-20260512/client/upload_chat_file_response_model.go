// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadChatFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadChatFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadChatFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadChatFileResponseBody) *UploadChatFileResponse
	GetBody() *UploadChatFileResponseBody
}

type UploadChatFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadChatFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadChatFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileResponse) GoString() string {
	return s.String()
}

func (s *UploadChatFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadChatFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadChatFileResponse) GetBody() *UploadChatFileResponseBody {
	return s.Body
}

func (s *UploadChatFileResponse) SetHeaders(v map[string]*string) *UploadChatFileResponse {
	s.Headers = v
	return s
}

func (s *UploadChatFileResponse) SetStatusCode(v int32) *UploadChatFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadChatFileResponse) SetBody(v *UploadChatFileResponseBody) *UploadChatFileResponse {
	s.Body = v
	return s
}

func (s *UploadChatFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
