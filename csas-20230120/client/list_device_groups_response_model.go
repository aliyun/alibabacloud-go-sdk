// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceGroupsResponseBody) *ListDeviceGroupsResponse
	GetBody() *ListDeviceGroupsResponseBody
}

type ListDeviceGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceGroupsResponse) GetBody() *ListDeviceGroupsResponseBody {
	return s.Body
}

func (s *ListDeviceGroupsResponse) SetHeaders(v map[string]*string) *ListDeviceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceGroupsResponse) SetStatusCode(v int32) *ListDeviceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceGroupsResponse) SetBody(v *ListDeviceGroupsResponseBody) *ListDeviceGroupsResponse {
	s.Body = v
	return s
}

func (s *ListDeviceGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
