// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCuHoursResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCuHoursResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCuHoursResponse
	GetStatusCode() *int32
	SetBody(v *GetCuHoursResponseBody) *GetCuHoursResponse
	GetBody() *GetCuHoursResponseBody
}

type GetCuHoursResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCuHoursResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCuHoursResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCuHoursResponse) GoString() string {
	return s.String()
}

func (s *GetCuHoursResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCuHoursResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCuHoursResponse) GetBody() *GetCuHoursResponseBody {
	return s.Body
}

func (s *GetCuHoursResponse) SetHeaders(v map[string]*string) *GetCuHoursResponse {
	s.Headers = v
	return s
}

func (s *GetCuHoursResponse) SetStatusCode(v int32) *GetCuHoursResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCuHoursResponse) SetBody(v *GetCuHoursResponseBody) *GetCuHoursResponse {
	s.Body = v
	return s
}

func (s *GetCuHoursResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
