// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoRemakeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoRemakeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoRemakeJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoRemakeJobResponseBody) *SubmitVideoRemakeJobResponse
	GetBody() *SubmitVideoRemakeJobResponseBody
}

type SubmitVideoRemakeJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoRemakeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoRemakeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoRemakeJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoRemakeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoRemakeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoRemakeJobResponse) GetBody() *SubmitVideoRemakeJobResponseBody {
	return s.Body
}

func (s *SubmitVideoRemakeJobResponse) SetHeaders(v map[string]*string) *SubmitVideoRemakeJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoRemakeJobResponse) SetStatusCode(v int32) *SubmitVideoRemakeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoRemakeJobResponse) SetBody(v *SubmitVideoRemakeJobResponseBody) *SubmitVideoRemakeJobResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoRemakeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
