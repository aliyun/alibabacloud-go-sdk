// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediaInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetMediaInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetMediaInfosResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetMediaInfosResponseBody) *BatchGetMediaInfosResponse
	GetBody() *BatchGetMediaInfosResponseBody
}

type BatchGetMediaInfosResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetMediaInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetMediaInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetMediaInfosResponse) GetBody() *BatchGetMediaInfosResponseBody {
	return s.Body
}

func (s *BatchGetMediaInfosResponse) SetHeaders(v map[string]*string) *BatchGetMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *BatchGetMediaInfosResponse) SetStatusCode(v int32) *BatchGetMediaInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetMediaInfosResponse) SetBody(v *BatchGetMediaInfosResponseBody) *BatchGetMediaInfosResponse {
	s.Body = v
	return s
}

func (s *BatchGetMediaInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
