// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAvatarTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAvatarTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAvatarTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAvatarTrainingJobResponseBody) *DeleteAvatarTrainingJobResponse
	GetBody() *DeleteAvatarTrainingJobResponseBody
}

type DeleteAvatarTrainingJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAvatarTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAvatarTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAvatarTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteAvatarTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAvatarTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAvatarTrainingJobResponse) GetBody() *DeleteAvatarTrainingJobResponseBody {
	return s.Body
}

func (s *DeleteAvatarTrainingJobResponse) SetHeaders(v map[string]*string) *DeleteAvatarTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteAvatarTrainingJobResponse) SetStatusCode(v int32) *DeleteAvatarTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAvatarTrainingJobResponse) SetBody(v *DeleteAvatarTrainingJobResponseBody) *DeleteAvatarTrainingJobResponse {
	s.Body = v
	return s
}

func (s *DeleteAvatarTrainingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
