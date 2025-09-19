// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationProcessDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationProcessDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationProcessDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationProcessDetailResponseBody) *ListOperationProcessDetailResponse
	GetBody() *ListOperationProcessDetailResponseBody
}

type ListOperationProcessDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationProcessDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationProcessDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessDetailResponse) GoString() string {
	return s.String()
}

func (s *ListOperationProcessDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationProcessDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationProcessDetailResponse) GetBody() *ListOperationProcessDetailResponseBody {
	return s.Body
}

func (s *ListOperationProcessDetailResponse) SetHeaders(v map[string]*string) *ListOperationProcessDetailResponse {
	s.Headers = v
	return s
}

func (s *ListOperationProcessDetailResponse) SetStatusCode(v int32) *ListOperationProcessDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationProcessDetailResponse) SetBody(v *ListOperationProcessDetailResponseBody) *ListOperationProcessDetailResponse {
	s.Body = v
	return s
}

func (s *ListOperationProcessDetailResponse) Validate() error {
	return dara.Validate(s)
}
