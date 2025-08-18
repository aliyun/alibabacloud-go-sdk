// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchPutKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchPutKvResponse
	GetStatusCode() *int32
	SetBody(v *BatchPutKvResponseBody) *BatchPutKvResponse
	GetBody() *BatchPutKvResponseBody
}

type BatchPutKvResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchPutKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchPutKvResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvResponse) GoString() string {
	return s.String()
}

func (s *BatchPutKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchPutKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchPutKvResponse) GetBody() *BatchPutKvResponseBody {
	return s.Body
}

func (s *BatchPutKvResponse) SetHeaders(v map[string]*string) *BatchPutKvResponse {
	s.Headers = v
	return s
}

func (s *BatchPutKvResponse) SetStatusCode(v int32) *BatchPutKvResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchPutKvResponse) SetBody(v *BatchPutKvResponseBody) *BatchPutKvResponse {
	s.Body = v
	return s
}

func (s *BatchPutKvResponse) Validate() error {
	return dara.Validate(s)
}
