// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoFaceRecogJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIVideoFaceRecogJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIVideoFaceRecogJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIVideoFaceRecogJobResponseBody) *ListAIVideoFaceRecogJobResponse
	GetBody() *ListAIVideoFaceRecogJobResponseBody
}

type ListAIVideoFaceRecogJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIVideoFaceRecogJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIVideoFaceRecogJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoFaceRecogJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIVideoFaceRecogJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIVideoFaceRecogJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIVideoFaceRecogJobResponse) GetBody() *ListAIVideoFaceRecogJobResponseBody {
	return s.Body
}

func (s *ListAIVideoFaceRecogJobResponse) SetHeaders(v map[string]*string) *ListAIVideoFaceRecogJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIVideoFaceRecogJobResponse) SetStatusCode(v int32) *ListAIVideoFaceRecogJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponse) SetBody(v *ListAIVideoFaceRecogJobResponseBody) *ListAIVideoFaceRecogJobResponse {
	s.Body = v
	return s
}

func (s *ListAIVideoFaceRecogJobResponse) Validate() error {
	return dara.Validate(s)
}
