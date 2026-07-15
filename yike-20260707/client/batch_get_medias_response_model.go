// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetMediasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetMediasResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetMediasResponseBody) *BatchGetMediasResponse
	GetBody() *BatchGetMediasResponseBody
}

type BatchGetMediasResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetMediasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetMediasResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponse) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetMediasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetMediasResponse) GetBody() *BatchGetMediasResponseBody {
	return s.Body
}

func (s *BatchGetMediasResponse) SetHeaders(v map[string]*string) *BatchGetMediasResponse {
	s.Headers = v
	return s
}

func (s *BatchGetMediasResponse) SetStatusCode(v int32) *BatchGetMediasResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetMediasResponse) SetBody(v *BatchGetMediasResponseBody) *BatchGetMediasResponse {
	s.Body = v
	return s
}

func (s *BatchGetMediasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
