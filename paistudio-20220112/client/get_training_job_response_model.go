// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *GetTrainingJobResponseBody) *GetTrainingJobResponse
	GetBody() *GetTrainingJobResponseBody
}

type GetTrainingJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrainingJobResponse) GetBody() *GetTrainingJobResponseBody {
	return s.Body
}

func (s *GetTrainingJobResponse) SetHeaders(v map[string]*string) *GetTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *GetTrainingJobResponse) SetStatusCode(v int32) *GetTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainingJobResponse) SetBody(v *GetTrainingJobResponseBody) *GetTrainingJobResponse {
	s.Body = v
	return s
}

func (s *GetTrainingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
