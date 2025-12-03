// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficSpecialControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrafficSpecialControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrafficSpecialControlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrafficSpecialControlResponseBody) *DeleteTrafficSpecialControlResponse
	GetBody() *DeleteTrafficSpecialControlResponseBody
}

type DeleteTrafficSpecialControlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficSpecialControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficSpecialControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficSpecialControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrafficSpecialControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrafficSpecialControlResponse) GetBody() *DeleteTrafficSpecialControlResponseBody {
	return s.Body
}

func (s *DeleteTrafficSpecialControlResponse) SetHeaders(v map[string]*string) *DeleteTrafficSpecialControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficSpecialControlResponse) SetStatusCode(v int32) *DeleteTrafficSpecialControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficSpecialControlResponse) SetBody(v *DeleteTrafficSpecialControlResponseBody) *DeleteTrafficSpecialControlResponse {
	s.Body = v
	return s
}

func (s *DeleteTrafficSpecialControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
