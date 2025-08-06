// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIASRJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIASRJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIASRJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIASRJobResponseBody) *SubmitAIASRJobResponse
	GetBody() *SubmitAIASRJobResponseBody
}

type SubmitAIASRJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIASRJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIASRJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIASRJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIASRJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIASRJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIASRJobResponse) GetBody() *SubmitAIASRJobResponseBody {
	return s.Body
}

func (s *SubmitAIASRJobResponse) SetHeaders(v map[string]*string) *SubmitAIASRJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIASRJobResponse) SetStatusCode(v int32) *SubmitAIASRJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIASRJobResponse) SetBody(v *SubmitAIASRJobResponseBody) *SubmitAIASRJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIASRJobResponse) Validate() error {
	return dara.Validate(s)
}
