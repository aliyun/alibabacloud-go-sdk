// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceVideoTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceVideoTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceVideoTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceVideoTemplateResponseBody) *DeleteFaceVideoTemplateResponse
	GetBody() *DeleteFaceVideoTemplateResponseBody
}

type DeleteFaceVideoTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceVideoTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceVideoTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceVideoTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceVideoTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceVideoTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceVideoTemplateResponse) GetBody() *DeleteFaceVideoTemplateResponseBody {
	return s.Body
}

func (s *DeleteFaceVideoTemplateResponse) SetHeaders(v map[string]*string) *DeleteFaceVideoTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceVideoTemplateResponse) SetStatusCode(v int32) *DeleteFaceVideoTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceVideoTemplateResponse) SetBody(v *DeleteFaceVideoTemplateResponseBody) *DeleteFaceVideoTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceVideoTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
