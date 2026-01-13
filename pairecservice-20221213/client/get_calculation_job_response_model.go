// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCalculationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCalculationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCalculationJobResponse
	GetStatusCode() *int32
	SetBody(v *GetCalculationJobResponseBody) *GetCalculationJobResponse
	GetBody() *GetCalculationJobResponseBody
}

type GetCalculationJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCalculationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCalculationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCalculationJobResponse) GoString() string {
	return s.String()
}

func (s *GetCalculationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCalculationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCalculationJobResponse) GetBody() *GetCalculationJobResponseBody {
	return s.Body
}

func (s *GetCalculationJobResponse) SetHeaders(v map[string]*string) *GetCalculationJobResponse {
	s.Headers = v
	return s
}

func (s *GetCalculationJobResponse) SetStatusCode(v int32) *GetCalculationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCalculationJobResponse) SetBody(v *GetCalculationJobResponseBody) *GetCalculationJobResponse {
	s.Body = v
	return s
}

func (s *GetCalculationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
