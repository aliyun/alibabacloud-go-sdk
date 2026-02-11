// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMseIncidentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMseIncidentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMseIncidentResponse
	GetStatusCode() *int32
	SetBody(v *SendMseIncidentResponseBody) *SendMseIncidentResponse
	GetBody() *SendMseIncidentResponseBody
}

type SendMseIncidentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMseIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMseIncidentResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMseIncidentResponse) GoString() string {
	return s.String()
}

func (s *SendMseIncidentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMseIncidentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMseIncidentResponse) GetBody() *SendMseIncidentResponseBody {
	return s.Body
}

func (s *SendMseIncidentResponse) SetHeaders(v map[string]*string) *SendMseIncidentResponse {
	s.Headers = v
	return s
}

func (s *SendMseIncidentResponse) SetStatusCode(v int32) *SendMseIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMseIncidentResponse) SetBody(v *SendMseIncidentResponseBody) *SendMseIncidentResponse {
	s.Body = v
	return s
}

func (s *SendMseIncidentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
