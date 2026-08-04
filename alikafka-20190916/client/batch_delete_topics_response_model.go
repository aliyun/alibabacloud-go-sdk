// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteTopicsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteTopicsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteTopicsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteTopicsResponseBody) *BatchDeleteTopicsResponse
	GetBody() *BatchDeleteTopicsResponseBody
}

type BatchDeleteTopicsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteTopicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteTopicsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteTopicsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteTopicsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteTopicsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteTopicsResponse) GetBody() *BatchDeleteTopicsResponseBody {
	return s.Body
}

func (s *BatchDeleteTopicsResponse) SetHeaders(v map[string]*string) *BatchDeleteTopicsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteTopicsResponse) SetStatusCode(v int32) *BatchDeleteTopicsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteTopicsResponse) SetBody(v *BatchDeleteTopicsResponseBody) *BatchDeleteTopicsResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteTopicsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
