// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTrialFixCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckTrialFixCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckTrialFixCountResponse
	GetStatusCode() *int32
	SetBody(v *CheckTrialFixCountResponseBody) *CheckTrialFixCountResponse
	GetBody() *CheckTrialFixCountResponseBody
}

type CheckTrialFixCountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckTrialFixCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckTrialFixCountResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckTrialFixCountResponse) GoString() string {
	return s.String()
}

func (s *CheckTrialFixCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckTrialFixCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckTrialFixCountResponse) GetBody() *CheckTrialFixCountResponseBody {
	return s.Body
}

func (s *CheckTrialFixCountResponse) SetHeaders(v map[string]*string) *CheckTrialFixCountResponse {
	s.Headers = v
	return s
}

func (s *CheckTrialFixCountResponse) SetStatusCode(v int32) *CheckTrialFixCountResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTrialFixCountResponse) SetBody(v *CheckTrialFixCountResponseBody) *CheckTrialFixCountResponse {
	s.Body = v
	return s
}

func (s *CheckTrialFixCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
