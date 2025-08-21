// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUnbindTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUnbindTemplateResponse
	GetStatusCode() *int32
	SetBody(v *BatchUnbindTemplateResponseBody) *BatchUnbindTemplateResponse
	GetBody() *BatchUnbindTemplateResponseBody
}

type BatchUnbindTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUnbindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUnbindTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindTemplateResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUnbindTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUnbindTemplateResponse) GetBody() *BatchUnbindTemplateResponseBody {
	return s.Body
}

func (s *BatchUnbindTemplateResponse) SetHeaders(v map[string]*string) *BatchUnbindTemplateResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindTemplateResponse) SetStatusCode(v int32) *BatchUnbindTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUnbindTemplateResponse) SetBody(v *BatchUnbindTemplateResponseBody) *BatchUnbindTemplateResponse {
	s.Body = v
	return s
}

func (s *BatchUnbindTemplateResponse) Validate() error {
	return dara.Validate(s)
}
