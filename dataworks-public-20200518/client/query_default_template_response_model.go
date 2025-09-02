// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDefaultTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDefaultTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDefaultTemplateResponse
	GetStatusCode() *int32
	SetBody(v *QueryDefaultTemplateResponseBody) *QueryDefaultTemplateResponse
	GetBody() *QueryDefaultTemplateResponseBody
}

type QueryDefaultTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDefaultTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDefaultTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDefaultTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryDefaultTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDefaultTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDefaultTemplateResponse) GetBody() *QueryDefaultTemplateResponseBody {
	return s.Body
}

func (s *QueryDefaultTemplateResponse) SetHeaders(v map[string]*string) *QueryDefaultTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryDefaultTemplateResponse) SetStatusCode(v int32) *QueryDefaultTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDefaultTemplateResponse) SetBody(v *QueryDefaultTemplateResponseBody) *QueryDefaultTemplateResponse {
	s.Body = v
	return s
}

func (s *QueryDefaultTemplateResponse) Validate() error {
	return dara.Validate(s)
}
