// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfraredRemoteControllersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInfraredRemoteControllersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInfraredRemoteControllersResponse
	GetStatusCode() *int32
	SetBody(v *ListInfraredRemoteControllersResponseBody) *ListInfraredRemoteControllersResponse
	GetBody() *ListInfraredRemoteControllersResponseBody
}

type ListInfraredRemoteControllersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInfraredRemoteControllersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInfraredRemoteControllersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredRemoteControllersResponse) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInfraredRemoteControllersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInfraredRemoteControllersResponse) GetBody() *ListInfraredRemoteControllersResponseBody {
	return s.Body
}

func (s *ListInfraredRemoteControllersResponse) SetHeaders(v map[string]*string) *ListInfraredRemoteControllersResponse {
	s.Headers = v
	return s
}

func (s *ListInfraredRemoteControllersResponse) SetStatusCode(v int32) *ListInfraredRemoteControllersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInfraredRemoteControllersResponse) SetBody(v *ListInfraredRemoteControllersResponseBody) *ListInfraredRemoteControllersResponse {
	s.Body = v
	return s
}

func (s *ListInfraredRemoteControllersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
