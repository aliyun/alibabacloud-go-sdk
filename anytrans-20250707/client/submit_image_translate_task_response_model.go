// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitImageTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitImageTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitImageTranslateTaskResponseBody) *SubmitImageTranslateTaskResponse
	GetBody() *SubmitImageTranslateTaskResponseBody
}

type SubmitImageTranslateTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitImageTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitImageTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitImageTranslateTaskResponse) GetBody() *SubmitImageTranslateTaskResponseBody {
	return s.Body
}

func (s *SubmitImageTranslateTaskResponse) SetHeaders(v map[string]*string) *SubmitImageTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitImageTranslateTaskResponse) SetStatusCode(v int32) *SubmitImageTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImageTranslateTaskResponse) SetBody(v *SubmitImageTranslateTaskResponseBody) *SubmitImageTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitImageTranslateTaskResponse) Validate() error {
	return dara.Validate(s)
}
