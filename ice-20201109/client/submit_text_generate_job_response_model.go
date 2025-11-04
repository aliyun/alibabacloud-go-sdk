// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTextGenerateJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTextGenerateJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTextGenerateJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTextGenerateJobResponseBody) *SubmitTextGenerateJobResponse
	GetBody() *SubmitTextGenerateJobResponseBody
}

type SubmitTextGenerateJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTextGenerateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTextGenerateJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTextGenerateJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitTextGenerateJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTextGenerateJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTextGenerateJobResponse) GetBody() *SubmitTextGenerateJobResponseBody {
	return s.Body
}

func (s *SubmitTextGenerateJobResponse) SetHeaders(v map[string]*string) *SubmitTextGenerateJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitTextGenerateJobResponse) SetStatusCode(v int32) *SubmitTextGenerateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTextGenerateJobResponse) SetBody(v *SubmitTextGenerateJobResponseBody) *SubmitTextGenerateJobResponse {
	s.Body = v
	return s
}

func (s *SubmitTextGenerateJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
