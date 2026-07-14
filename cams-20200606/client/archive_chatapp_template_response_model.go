// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveChatappTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ArchiveChatappTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ArchiveChatappTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ArchiveChatappTemplateResponseBody) *ArchiveChatappTemplateResponse
	GetBody() *ArchiveChatappTemplateResponseBody
}

type ArchiveChatappTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ArchiveChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ArchiveChatappTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ArchiveChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *ArchiveChatappTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ArchiveChatappTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ArchiveChatappTemplateResponse) GetBody() *ArchiveChatappTemplateResponseBody {
	return s.Body
}

func (s *ArchiveChatappTemplateResponse) SetHeaders(v map[string]*string) *ArchiveChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *ArchiveChatappTemplateResponse) SetStatusCode(v int32) *ArchiveChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ArchiveChatappTemplateResponse) SetBody(v *ArchiveChatappTemplateResponseBody) *ArchiveChatappTemplateResponse {
	s.Body = v
	return s
}

func (s *ArchiveChatappTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
