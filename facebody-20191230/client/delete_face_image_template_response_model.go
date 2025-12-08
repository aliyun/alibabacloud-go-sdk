// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceImageTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceImageTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceImageTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceImageTemplateResponseBody) *DeleteFaceImageTemplateResponse
	GetBody() *DeleteFaceImageTemplateResponseBody
}

type DeleteFaceImageTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceImageTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceImageTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceImageTemplateResponse) GetBody() *DeleteFaceImageTemplateResponseBody {
	return s.Body
}

func (s *DeleteFaceImageTemplateResponse) SetHeaders(v map[string]*string) *DeleteFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceImageTemplateResponse) SetStatusCode(v int32) *DeleteFaceImageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceImageTemplateResponse) SetBody(v *DeleteFaceImageTemplateResponseBody) *DeleteFaceImageTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceImageTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
