// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteKvResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteKvResponseBody) *BatchDeleteKvResponse
	GetBody() *BatchDeleteKvResponseBody
}

type BatchDeleteKvResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteKvResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteKvResponse) GetBody() *BatchDeleteKvResponseBody {
	return s.Body
}

func (s *BatchDeleteKvResponse) SetHeaders(v map[string]*string) *BatchDeleteKvResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteKvResponse) SetStatusCode(v int32) *BatchDeleteKvResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteKvResponse) SetBody(v *BatchDeleteKvResponseBody) *BatchDeleteKvResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteKvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
