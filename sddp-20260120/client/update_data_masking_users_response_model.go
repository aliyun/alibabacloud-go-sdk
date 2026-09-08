// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataMaskingUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataMaskingUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataMaskingUsersResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataMaskingUsersResponseBody) *UpdateDataMaskingUsersResponse
	GetBody() *UpdateDataMaskingUsersResponseBody
}

type UpdateDataMaskingUsersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataMaskingUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataMaskingUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingUsersResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataMaskingUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataMaskingUsersResponse) GetBody() *UpdateDataMaskingUsersResponseBody {
	return s.Body
}

func (s *UpdateDataMaskingUsersResponse) SetHeaders(v map[string]*string) *UpdateDataMaskingUsersResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataMaskingUsersResponse) SetStatusCode(v int32) *UpdateDataMaskingUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataMaskingUsersResponse) SetBody(v *UpdateDataMaskingUsersResponseBody) *UpdateDataMaskingUsersResponse {
	s.Body = v
	return s
}

func (s *UpdateDataMaskingUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
