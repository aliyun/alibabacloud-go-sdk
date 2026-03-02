// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProfileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProfileResponseBody) *DeleteProfileResponse
	GetBody() *DeleteProfileResponseBody
}

type DeleteProfileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProfileResponse) GetBody() *DeleteProfileResponseBody {
	return s.Body
}

func (s *DeleteProfileResponse) SetHeaders(v map[string]*string) *DeleteProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteProfileResponse) SetStatusCode(v int32) *DeleteProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProfileResponse) SetBody(v *DeleteProfileResponseBody) *DeleteProfileResponse {
	s.Body = v
	return s
}

func (s *DeleteProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
