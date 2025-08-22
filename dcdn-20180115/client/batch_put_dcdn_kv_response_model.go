// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutDcdnKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchPutDcdnKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchPutDcdnKvResponse
	GetStatusCode() *int32
	SetBody(v *BatchPutDcdnKvResponseBody) *BatchPutDcdnKvResponse
	GetBody() *BatchPutDcdnKvResponseBody
}

type BatchPutDcdnKvResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchPutDcdnKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchPutDcdnKvResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvResponse) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchPutDcdnKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchPutDcdnKvResponse) GetBody() *BatchPutDcdnKvResponseBody {
	return s.Body
}

func (s *BatchPutDcdnKvResponse) SetHeaders(v map[string]*string) *BatchPutDcdnKvResponse {
	s.Headers = v
	return s
}

func (s *BatchPutDcdnKvResponse) SetStatusCode(v int32) *BatchPutDcdnKvResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchPutDcdnKvResponse) SetBody(v *BatchPutDcdnKvResponseBody) *BatchPutDcdnKvResponse {
	s.Body = v
	return s
}

func (s *BatchPutDcdnKvResponse) Validate() error {
	return dara.Validate(s)
}
