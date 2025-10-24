// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpDBDeleteJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitFpDBDeleteJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitFpDBDeleteJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitFpDBDeleteJobResponseBody) *SubmitFpDBDeleteJobResponse
	GetBody() *SubmitFpDBDeleteJobResponseBody
}

type SubmitFpDBDeleteJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitFpDBDeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitFpDBDeleteJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpDBDeleteJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitFpDBDeleteJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitFpDBDeleteJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitFpDBDeleteJobResponse) GetBody() *SubmitFpDBDeleteJobResponseBody {
	return s.Body
}

func (s *SubmitFpDBDeleteJobResponse) SetHeaders(v map[string]*string) *SubmitFpDBDeleteJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitFpDBDeleteJobResponse) SetStatusCode(v int32) *SubmitFpDBDeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitFpDBDeleteJobResponse) SetBody(v *SubmitFpDBDeleteJobResponseBody) *SubmitFpDBDeleteJobResponse {
	s.Body = v
	return s
}

func (s *SubmitFpDBDeleteJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
