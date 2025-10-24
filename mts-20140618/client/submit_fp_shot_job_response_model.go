// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpShotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitFpShotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitFpShotJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitFpShotJobResponseBody) *SubmitFpShotJobResponse
	GetBody() *SubmitFpShotJobResponseBody
}

type SubmitFpShotJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitFpShotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitFpShotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpShotJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitFpShotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitFpShotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitFpShotJobResponse) GetBody() *SubmitFpShotJobResponseBody {
	return s.Body
}

func (s *SubmitFpShotJobResponse) SetHeaders(v map[string]*string) *SubmitFpShotJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitFpShotJobResponse) SetStatusCode(v int32) *SubmitFpShotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitFpShotJobResponse) SetBody(v *SubmitFpShotJobResponseBody) *SubmitFpShotJobResponse {
	s.Body = v
	return s
}

func (s *SubmitFpShotJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
