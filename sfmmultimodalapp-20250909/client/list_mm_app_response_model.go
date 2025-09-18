// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmAppResponse
	GetStatusCode() *int32
	SetBody(v *ListMmAppResponseBody) *ListMmAppResponse
	GetBody() *ListMmAppResponseBody
}

type ListMmAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmAppResponse) GoString() string {
	return s.String()
}

func (s *ListMmAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmAppResponse) GetBody() *ListMmAppResponseBody {
	return s.Body
}

func (s *ListMmAppResponse) SetHeaders(v map[string]*string) *ListMmAppResponse {
	s.Headers = v
	return s
}

func (s *ListMmAppResponse) SetStatusCode(v int32) *ListMmAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmAppResponse) SetBody(v *ListMmAppResponseBody) *ListMmAppResponse {
	s.Body = v
	return s
}

func (s *ListMmAppResponse) Validate() error {
	return dara.Validate(s)
}
