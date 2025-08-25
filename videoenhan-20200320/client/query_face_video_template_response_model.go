// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceVideoTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceVideoTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceVideoTemplateResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceVideoTemplateResponseBody) *QueryFaceVideoTemplateResponse
	GetBody() *QueryFaceVideoTemplateResponseBody
}

type QueryFaceVideoTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceVideoTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceVideoTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceVideoTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceVideoTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceVideoTemplateResponse) GetBody() *QueryFaceVideoTemplateResponseBody {
	return s.Body
}

func (s *QueryFaceVideoTemplateResponse) SetHeaders(v map[string]*string) *QueryFaceVideoTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceVideoTemplateResponse) SetStatusCode(v int32) *QueryFaceVideoTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceVideoTemplateResponse) SetBody(v *QueryFaceVideoTemplateResponseBody) *QueryFaceVideoTemplateResponse {
	s.Body = v
	return s
}

func (s *QueryFaceVideoTemplateResponse) Validate() error {
	return dara.Validate(s)
}
