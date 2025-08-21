// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceIdByIdentitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceIdByIdentitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceIdByIdentitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceIdByIdentitiesResponseBody) *ListDeviceIdByIdentitiesResponse
	GetBody() *ListDeviceIdByIdentitiesResponseBody
}

type ListDeviceIdByIdentitiesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceIdByIdentitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceIdByIdentitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceIdByIdentitiesResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceIdByIdentitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceIdByIdentitiesResponse) GetBody() *ListDeviceIdByIdentitiesResponseBody {
	return s.Body
}

func (s *ListDeviceIdByIdentitiesResponse) SetHeaders(v map[string]*string) *ListDeviceIdByIdentitiesResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceIdByIdentitiesResponse) SetStatusCode(v int32) *ListDeviceIdByIdentitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceIdByIdentitiesResponse) SetBody(v *ListDeviceIdByIdentitiesResponseBody) *ListDeviceIdByIdentitiesResponse {
	s.Body = v
	return s
}

func (s *ListDeviceIdByIdentitiesResponse) Validate() error {
	return dara.Validate(s)
}
