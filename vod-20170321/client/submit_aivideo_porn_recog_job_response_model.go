// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoPornRecogJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIVideoPornRecogJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIVideoPornRecogJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIVideoPornRecogJobResponseBody) *SubmitAIVideoPornRecogJobResponse
	GetBody() *SubmitAIVideoPornRecogJobResponseBody
}

type SubmitAIVideoPornRecogJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIVideoPornRecogJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIVideoPornRecogJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoPornRecogJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoPornRecogJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIVideoPornRecogJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIVideoPornRecogJobResponse) GetBody() *SubmitAIVideoPornRecogJobResponseBody {
	return s.Body
}

func (s *SubmitAIVideoPornRecogJobResponse) SetHeaders(v map[string]*string) *SubmitAIVideoPornRecogJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponse) SetStatusCode(v int32) *SubmitAIVideoPornRecogJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponse) SetBody(v *SubmitAIVideoPornRecogJobResponseBody) *SubmitAIVideoPornRecogJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponse) Validate() error {
	return dara.Validate(s)
}
