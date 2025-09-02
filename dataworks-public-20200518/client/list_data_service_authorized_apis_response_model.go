// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAuthorizedApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceAuthorizedApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceAuthorizedApisResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceAuthorizedApisResponseBody) *ListDataServiceAuthorizedApisResponse
	GetBody() *ListDataServiceAuthorizedApisResponseBody
}

type ListDataServiceAuthorizedApisResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceAuthorizedApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceAuthorizedApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedApisResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceAuthorizedApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceAuthorizedApisResponse) GetBody() *ListDataServiceAuthorizedApisResponseBody {
	return s.Body
}

func (s *ListDataServiceAuthorizedApisResponse) SetHeaders(v map[string]*string) *ListDataServiceAuthorizedApisResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceAuthorizedApisResponse) SetStatusCode(v int32) *ListDataServiceAuthorizedApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponse) SetBody(v *ListDataServiceAuthorizedApisResponseBody) *ListDataServiceAuthorizedApisResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceAuthorizedApisResponse) Validate() error {
	return dara.Validate(s)
}
