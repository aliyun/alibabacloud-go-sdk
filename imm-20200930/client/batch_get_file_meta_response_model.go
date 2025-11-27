// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetFileMetaResponseBody) *BatchGetFileMetaResponse
	GetBody() *BatchGetFileMetaResponseBody
}

type BatchGetFileMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFileMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetFileMetaResponse) GetBody() *BatchGetFileMetaResponseBody {
	return s.Body
}

func (s *BatchGetFileMetaResponse) SetHeaders(v map[string]*string) *BatchGetFileMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchGetFileMetaResponse) SetStatusCode(v int32) *BatchGetFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetFileMetaResponse) SetBody(v *BatchGetFileMetaResponseBody) *BatchGetFileMetaResponse {
	s.Body = v
	return s
}

func (s *BatchGetFileMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
