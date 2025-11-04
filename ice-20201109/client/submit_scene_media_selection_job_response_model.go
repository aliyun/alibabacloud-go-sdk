// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneMediaSelectionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSceneMediaSelectionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSceneMediaSelectionJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSceneMediaSelectionJobResponseBody) *SubmitSceneMediaSelectionJobResponse
	GetBody() *SubmitSceneMediaSelectionJobResponseBody
}

type SubmitSceneMediaSelectionJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSceneMediaSelectionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSceneMediaSelectionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneMediaSelectionJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSceneMediaSelectionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSceneMediaSelectionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSceneMediaSelectionJobResponse) GetBody() *SubmitSceneMediaSelectionJobResponseBody {
	return s.Body
}

func (s *SubmitSceneMediaSelectionJobResponse) SetHeaders(v map[string]*string) *SubmitSceneMediaSelectionJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSceneMediaSelectionJobResponse) SetStatusCode(v int32) *SubmitSceneMediaSelectionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSceneMediaSelectionJobResponse) SetBody(v *SubmitSceneMediaSelectionJobResponseBody) *SubmitSceneMediaSelectionJobResponse {
	s.Body = v
	return s
}

func (s *SubmitSceneMediaSelectionJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
