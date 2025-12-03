// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectWorkitemTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectWorkitemTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectWorkitemTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectWorkitemTypesResponseBody) *ListProjectWorkitemTypesResponse
	GetBody() *ListProjectWorkitemTypesResponseBody
}

type ListProjectWorkitemTypesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectWorkitemTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectWorkitemTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectWorkitemTypesResponse) GoString() string {
	return s.String()
}

func (s *ListProjectWorkitemTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectWorkitemTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectWorkitemTypesResponse) GetBody() *ListProjectWorkitemTypesResponseBody {
	return s.Body
}

func (s *ListProjectWorkitemTypesResponse) SetHeaders(v map[string]*string) *ListProjectWorkitemTypesResponse {
	s.Headers = v
	return s
}

func (s *ListProjectWorkitemTypesResponse) SetStatusCode(v int32) *ListProjectWorkitemTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectWorkitemTypesResponse) SetBody(v *ListProjectWorkitemTypesResponseBody) *ListProjectWorkitemTypesResponse {
	s.Body = v
	return s
}

func (s *ListProjectWorkitemTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
