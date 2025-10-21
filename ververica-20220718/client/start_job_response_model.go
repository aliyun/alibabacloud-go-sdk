// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartJobResponse
	GetStatusCode() *int32
	SetBody(v *StartJobResponseBody) *StartJobResponse
	GetBody() *StartJobResponseBody
}

type StartJobResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartJobResponse) GoString() string {
	return s.String()
}

func (s *StartJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartJobResponse) GetBody() *StartJobResponseBody {
	return s.Body
}

func (s *StartJobResponse) SetHeaders(v map[string]*string) *StartJobResponse {
	s.Headers = v
	return s
}

func (s *StartJobResponse) SetStatusCode(v int32) *StartJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartJobResponse) SetBody(v *StartJobResponseBody) *StartJobResponse {
	s.Body = v
	return s
}

func (s *StartJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
