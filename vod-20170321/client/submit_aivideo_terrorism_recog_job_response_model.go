// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoTerrorismRecogJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIVideoTerrorismRecogJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIVideoTerrorismRecogJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIVideoTerrorismRecogJobResponseBody) *SubmitAIVideoTerrorismRecogJobResponse
	GetBody() *SubmitAIVideoTerrorismRecogJobResponseBody
}

type SubmitAIVideoTerrorismRecogJobResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIVideoTerrorismRecogJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIVideoTerrorismRecogJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoTerrorismRecogJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoTerrorismRecogJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIVideoTerrorismRecogJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIVideoTerrorismRecogJobResponse) GetBody() *SubmitAIVideoTerrorismRecogJobResponseBody {
	return s.Body
}

func (s *SubmitAIVideoTerrorismRecogJobResponse) SetHeaders(v map[string]*string) *SubmitAIVideoTerrorismRecogJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponse) SetStatusCode(v int32) *SubmitAIVideoTerrorismRecogJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponse) SetBody(v *SubmitAIVideoTerrorismRecogJobResponseBody) *SubmitAIVideoTerrorismRecogJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponse) Validate() error {
	return dara.Validate(s)
}
