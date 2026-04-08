// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDIJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDIJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDIJobResponse
	GetStatusCode() *int32
	SetBody(v *StartDIJobResponseBody) *StartDIJobResponse
	GetBody() *StartDIJobResponseBody
}

type StartDIJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDIJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDIJobResponse) GoString() string {
	return s.String()
}

func (s *StartDIJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDIJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDIJobResponse) GetBody() *StartDIJobResponseBody {
	return s.Body
}

func (s *StartDIJobResponse) SetHeaders(v map[string]*string) *StartDIJobResponse {
	s.Headers = v
	return s
}

func (s *StartDIJobResponse) SetStatusCode(v int32) *StartDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDIJobResponse) SetBody(v *StartDIJobResponseBody) *StartDIJobResponse {
	s.Body = v
	return s
}

func (s *StartDIJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
