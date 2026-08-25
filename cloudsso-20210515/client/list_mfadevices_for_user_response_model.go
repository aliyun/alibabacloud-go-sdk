// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMFADevicesForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMFADevicesForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMFADevicesForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListMFADevicesForUserResponseBody) *ListMFADevicesForUserResponse
	GetBody() *ListMFADevicesForUserResponseBody
}

type ListMFADevicesForUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMFADevicesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMFADevicesForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMFADevicesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListMFADevicesForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMFADevicesForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMFADevicesForUserResponse) GetBody() *ListMFADevicesForUserResponseBody {
	return s.Body
}

func (s *ListMFADevicesForUserResponse) SetHeaders(v map[string]*string) *ListMFADevicesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListMFADevicesForUserResponse) SetStatusCode(v int32) *ListMFADevicesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMFADevicesForUserResponse) SetBody(v *ListMFADevicesForUserResponseBody) *ListMFADevicesForUserResponse {
	s.Body = v
	return s
}

func (s *ListMFADevicesForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
