// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLongTextTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLongTextTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLongTextTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetLongTextTranslateTaskResponseBody) *GetLongTextTranslateTaskResponse
	GetBody() *GetLongTextTranslateTaskResponseBody
}

type GetLongTextTranslateTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLongTextTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLongTextTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLongTextTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *GetLongTextTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLongTextTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLongTextTranslateTaskResponse) GetBody() *GetLongTextTranslateTaskResponseBody {
	return s.Body
}

func (s *GetLongTextTranslateTaskResponse) SetHeaders(v map[string]*string) *GetLongTextTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *GetLongTextTranslateTaskResponse) SetStatusCode(v int32) *GetLongTextTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLongTextTranslateTaskResponse) SetBody(v *GetLongTextTranslateTaskResponseBody) *GetLongTextTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *GetLongTextTranslateTaskResponse) Validate() error {
	return dara.Validate(s)
}
