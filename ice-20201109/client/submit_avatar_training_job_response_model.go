// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAvatarTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAvatarTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAvatarTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAvatarTrainingJobResponseBody) *SubmitAvatarTrainingJobResponse
	GetBody() *SubmitAvatarTrainingJobResponseBody
}

type SubmitAvatarTrainingJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAvatarTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAvatarTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAvatarTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAvatarTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAvatarTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAvatarTrainingJobResponse) GetBody() *SubmitAvatarTrainingJobResponseBody {
	return s.Body
}

func (s *SubmitAvatarTrainingJobResponse) SetHeaders(v map[string]*string) *SubmitAvatarTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAvatarTrainingJobResponse) SetStatusCode(v int32) *SubmitAvatarTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAvatarTrainingJobResponse) SetBody(v *SubmitAvatarTrainingJobResponseBody) *SubmitAvatarTrainingJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAvatarTrainingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
