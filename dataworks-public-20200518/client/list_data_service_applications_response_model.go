// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceApplicationsResponseBody) *ListDataServiceApplicationsResponse
	GetBody() *ListDataServiceApplicationsResponseBody
}

type ListDataServiceApplicationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceApplicationsResponse) GetBody() *ListDataServiceApplicationsResponseBody {
	return s.Body
}

func (s *ListDataServiceApplicationsResponse) SetHeaders(v map[string]*string) *ListDataServiceApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceApplicationsResponse) SetStatusCode(v int32) *ListDataServiceApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceApplicationsResponse) SetBody(v *ListDataServiceApplicationsResponseBody) *ListDataServiceApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
