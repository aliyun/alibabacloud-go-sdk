// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVolumesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVolumesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVolumesResponse
	GetStatusCode() *int32
	SetBody(v *ListVolumesResponseBody) *ListVolumesResponse
	GetBody() *ListVolumesResponseBody
}

type ListVolumesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVolumesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVolumesResponse) GoString() string {
	return s.String()
}

func (s *ListVolumesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVolumesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVolumesResponse) GetBody() *ListVolumesResponseBody {
	return s.Body
}

func (s *ListVolumesResponse) SetHeaders(v map[string]*string) *ListVolumesResponse {
	s.Headers = v
	return s
}

func (s *ListVolumesResponse) SetStatusCode(v int32) *ListVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVolumesResponse) SetBody(v *ListVolumesResponseBody) *ListVolumesResponse {
	s.Body = v
	return s
}

func (s *ListVolumesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
