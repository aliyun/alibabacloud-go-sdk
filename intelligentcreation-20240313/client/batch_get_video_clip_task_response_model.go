// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetVideoClipTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetVideoClipTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetVideoClipTaskResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetVideoClipTaskResponseBody) *BatchGetVideoClipTaskResponse
	GetBody() *BatchGetVideoClipTaskResponseBody
}

type BatchGetVideoClipTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetVideoClipTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetVideoClipTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoClipTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetVideoClipTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetVideoClipTaskResponse) GetBody() *BatchGetVideoClipTaskResponseBody {
	return s.Body
}

func (s *BatchGetVideoClipTaskResponse) SetHeaders(v map[string]*string) *BatchGetVideoClipTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchGetVideoClipTaskResponse) SetStatusCode(v int32) *BatchGetVideoClipTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetVideoClipTaskResponse) SetBody(v *BatchGetVideoClipTaskResponseBody) *BatchGetVideoClipTaskResponse {
	s.Body = v
	return s
}

func (s *BatchGetVideoClipTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
