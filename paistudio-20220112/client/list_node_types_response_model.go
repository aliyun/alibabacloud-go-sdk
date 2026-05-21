// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeTypesResponseBody) *ListNodeTypesResponse
	GetBody() *ListNodeTypesResponseBody
}

type ListNodeTypesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeTypesResponse) GoString() string {
	return s.String()
}

func (s *ListNodeTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeTypesResponse) GetBody() *ListNodeTypesResponseBody {
	return s.Body
}

func (s *ListNodeTypesResponse) SetHeaders(v map[string]*string) *ListNodeTypesResponse {
	s.Headers = v
	return s
}

func (s *ListNodeTypesResponse) SetStatusCode(v int32) *ListNodeTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeTypesResponse) SetBody(v *ListNodeTypesResponseBody) *ListNodeTypesResponse {
	s.Body = v
	return s
}

func (s *ListNodeTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
