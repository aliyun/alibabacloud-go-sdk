// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistrantProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRegistrantProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRegistrantProfileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRegistrantProfileResponseBody) *DeleteRegistrantProfileResponse
	GetBody() *DeleteRegistrantProfileResponseBody
}

type DeleteRegistrantProfileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRegistrantProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRegistrantProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistrantProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegistrantProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRegistrantProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRegistrantProfileResponse) GetBody() *DeleteRegistrantProfileResponseBody {
	return s.Body
}

func (s *DeleteRegistrantProfileResponse) SetHeaders(v map[string]*string) *DeleteRegistrantProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegistrantProfileResponse) SetStatusCode(v int32) *DeleteRegistrantProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRegistrantProfileResponse) SetBody(v *DeleteRegistrantProfileResponseBody) *DeleteRegistrantProfileResponse {
	s.Body = v
	return s
}

func (s *DeleteRegistrantProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
