// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushTemplateResponse
	GetStatusCode() *int32
	SetBody(v *PushTemplateResponseBody) *PushTemplateResponse
	GetBody() *PushTemplateResponseBody
}

type PushTemplateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s PushTemplateResponse) GoString() string {
	return s.String()
}

func (s *PushTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushTemplateResponse) GetBody() *PushTemplateResponseBody {
	return s.Body
}

func (s *PushTemplateResponse) SetHeaders(v map[string]*string) *PushTemplateResponse {
	s.Headers = v
	return s
}

func (s *PushTemplateResponse) SetStatusCode(v int32) *PushTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *PushTemplateResponse) SetBody(v *PushTemplateResponseBody) *PushTemplateResponse {
	s.Body = v
	return s
}

func (s *PushTemplateResponse) Validate() error {
	return dara.Validate(s)
}
