// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneTimelineOrganizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSceneTimelineOrganizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSceneTimelineOrganizationJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSceneTimelineOrganizationJobResponseBody) *SubmitSceneTimelineOrganizationJobResponse
	GetBody() *SubmitSceneTimelineOrganizationJobResponseBody
}

type SubmitSceneTimelineOrganizationJobResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSceneTimelineOrganizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSceneTimelineOrganizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneTimelineOrganizationJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSceneTimelineOrganizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSceneTimelineOrganizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSceneTimelineOrganizationJobResponse) GetBody() *SubmitSceneTimelineOrganizationJobResponseBody {
	return s.Body
}

func (s *SubmitSceneTimelineOrganizationJobResponse) SetHeaders(v map[string]*string) *SubmitSceneTimelineOrganizationJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobResponse) SetStatusCode(v int32) *SubmitSceneTimelineOrganizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobResponse) SetBody(v *SubmitSceneTimelineOrganizationJobResponseBody) *SubmitSceneTimelineOrganizationJobResponse {
	s.Body = v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
