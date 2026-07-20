// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountSafetyIncidentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountSafetyIncidentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountSafetyIncidentResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountSafetyIncidentResponseBody) *QueryAccountSafetyIncidentResponse
	GetBody() *QueryAccountSafetyIncidentResponseBody
}

type QueryAccountSafetyIncidentResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountSafetyIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountSafetyIncidentResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSafetyIncidentResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountSafetyIncidentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountSafetyIncidentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountSafetyIncidentResponse) GetBody() *QueryAccountSafetyIncidentResponseBody {
	return s.Body
}

func (s *QueryAccountSafetyIncidentResponse) SetHeaders(v map[string]*string) *QueryAccountSafetyIncidentResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountSafetyIncidentResponse) SetStatusCode(v int32) *QueryAccountSafetyIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponse) SetBody(v *QueryAccountSafetyIncidentResponseBody) *QueryAccountSafetyIncidentResponse {
	s.Body = v
	return s
}

func (s *QueryAccountSafetyIncidentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
