// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDataLakePartitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateDataLakePartitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateDataLakePartitionsResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateDataLakePartitionsResponseBody) *BatchCreateDataLakePartitionsResponse
	GetBody() *BatchCreateDataLakePartitionsResponseBody
}

type BatchCreateDataLakePartitionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateDataLakePartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateDataLakePartitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDataLakePartitionsResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateDataLakePartitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateDataLakePartitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateDataLakePartitionsResponse) GetBody() *BatchCreateDataLakePartitionsResponseBody {
	return s.Body
}

func (s *BatchCreateDataLakePartitionsResponse) SetHeaders(v map[string]*string) *BatchCreateDataLakePartitionsResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateDataLakePartitionsResponse) SetStatusCode(v int32) *BatchCreateDataLakePartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateDataLakePartitionsResponse) SetBody(v *BatchCreateDataLakePartitionsResponseBody) *BatchCreateDataLakePartitionsResponse {
	s.Body = v
	return s
}

func (s *BatchCreateDataLakePartitionsResponse) Validate() error {
	return dara.Validate(s)
}
