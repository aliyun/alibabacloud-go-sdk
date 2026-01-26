// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAlertResponse
	GetStatusCode() *int32
	SetBody(v *StartAlertResponseBody) *StartAlertResponse
	GetBody() *StartAlertResponseBody
}

type StartAlertResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAlertResponse) GoString() string {
	return s.String()
}

func (s *StartAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAlertResponse) GetBody() *StartAlertResponseBody {
	return s.Body
}

func (s *StartAlertResponse) SetHeaders(v map[string]*string) *StartAlertResponse {
	s.Headers = v
	return s
}

func (s *StartAlertResponse) SetStatusCode(v int32) *StartAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAlertResponse) SetBody(v *StartAlertResponseBody) *StartAlertResponse {
	s.Body = v
	return s
}

func (s *StartAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
