// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiCallsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceApiCallsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceApiCallsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceApiCallsResponseBody) *ListDataServiceApiCallsResponse
	GetBody() *ListDataServiceApiCallsResponseBody
}

type ListDataServiceApiCallsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceApiCallsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceApiCallsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceApiCallsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceApiCallsResponse) GetBody() *ListDataServiceApiCallsResponseBody {
	return s.Body
}

func (s *ListDataServiceApiCallsResponse) SetHeaders(v map[string]*string) *ListDataServiceApiCallsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceApiCallsResponse) SetStatusCode(v int32) *ListDataServiceApiCallsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceApiCallsResponse) SetBody(v *ListDataServiceApiCallsResponseBody) *ListDataServiceApiCallsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceApiCallsResponse) Validate() error {
	return dara.Validate(s)
}
