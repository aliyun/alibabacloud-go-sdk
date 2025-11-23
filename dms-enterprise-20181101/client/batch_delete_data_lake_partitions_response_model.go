// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDataLakePartitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteDataLakePartitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteDataLakePartitionsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteDataLakePartitionsResponseBody) *BatchDeleteDataLakePartitionsResponse
	GetBody() *BatchDeleteDataLakePartitionsResponseBody
}

type BatchDeleteDataLakePartitionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteDataLakePartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteDataLakePartitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDataLakePartitionsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDataLakePartitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteDataLakePartitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteDataLakePartitionsResponse) GetBody() *BatchDeleteDataLakePartitionsResponseBody {
	return s.Body
}

func (s *BatchDeleteDataLakePartitionsResponse) SetHeaders(v map[string]*string) *BatchDeleteDataLakePartitionsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDataLakePartitionsResponse) SetStatusCode(v int32) *BatchDeleteDataLakePartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsResponse) SetBody(v *BatchDeleteDataLakePartitionsResponseBody) *BatchDeleteDataLakePartitionsResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteDataLakePartitionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
