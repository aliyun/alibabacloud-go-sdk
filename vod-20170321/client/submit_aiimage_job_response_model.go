// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIImageJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIImageJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIImageJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIImageJobResponseBody) *SubmitAIImageJobResponse
	GetBody() *SubmitAIImageJobResponseBody
}

type SubmitAIImageJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIImageJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIImageJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIImageJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIImageJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIImageJobResponse) GetBody() *SubmitAIImageJobResponseBody {
	return s.Body
}

func (s *SubmitAIImageJobResponse) SetHeaders(v map[string]*string) *SubmitAIImageJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIImageJobResponse) SetStatusCode(v int32) *SubmitAIImageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIImageJobResponse) SetBody(v *SubmitAIImageJobResponseBody) *SubmitAIImageJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIImageJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
