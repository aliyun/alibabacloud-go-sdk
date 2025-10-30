// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDynamicRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDynamicRoutesResponse
	GetStatusCode() *int32
	SetBody(v *ListDynamicRoutesResponseBody) *ListDynamicRoutesResponse
	GetBody() *ListDynamicRoutesResponseBody
}

type ListDynamicRoutesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDynamicRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDynamicRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDynamicRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDynamicRoutesResponse) GetBody() *ListDynamicRoutesResponseBody {
	return s.Body
}

func (s *ListDynamicRoutesResponse) SetHeaders(v map[string]*string) *ListDynamicRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicRoutesResponse) SetStatusCode(v int32) *ListDynamicRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDynamicRoutesResponse) SetBody(v *ListDynamicRoutesResponseBody) *ListDynamicRoutesResponse {
	s.Body = v
	return s
}

func (s *ListDynamicRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
