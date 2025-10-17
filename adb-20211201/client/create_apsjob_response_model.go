// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAPSJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAPSJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAPSJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateAPSJobResponseBody) *CreateAPSJobResponse
	GetBody() *CreateAPSJobResponseBody
}

type CreateAPSJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAPSJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAPSJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAPSJobResponse) GoString() string {
	return s.String()
}

func (s *CreateAPSJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAPSJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAPSJobResponse) GetBody() *CreateAPSJobResponseBody {
	return s.Body
}

func (s *CreateAPSJobResponse) SetHeaders(v map[string]*string) *CreateAPSJobResponse {
	s.Headers = v
	return s
}

func (s *CreateAPSJobResponse) SetStatusCode(v int32) *CreateAPSJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAPSJobResponse) SetBody(v *CreateAPSJobResponseBody) *CreateAPSJobResponse {
	s.Body = v
	return s
}

func (s *CreateAPSJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
