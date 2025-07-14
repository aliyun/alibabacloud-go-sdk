// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryShareListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryShareListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryShareListResponse
	GetStatusCode() *int32
	SetBody(v *QueryShareListResponseBody) *QueryShareListResponse
	GetBody() *QueryShareListResponseBody
}

type QueryShareListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryShareListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryShareListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryShareListResponse) GoString() string {
	return s.String()
}

func (s *QueryShareListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryShareListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryShareListResponse) GetBody() *QueryShareListResponseBody {
	return s.Body
}

func (s *QueryShareListResponse) SetHeaders(v map[string]*string) *QueryShareListResponse {
	s.Headers = v
	return s
}

func (s *QueryShareListResponse) SetStatusCode(v int32) *QueryShareListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryShareListResponse) SetBody(v *QueryShareListResponseBody) *QueryShareListResponse {
	s.Body = v
	return s
}

func (s *QueryShareListResponse) Validate() error {
	return dara.Validate(s)
}
