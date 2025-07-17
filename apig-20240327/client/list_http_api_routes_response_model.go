// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApiRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpApiRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpApiRoutesResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpApiRoutesResponseBody) *ListHttpApiRoutesResponse
	GetBody() *ListHttpApiRoutesResponseBody
}

type ListHttpApiRoutesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpApiRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpApiRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApiRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListHttpApiRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpApiRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpApiRoutesResponse) GetBody() *ListHttpApiRoutesResponseBody {
	return s.Body
}

func (s *ListHttpApiRoutesResponse) SetHeaders(v map[string]*string) *ListHttpApiRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListHttpApiRoutesResponse) SetStatusCode(v int32) *ListHttpApiRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpApiRoutesResponse) SetBody(v *ListHttpApiRoutesResponseBody) *ListHttpApiRoutesResponse {
	s.Body = v
	return s
}

func (s *ListHttpApiRoutesResponse) Validate() error {
	return dara.Validate(s)
}
