// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteFunctionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSiteFunctionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSiteFunctionsResponse
	GetStatusCode() *int32
	SetBody(v *ListSiteFunctionsResponseBody) *ListSiteFunctionsResponse
	GetBody() *ListSiteFunctionsResponseBody
}

type ListSiteFunctionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSiteFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSiteFunctionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSiteFunctionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSiteFunctionsResponse) GetBody() *ListSiteFunctionsResponseBody {
	return s.Body
}

func (s *ListSiteFunctionsResponse) SetHeaders(v map[string]*string) *ListSiteFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListSiteFunctionsResponse) SetStatusCode(v int32) *ListSiteFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSiteFunctionsResponse) SetBody(v *ListSiteFunctionsResponseBody) *ListSiteFunctionsResponse {
	s.Body = v
	return s
}

func (s *ListSiteFunctionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
