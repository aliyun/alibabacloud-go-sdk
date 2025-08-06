// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoFaceRecogJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIVideoFaceRecogJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIVideoFaceRecogJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIVideoFaceRecogJobResponseBody) *SubmitAIVideoFaceRecogJobResponse
	GetBody() *SubmitAIVideoFaceRecogJobResponseBody
}

type SubmitAIVideoFaceRecogJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIVideoFaceRecogJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIVideoFaceRecogJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoFaceRecogJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoFaceRecogJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIVideoFaceRecogJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIVideoFaceRecogJobResponse) GetBody() *SubmitAIVideoFaceRecogJobResponseBody {
	return s.Body
}

func (s *SubmitAIVideoFaceRecogJobResponse) SetHeaders(v map[string]*string) *SubmitAIVideoFaceRecogJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponse) SetStatusCode(v int32) *SubmitAIVideoFaceRecogJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponse) SetBody(v *SubmitAIVideoFaceRecogJobResponseBody) *SubmitAIVideoFaceRecogJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponse) Validate() error {
	return dara.Validate(s)
}
