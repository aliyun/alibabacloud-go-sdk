// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateDataLakePartitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateDataLakePartitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateDataLakePartitionsResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateDataLakePartitionsResponseBody) *BatchUpdateDataLakePartitionsResponse
	GetBody() *BatchUpdateDataLakePartitionsResponseBody
}

type BatchUpdateDataLakePartitionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateDataLakePartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateDataLakePartitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateDataLakePartitionsResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateDataLakePartitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateDataLakePartitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateDataLakePartitionsResponse) GetBody() *BatchUpdateDataLakePartitionsResponseBody {
	return s.Body
}

func (s *BatchUpdateDataLakePartitionsResponse) SetHeaders(v map[string]*string) *BatchUpdateDataLakePartitionsResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateDataLakePartitionsResponse) SetStatusCode(v int32) *BatchUpdateDataLakePartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsResponse) SetBody(v *BatchUpdateDataLakePartitionsResponseBody) *BatchUpdateDataLakePartitionsResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateDataLakePartitionsResponse) Validate() error {
	return dara.Validate(s)
}
