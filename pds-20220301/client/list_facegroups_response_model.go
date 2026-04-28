// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFacegroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFacegroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFacegroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListFacegroupsResponseBody) *ListFacegroupsResponse
	GetBody() *ListFacegroupsResponseBody
}

type ListFacegroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFacegroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFacegroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFacegroupsResponse) GoString() string {
	return s.String()
}

func (s *ListFacegroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFacegroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFacegroupsResponse) GetBody() *ListFacegroupsResponseBody {
	return s.Body
}

func (s *ListFacegroupsResponse) SetHeaders(v map[string]*string) *ListFacegroupsResponse {
	s.Headers = v
	return s
}

func (s *ListFacegroupsResponse) SetStatusCode(v int32) *ListFacegroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFacegroupsResponse) SetBody(v *ListFacegroupsResponseBody) *ListFacegroupsResponse {
	s.Body = v
	return s
}

func (s *ListFacegroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
