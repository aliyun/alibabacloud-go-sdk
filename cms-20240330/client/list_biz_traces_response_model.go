// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizTracesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBizTracesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBizTracesResponse
	GetStatusCode() *int32
	SetBody(v *ListBizTracesResponseBody) *ListBizTracesResponse
	GetBody() *ListBizTracesResponseBody
}

type ListBizTracesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBizTracesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBizTracesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBizTracesResponse) GoString() string {
	return s.String()
}

func (s *ListBizTracesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBizTracesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBizTracesResponse) GetBody() *ListBizTracesResponseBody {
	return s.Body
}

func (s *ListBizTracesResponse) SetHeaders(v map[string]*string) *ListBizTracesResponse {
	s.Headers = v
	return s
}

func (s *ListBizTracesResponse) SetStatusCode(v int32) *ListBizTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBizTracesResponse) SetBody(v *ListBizTracesResponseBody) *ListBizTracesResponse {
	s.Body = v
	return s
}

func (s *ListBizTracesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
