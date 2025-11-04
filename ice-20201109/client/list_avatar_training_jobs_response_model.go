// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarTrainingJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvatarTrainingJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvatarTrainingJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListAvatarTrainingJobsResponseBody) *ListAvatarTrainingJobsResponse
	GetBody() *ListAvatarTrainingJobsResponseBody
}

type ListAvatarTrainingJobsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvatarTrainingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvatarTrainingJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarTrainingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListAvatarTrainingJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvatarTrainingJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvatarTrainingJobsResponse) GetBody() *ListAvatarTrainingJobsResponseBody {
	return s.Body
}

func (s *ListAvatarTrainingJobsResponse) SetHeaders(v map[string]*string) *ListAvatarTrainingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListAvatarTrainingJobsResponse) SetStatusCode(v int32) *ListAvatarTrainingJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvatarTrainingJobsResponse) SetBody(v *ListAvatarTrainingJobsResponseBody) *ListAvatarTrainingJobsResponse {
	s.Body = v
	return s
}

func (s *ListAvatarTrainingJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
