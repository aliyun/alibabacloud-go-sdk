// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBusinessHoursResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckBusinessHoursResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckBusinessHoursResponse
	GetStatusCode() *int32
	SetBody(v *CheckBusinessHoursResponseBody) *CheckBusinessHoursResponse
	GetBody() *CheckBusinessHoursResponseBody
}

type CheckBusinessHoursResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckBusinessHoursResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckBusinessHoursResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckBusinessHoursResponse) GoString() string {
	return s.String()
}

func (s *CheckBusinessHoursResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckBusinessHoursResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckBusinessHoursResponse) GetBody() *CheckBusinessHoursResponseBody {
	return s.Body
}

func (s *CheckBusinessHoursResponse) SetHeaders(v map[string]*string) *CheckBusinessHoursResponse {
	s.Headers = v
	return s
}

func (s *CheckBusinessHoursResponse) SetStatusCode(v int32) *CheckBusinessHoursResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckBusinessHoursResponse) SetBody(v *CheckBusinessHoursResponseBody) *CheckBusinessHoursResponse {
	s.Body = v
	return s
}

func (s *CheckBusinessHoursResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
