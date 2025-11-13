// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneBatchEditingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSceneBatchEditingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSceneBatchEditingJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSceneBatchEditingJobResponseBody) *SubmitSceneBatchEditingJobResponse
	GetBody() *SubmitSceneBatchEditingJobResponseBody
}

type SubmitSceneBatchEditingJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSceneBatchEditingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSceneBatchEditingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneBatchEditingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSceneBatchEditingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSceneBatchEditingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSceneBatchEditingJobResponse) GetBody() *SubmitSceneBatchEditingJobResponseBody {
	return s.Body
}

func (s *SubmitSceneBatchEditingJobResponse) SetHeaders(v map[string]*string) *SubmitSceneBatchEditingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSceneBatchEditingJobResponse) SetStatusCode(v int32) *SubmitSceneBatchEditingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSceneBatchEditingJobResponse) SetBody(v *SubmitSceneBatchEditingJobResponseBody) *SubmitSceneBatchEditingJobResponse {
	s.Body = v
	return s
}

func (s *SubmitSceneBatchEditingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
