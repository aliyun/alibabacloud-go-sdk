// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteFileMetaResponseBody) *BatchDeleteFileMetaResponse
	GetBody() *BatchDeleteFileMetaResponseBody
}

type BatchDeleteFileMetaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteFileMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteFileMetaResponse) GetBody() *BatchDeleteFileMetaResponseBody {
	return s.Body
}

func (s *BatchDeleteFileMetaResponse) SetHeaders(v map[string]*string) *BatchDeleteFileMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteFileMetaResponse) SetStatusCode(v int32) *BatchDeleteFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteFileMetaResponse) SetBody(v *BatchDeleteFileMetaResponseBody) *BatchDeleteFileMetaResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteFileMetaResponse) Validate() error {
	return dara.Validate(s)
}
