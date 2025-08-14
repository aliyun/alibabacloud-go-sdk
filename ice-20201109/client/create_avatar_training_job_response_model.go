// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAvatarTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAvatarTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAvatarTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateAvatarTrainingJobResponseBody) *CreateAvatarTrainingJobResponse
	GetBody() *CreateAvatarTrainingJobResponseBody
}

type CreateAvatarTrainingJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAvatarTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAvatarTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAvatarTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *CreateAvatarTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAvatarTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAvatarTrainingJobResponse) GetBody() *CreateAvatarTrainingJobResponseBody {
	return s.Body
}

func (s *CreateAvatarTrainingJobResponse) SetHeaders(v map[string]*string) *CreateAvatarTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *CreateAvatarTrainingJobResponse) SetStatusCode(v int32) *CreateAvatarTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAvatarTrainingJobResponse) SetBody(v *CreateAvatarTrainingJobResponseBody) *CreateAvatarTrainingJobResponse {
	s.Body = v
	return s
}

func (s *CreateAvatarTrainingJobResponse) Validate() error {
	return dara.Validate(s)
}
