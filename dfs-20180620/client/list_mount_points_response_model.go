// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMountPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMountPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMountPointsResponse
	GetStatusCode() *int32
	SetBody(v *ListMountPointsResponseBody) *ListMountPointsResponse
	GetBody() *ListMountPointsResponseBody
}

type ListMountPointsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMountPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMountPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMountPointsResponse) GoString() string {
	return s.String()
}

func (s *ListMountPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMountPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMountPointsResponse) GetBody() *ListMountPointsResponseBody {
	return s.Body
}

func (s *ListMountPointsResponse) SetHeaders(v map[string]*string) *ListMountPointsResponse {
	s.Headers = v
	return s
}

func (s *ListMountPointsResponse) SetStatusCode(v int32) *ListMountPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMountPointsResponse) SetBody(v *ListMountPointsResponseBody) *ListMountPointsResponse {
	s.Body = v
	return s
}

func (s *ListMountPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
