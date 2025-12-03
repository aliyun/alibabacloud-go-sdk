// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHbaseHaSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHbaseHaSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHbaseHaSlbResponse
	GetStatusCode() *int32
	SetBody(v *CreateHbaseHaSlbResponseBody) *CreateHbaseHaSlbResponse
	GetBody() *CreateHbaseHaSlbResponseBody
}

type CreateHbaseHaSlbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHbaseHaSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHbaseHaSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHbaseHaSlbResponse) GetBody() *CreateHbaseHaSlbResponseBody {
	return s.Body
}

func (s *CreateHbaseHaSlbResponse) SetHeaders(v map[string]*string) *CreateHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *CreateHbaseHaSlbResponse) SetStatusCode(v int32) *CreateHbaseHaSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHbaseHaSlbResponse) SetBody(v *CreateHbaseHaSlbResponseBody) *CreateHbaseHaSlbResponse {
	s.Body = v
	return s
}

func (s *CreateHbaseHaSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
