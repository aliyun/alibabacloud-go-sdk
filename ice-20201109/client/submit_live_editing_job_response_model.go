// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveEditingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitLiveEditingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitLiveEditingJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitLiveEditingJobResponseBody) *SubmitLiveEditingJobResponse
	GetBody() *SubmitLiveEditingJobResponseBody
}

type SubmitLiveEditingJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitLiveEditingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitLiveEditingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveEditingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitLiveEditingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitLiveEditingJobResponse) GetBody() *SubmitLiveEditingJobResponseBody {
	return s.Body
}

func (s *SubmitLiveEditingJobResponse) SetHeaders(v map[string]*string) *SubmitLiveEditingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitLiveEditingJobResponse) SetStatusCode(v int32) *SubmitLiveEditingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitLiveEditingJobResponse) SetBody(v *SubmitLiveEditingJobResponseBody) *SubmitLiveEditingJobResponse {
	s.Body = v
	return s
}

func (s *SubmitLiveEditingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
