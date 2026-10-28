// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewNotifyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewNotifyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewNotifyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *PreviewNotifyTemplateResponseBody) *PreviewNotifyTemplateResponse
	GetBody() *PreviewNotifyTemplateResponseBody
}

type PreviewNotifyTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewNotifyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewNotifyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewNotifyTemplateResponse) GoString() string {
	return s.String()
}

func (s *PreviewNotifyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewNotifyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewNotifyTemplateResponse) GetBody() *PreviewNotifyTemplateResponseBody {
	return s.Body
}

func (s *PreviewNotifyTemplateResponse) SetHeaders(v map[string]*string) *PreviewNotifyTemplateResponse {
	s.Headers = v
	return s
}

func (s *PreviewNotifyTemplateResponse) SetStatusCode(v int32) *PreviewNotifyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewNotifyTemplateResponse) SetBody(v *PreviewNotifyTemplateResponseBody) *PreviewNotifyTemplateResponse {
	s.Body = v
	return s
}

func (s *PreviewNotifyTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
