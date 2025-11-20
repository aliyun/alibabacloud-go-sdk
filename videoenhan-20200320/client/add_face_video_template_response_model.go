// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceVideoTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceVideoTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceVideoTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceVideoTemplateResponseBody) *AddFaceVideoTemplateResponse
	GetBody() *AddFaceVideoTemplateResponseBody
}

type AddFaceVideoTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceVideoTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceVideoTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceVideoTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceVideoTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceVideoTemplateResponse) GetBody() *AddFaceVideoTemplateResponseBody {
	return s.Body
}

func (s *AddFaceVideoTemplateResponse) SetHeaders(v map[string]*string) *AddFaceVideoTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddFaceVideoTemplateResponse) SetStatusCode(v int32) *AddFaceVideoTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceVideoTemplateResponse) SetBody(v *AddFaceVideoTemplateResponseBody) *AddFaceVideoTemplateResponse {
	s.Body = v
	return s
}

func (s *AddFaceVideoTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
