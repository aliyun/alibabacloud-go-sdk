// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaInfoJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMediaInfoJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMediaInfoJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMediaInfoJobResponseBody) *SubmitMediaInfoJobResponse
	GetBody() *SubmitMediaInfoJobResponseBody
}

type SubmitMediaInfoJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMediaInfoJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMediaInfoJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMediaInfoJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMediaInfoJobResponse) GetBody() *SubmitMediaInfoJobResponseBody {
	return s.Body
}

func (s *SubmitMediaInfoJobResponse) SetHeaders(v map[string]*string) *SubmitMediaInfoJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaInfoJobResponse) SetStatusCode(v int32) *SubmitMediaInfoJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMediaInfoJobResponse) SetBody(v *SubmitMediaInfoJobResponseBody) *SubmitMediaInfoJobResponse {
	s.Body = v
	return s
}

func (s *SubmitMediaInfoJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
