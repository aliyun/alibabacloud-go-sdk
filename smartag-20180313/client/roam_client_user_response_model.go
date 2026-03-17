// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoamClientUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RoamClientUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RoamClientUserResponse
	GetStatusCode() *int32
	SetBody(v *RoamClientUserResponseBody) *RoamClientUserResponse
	GetBody() *RoamClientUserResponseBody
}

type RoamClientUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RoamClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RoamClientUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RoamClientUserResponse) GoString() string {
	return s.String()
}

func (s *RoamClientUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RoamClientUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RoamClientUserResponse) GetBody() *RoamClientUserResponseBody {
	return s.Body
}

func (s *RoamClientUserResponse) SetHeaders(v map[string]*string) *RoamClientUserResponse {
	s.Headers = v
	return s
}

func (s *RoamClientUserResponse) SetStatusCode(v int32) *RoamClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RoamClientUserResponse) SetBody(v *RoamClientUserResponseBody) *RoamClientUserResponse {
	s.Body = v
	return s
}

func (s *RoamClientUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
