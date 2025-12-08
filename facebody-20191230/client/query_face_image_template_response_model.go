// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceImageTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceImageTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceImageTemplateResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceImageTemplateResponseBody) *QueryFaceImageTemplateResponse
	GetBody() *QueryFaceImageTemplateResponseBody
}

type QueryFaceImageTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceImageTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceImageTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceImageTemplateResponse) GetBody() *QueryFaceImageTemplateResponseBody {
	return s.Body
}

func (s *QueryFaceImageTemplateResponse) SetHeaders(v map[string]*string) *QueryFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceImageTemplateResponse) SetStatusCode(v int32) *QueryFaceImageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceImageTemplateResponse) SetBody(v *QueryFaceImageTemplateResponseBody) *QueryFaceImageTemplateResponse {
	s.Body = v
	return s
}

func (s *QueryFaceImageTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
