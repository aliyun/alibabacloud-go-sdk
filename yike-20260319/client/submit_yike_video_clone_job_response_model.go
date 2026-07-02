// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeVideoCloneJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitYikeVideoCloneJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitYikeVideoCloneJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitYikeVideoCloneJobResponseBody) *SubmitYikeVideoCloneJobResponse
	GetBody() *SubmitYikeVideoCloneJobResponseBody
}

type SubmitYikeVideoCloneJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitYikeVideoCloneJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitYikeVideoCloneJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeVideoCloneJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitYikeVideoCloneJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitYikeVideoCloneJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitYikeVideoCloneJobResponse) GetBody() *SubmitYikeVideoCloneJobResponseBody {
	return s.Body
}

func (s *SubmitYikeVideoCloneJobResponse) SetHeaders(v map[string]*string) *SubmitYikeVideoCloneJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitYikeVideoCloneJobResponse) SetStatusCode(v int32) *SubmitYikeVideoCloneJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitYikeVideoCloneJobResponse) SetBody(v *SubmitYikeVideoCloneJobResponseBody) *SubmitYikeVideoCloneJobResponse {
	s.Body = v
	return s
}

func (s *SubmitYikeVideoCloneJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
