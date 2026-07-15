// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoGenerationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoGenerationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoGenerationJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoGenerationJobResponseBody) *SubmitVideoGenerationJobResponse
	GetBody() *SubmitVideoGenerationJobResponseBody
}

type SubmitVideoGenerationJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoGenerationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoGenerationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoGenerationJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoGenerationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoGenerationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoGenerationJobResponse) GetBody() *SubmitVideoGenerationJobResponseBody {
	return s.Body
}

func (s *SubmitVideoGenerationJobResponse) SetHeaders(v map[string]*string) *SubmitVideoGenerationJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoGenerationJobResponse) SetStatusCode(v int32) *SubmitVideoGenerationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoGenerationJobResponse) SetBody(v *SubmitVideoGenerationJobResponseBody) *SubmitVideoGenerationJobResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoGenerationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
