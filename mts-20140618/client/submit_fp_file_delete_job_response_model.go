// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpFileDeleteJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitFpFileDeleteJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitFpFileDeleteJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitFpFileDeleteJobResponseBody) *SubmitFpFileDeleteJobResponse
	GetBody() *SubmitFpFileDeleteJobResponseBody
}

type SubmitFpFileDeleteJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitFpFileDeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitFpFileDeleteJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpFileDeleteJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitFpFileDeleteJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitFpFileDeleteJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitFpFileDeleteJobResponse) GetBody() *SubmitFpFileDeleteJobResponseBody {
	return s.Body
}

func (s *SubmitFpFileDeleteJobResponse) SetHeaders(v map[string]*string) *SubmitFpFileDeleteJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitFpFileDeleteJobResponse) SetStatusCode(v int32) *SubmitFpFileDeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitFpFileDeleteJobResponse) SetBody(v *SubmitFpFileDeleteJobResponseBody) *SubmitFpFileDeleteJobResponse {
	s.Body = v
	return s
}

func (s *SubmitFpFileDeleteJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
