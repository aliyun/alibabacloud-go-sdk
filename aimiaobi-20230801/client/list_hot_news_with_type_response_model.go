// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotNewsWithTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotNewsWithTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotNewsWithTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListHotNewsWithTypeResponseBody) *ListHotNewsWithTypeResponse
	GetBody() *ListHotNewsWithTypeResponseBody
}

type ListHotNewsWithTypeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotNewsWithTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotNewsWithTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotNewsWithTypeResponse) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotNewsWithTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotNewsWithTypeResponse) GetBody() *ListHotNewsWithTypeResponseBody {
	return s.Body
}

func (s *ListHotNewsWithTypeResponse) SetHeaders(v map[string]*string) *ListHotNewsWithTypeResponse {
	s.Headers = v
	return s
}

func (s *ListHotNewsWithTypeResponse) SetStatusCode(v int32) *ListHotNewsWithTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotNewsWithTypeResponse) SetBody(v *ListHotNewsWithTypeResponseBody) *ListHotNewsWithTypeResponse {
	s.Body = v
	return s
}

func (s *ListHotNewsWithTypeResponse) Validate() error {
	return dara.Validate(s)
}
