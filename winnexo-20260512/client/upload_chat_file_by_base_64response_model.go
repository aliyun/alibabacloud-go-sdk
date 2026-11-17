// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadChatFileByBase64Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadChatFileByBase64Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadChatFileByBase64Response
	GetStatusCode() *int32
	SetBody(v *UploadChatFileByBase64ResponseBody) *UploadChatFileByBase64Response
	GetBody() *UploadChatFileByBase64ResponseBody
}

type UploadChatFileByBase64Response struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadChatFileByBase64ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadChatFileByBase64Response) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileByBase64Response) GoString() string {
	return s.String()
}

func (s *UploadChatFileByBase64Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadChatFileByBase64Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadChatFileByBase64Response) GetBody() *UploadChatFileByBase64ResponseBody {
	return s.Body
}

func (s *UploadChatFileByBase64Response) SetHeaders(v map[string]*string) *UploadChatFileByBase64Response {
	s.Headers = v
	return s
}

func (s *UploadChatFileByBase64Response) SetStatusCode(v int32) *UploadChatFileByBase64Response {
	s.StatusCode = &v
	return s
}

func (s *UploadChatFileByBase64Response) SetBody(v *UploadChatFileByBase64ResponseBody) *UploadChatFileByBase64Response {
	s.Body = v
	return s
}

func (s *UploadChatFileByBase64Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
