// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CopyTemplateResponseBody) *CopyTemplateResponse
	GetBody() *CopyTemplateResponseBody
}

type CopyTemplateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyTemplateResponse) GoString() string {
	return s.String()
}

func (s *CopyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyTemplateResponse) GetBody() *CopyTemplateResponseBody {
	return s.Body
}

func (s *CopyTemplateResponse) SetHeaders(v map[string]*string) *CopyTemplateResponse {
	s.Headers = v
	return s
}

func (s *CopyTemplateResponse) SetStatusCode(v int32) *CopyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyTemplateResponse) SetBody(v *CopyTemplateResponseBody) *CopyTemplateResponse {
	s.Body = v
	return s
}

func (s *CopyTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
