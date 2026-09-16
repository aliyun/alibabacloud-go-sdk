// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLumaNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLumaNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *ListLumaNamespacesResponseBody) *ListLumaNamespacesResponse
	GetBody() *ListLumaNamespacesResponseBody
}

type ListLumaNamespacesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLumaNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLumaNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLumaNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListLumaNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLumaNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLumaNamespacesResponse) GetBody() *ListLumaNamespacesResponseBody {
	return s.Body
}

func (s *ListLumaNamespacesResponse) SetHeaders(v map[string]*string) *ListLumaNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListLumaNamespacesResponse) SetStatusCode(v int32) *ListLumaNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLumaNamespacesResponse) SetBody(v *ListLumaNamespacesResponseBody) *ListLumaNamespacesResponse {
	s.Body = v
	return s
}

func (s *ListLumaNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
