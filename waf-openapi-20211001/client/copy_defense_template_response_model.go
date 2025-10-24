// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDefenseTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyDefenseTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyDefenseTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CopyDefenseTemplateResponseBody) *CopyDefenseTemplateResponse
	GetBody() *CopyDefenseTemplateResponseBody
}

type CopyDefenseTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDefenseTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *CopyDefenseTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyDefenseTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyDefenseTemplateResponse) GetBody() *CopyDefenseTemplateResponseBody {
	return s.Body
}

func (s *CopyDefenseTemplateResponse) SetHeaders(v map[string]*string) *CopyDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *CopyDefenseTemplateResponse) SetStatusCode(v int32) *CopyDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDefenseTemplateResponse) SetBody(v *CopyDefenseTemplateResponseBody) *CopyDefenseTemplateResponse {
	s.Body = v
	return s
}

func (s *CopyDefenseTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
