// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceImageTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceImageTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceImageTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceImageTemplateResponseBody) *AddFaceImageTemplateResponse
	GetBody() *AddFaceImageTemplateResponseBody
}

type AddFaceImageTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceImageTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceImageTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceImageTemplateResponse) GetBody() *AddFaceImageTemplateResponseBody {
	return s.Body
}

func (s *AddFaceImageTemplateResponse) SetHeaders(v map[string]*string) *AddFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddFaceImageTemplateResponse) SetStatusCode(v int32) *AddFaceImageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceImageTemplateResponse) SetBody(v *AddFaceImageTemplateResponseBody) *AddFaceImageTemplateResponse {
	s.Body = v
	return s
}

func (s *AddFaceImageTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
