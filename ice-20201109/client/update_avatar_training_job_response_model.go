// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAvatarTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAvatarTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAvatarTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAvatarTrainingJobResponseBody) *UpdateAvatarTrainingJobResponse
	GetBody() *UpdateAvatarTrainingJobResponseBody
}

type UpdateAvatarTrainingJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAvatarTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAvatarTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAvatarTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateAvatarTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAvatarTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAvatarTrainingJobResponse) GetBody() *UpdateAvatarTrainingJobResponseBody {
	return s.Body
}

func (s *UpdateAvatarTrainingJobResponse) SetHeaders(v map[string]*string) *UpdateAvatarTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateAvatarTrainingJobResponse) SetStatusCode(v int32) *UpdateAvatarTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAvatarTrainingJobResponse) SetBody(v *UpdateAvatarTrainingJobResponseBody) *UpdateAvatarTrainingJobResponse {
	s.Body = v
	return s
}

func (s *UpdateAvatarTrainingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
