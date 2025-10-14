// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIncidentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIncidentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIncidentResponse
	GetStatusCode() *int32
	SetBody(v *GetIncidentResponseBody) *GetIncidentResponse
	GetBody() *GetIncidentResponseBody
}

type GetIncidentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIncidentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIncidentResponse) GoString() string {
	return s.String()
}

func (s *GetIncidentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIncidentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIncidentResponse) GetBody() *GetIncidentResponseBody {
	return s.Body
}

func (s *GetIncidentResponse) SetHeaders(v map[string]*string) *GetIncidentResponse {
	s.Headers = v
	return s
}

func (s *GetIncidentResponse) SetStatusCode(v int32) *GetIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIncidentResponse) SetBody(v *GetIncidentResponseBody) *GetIncidentResponse {
	s.Body = v
	return s
}

func (s *GetIncidentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
