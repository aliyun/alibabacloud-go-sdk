// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaintainWindowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMaintainWindowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMaintainWindowsResponse
	GetStatusCode() *int32
	SetBody(v *ListMaintainWindowsResponseBody) *ListMaintainWindowsResponse
	GetBody() *ListMaintainWindowsResponseBody
}

type ListMaintainWindowsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMaintainWindowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMaintainWindowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMaintainWindowsResponse) GoString() string {
	return s.String()
}

func (s *ListMaintainWindowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMaintainWindowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMaintainWindowsResponse) GetBody() *ListMaintainWindowsResponseBody {
	return s.Body
}

func (s *ListMaintainWindowsResponse) SetHeaders(v map[string]*string) *ListMaintainWindowsResponse {
	s.Headers = v
	return s
}

func (s *ListMaintainWindowsResponse) SetStatusCode(v int32) *ListMaintainWindowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMaintainWindowsResponse) SetBody(v *ListMaintainWindowsResponseBody) *ListMaintainWindowsResponse {
	s.Body = v
	return s
}

func (s *ListMaintainWindowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
