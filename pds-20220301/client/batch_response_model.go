// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchResponse
	GetStatusCode() *int32
	SetBody(v *BatchResponseBody) *BatchResponse
	GetBody() *BatchResponseBody
}

type BatchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchResponse) GoString() string {
	return s.String()
}

func (s *BatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchResponse) GetBody() *BatchResponseBody {
	return s.Body
}

func (s *BatchResponse) SetHeaders(v map[string]*string) *BatchResponse {
	s.Headers = v
	return s
}

func (s *BatchResponse) SetStatusCode(v int32) *BatchResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchResponse) SetBody(v *BatchResponseBody) *BatchResponse {
	s.Body = v
	return s
}

func (s *BatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
