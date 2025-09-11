// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHtmlTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHtmlTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHtmlTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetHtmlTranslateTaskResponseBody) *GetHtmlTranslateTaskResponse
	GetBody() *GetHtmlTranslateTaskResponseBody
}

type GetHtmlTranslateTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHtmlTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHtmlTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHtmlTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *GetHtmlTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHtmlTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHtmlTranslateTaskResponse) GetBody() *GetHtmlTranslateTaskResponseBody {
	return s.Body
}

func (s *GetHtmlTranslateTaskResponse) SetHeaders(v map[string]*string) *GetHtmlTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *GetHtmlTranslateTaskResponse) SetStatusCode(v int32) *GetHtmlTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHtmlTranslateTaskResponse) SetBody(v *GetHtmlTranslateTaskResponseBody) *GetHtmlTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *GetHtmlTranslateTaskResponse) Validate() error {
	return dara.Validate(s)
}
