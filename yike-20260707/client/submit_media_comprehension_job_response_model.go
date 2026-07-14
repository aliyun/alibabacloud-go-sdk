// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaComprehensionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMediaComprehensionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMediaComprehensionJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMediaComprehensionJobResponseBody) *SubmitMediaComprehensionJobResponse
	GetBody() *SubmitMediaComprehensionJobResponseBody
}

type SubmitMediaComprehensionJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMediaComprehensionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMediaComprehensionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaComprehensionJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaComprehensionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMediaComprehensionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMediaComprehensionJobResponse) GetBody() *SubmitMediaComprehensionJobResponseBody {
	return s.Body
}

func (s *SubmitMediaComprehensionJobResponse) SetHeaders(v map[string]*string) *SubmitMediaComprehensionJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaComprehensionJobResponse) SetStatusCode(v int32) *SubmitMediaComprehensionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMediaComprehensionJobResponse) SetBody(v *SubmitMediaComprehensionJobResponseBody) *SubmitMediaComprehensionJobResponse {
	s.Body = v
	return s
}

func (s *SubmitMediaComprehensionJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
