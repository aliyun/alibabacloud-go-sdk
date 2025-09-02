// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceApisResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceApisResponseBody) *ListDataServiceApisResponse
	GetBody() *ListDataServiceApisResponseBody
}

type ListDataServiceApisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceApisResponse) GetBody() *ListDataServiceApisResponseBody {
	return s.Body
}

func (s *ListDataServiceApisResponse) SetHeaders(v map[string]*string) *ListDataServiceApisResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceApisResponse) SetStatusCode(v int32) *ListDataServiceApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceApisResponse) SetBody(v *ListDataServiceApisResponseBody) *ListDataServiceApisResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceApisResponse) Validate() error {
	return dara.Validate(s)
}
