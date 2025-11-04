// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvatarTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAvatarTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAvatarTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *GetAvatarTrainingJobResponseBody) *GetAvatarTrainingJobResponse
	GetBody() *GetAvatarTrainingJobResponseBody
}

type GetAvatarTrainingJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAvatarTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAvatarTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *GetAvatarTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAvatarTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAvatarTrainingJobResponse) GetBody() *GetAvatarTrainingJobResponseBody {
	return s.Body
}

func (s *GetAvatarTrainingJobResponse) SetHeaders(v map[string]*string) *GetAvatarTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *GetAvatarTrainingJobResponse) SetStatusCode(v int32) *GetAvatarTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAvatarTrainingJobResponse) SetBody(v *GetAvatarTrainingJobResponseBody) *GetAvatarTrainingJobResponse {
	s.Body = v
	return s
}

func (s *GetAvatarTrainingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
