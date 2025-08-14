// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNameListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNameListResponseBody) *DeleteNameListResponse
	GetBody() *DeleteNameListResponseBody
}

type DeleteNameListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNameListResponse) GoString() string {
	return s.String()
}

func (s *DeleteNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNameListResponse) GetBody() *DeleteNameListResponseBody {
	return s.Body
}

func (s *DeleteNameListResponse) SetHeaders(v map[string]*string) *DeleteNameListResponse {
	s.Headers = v
	return s
}

func (s *DeleteNameListResponse) SetStatusCode(v int32) *DeleteNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNameListResponse) SetBody(v *DeleteNameListResponseBody) *DeleteNameListResponse {
	s.Body = v
	return s
}

func (s *DeleteNameListResponse) Validate() error {
	return dara.Validate(s)
}
