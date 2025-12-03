// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHbaseHaSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHbaseHaSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHbaseHaSlbResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHbaseHaSlbResponseBody) *DeleteHbaseHaSlbResponse
	GetBody() *DeleteHbaseHaSlbResponseBody
}

type DeleteHbaseHaSlbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHbaseHaSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHbaseHaSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHbaseHaSlbResponse) GetBody() *DeleteHbaseHaSlbResponseBody {
	return s.Body
}

func (s *DeleteHbaseHaSlbResponse) SetHeaders(v map[string]*string) *DeleteHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *DeleteHbaseHaSlbResponse) SetStatusCode(v int32) *DeleteHbaseHaSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHbaseHaSlbResponse) SetBody(v *DeleteHbaseHaSlbResponseBody) *DeleteHbaseHaSlbResponse {
	s.Body = v
	return s
}

func (s *DeleteHbaseHaSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
