// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigItemsResponseBody) *ListConfigItemsResponse
	GetBody() *ListConfigItemsResponseBody
}

type ListConfigItemsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigItemsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigItemsResponse) GetBody() *ListConfigItemsResponseBody {
	return s.Body
}

func (s *ListConfigItemsResponse) SetHeaders(v map[string]*string) *ListConfigItemsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigItemsResponse) SetStatusCode(v int32) *ListConfigItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigItemsResponse) SetBody(v *ListConfigItemsResponseBody) *ListConfigItemsResponse {
	s.Body = v
	return s
}

func (s *ListConfigItemsResponse) Validate() error {
	return dara.Validate(s)
}
