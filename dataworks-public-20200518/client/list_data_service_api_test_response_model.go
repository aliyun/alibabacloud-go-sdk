// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceApiTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceApiTestResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceApiTestResponseBody) *ListDataServiceApiTestResponse
	GetBody() *ListDataServiceApiTestResponseBody
}

type ListDataServiceApiTestResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceApiTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceApiTestResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiTestResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceApiTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceApiTestResponse) GetBody() *ListDataServiceApiTestResponseBody {
	return s.Body
}

func (s *ListDataServiceApiTestResponse) SetHeaders(v map[string]*string) *ListDataServiceApiTestResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceApiTestResponse) SetStatusCode(v int32) *ListDataServiceApiTestResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceApiTestResponse) SetBody(v *ListDataServiceApiTestResponseBody) *ListDataServiceApiTestResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceApiTestResponse) Validate() error {
	return dara.Validate(s)
}
