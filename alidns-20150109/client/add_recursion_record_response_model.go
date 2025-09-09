// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecursionRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRecursionRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRecursionRecordResponse
	GetStatusCode() *int32
	SetBody(v *AddRecursionRecordResponseBody) *AddRecursionRecordResponse
	GetBody() *AddRecursionRecordResponseBody
}

type AddRecursionRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecursionRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecursionRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRecursionRecordResponse) GoString() string {
	return s.String()
}

func (s *AddRecursionRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRecursionRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRecursionRecordResponse) GetBody() *AddRecursionRecordResponseBody {
	return s.Body
}

func (s *AddRecursionRecordResponse) SetHeaders(v map[string]*string) *AddRecursionRecordResponse {
	s.Headers = v
	return s
}

func (s *AddRecursionRecordResponse) SetStatusCode(v int32) *AddRecursionRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecursionRecordResponse) SetBody(v *AddRecursionRecordResponseBody) *AddRecursionRecordResponse {
	s.Body = v
	return s
}

func (s *AddRecursionRecordResponse) Validate() error {
	return dara.Validate(s)
}
