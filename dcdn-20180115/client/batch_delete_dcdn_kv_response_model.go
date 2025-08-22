// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteDcdnKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteDcdnKvResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteDcdnKvResponseBody) *BatchDeleteDcdnKvResponse
	GetBody() *BatchDeleteDcdnKvResponseBody
}

type BatchDeleteDcdnKvResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteDcdnKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteDcdnKvResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnKvResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteDcdnKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteDcdnKvResponse) GetBody() *BatchDeleteDcdnKvResponseBody {
	return s.Body
}

func (s *BatchDeleteDcdnKvResponse) SetHeaders(v map[string]*string) *BatchDeleteDcdnKvResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDcdnKvResponse) SetStatusCode(v int32) *BatchDeleteDcdnKvResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteDcdnKvResponse) SetBody(v *BatchDeleteDcdnKvResponseBody) *BatchDeleteDcdnKvResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteDcdnKvResponse) Validate() error {
	return dara.Validate(s)
}
