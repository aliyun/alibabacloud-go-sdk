// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportHotelConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportHotelConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportHotelConfigResponse
	GetStatusCode() *int32
	SetBody(v *ImportHotelConfigResponseBody) *ImportHotelConfigResponse
	GetBody() *ImportHotelConfigResponseBody
}

type ImportHotelConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportHotelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportHotelConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportHotelConfigResponse) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportHotelConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportHotelConfigResponse) GetBody() *ImportHotelConfigResponseBody {
	return s.Body
}

func (s *ImportHotelConfigResponse) SetHeaders(v map[string]*string) *ImportHotelConfigResponse {
	s.Headers = v
	return s
}

func (s *ImportHotelConfigResponse) SetStatusCode(v int32) *ImportHotelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportHotelConfigResponse) SetBody(v *ImportHotelConfigResponseBody) *ImportHotelConfigResponse {
	s.Body = v
	return s
}

func (s *ImportHotelConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
