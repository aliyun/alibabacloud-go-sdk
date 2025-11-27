// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateFileMetaResponseBody) *BatchUpdateFileMetaResponse
	GetBody() *BatchUpdateFileMetaResponseBody
}

type BatchUpdateFileMetaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateFileMetaResponse) GetBody() *BatchUpdateFileMetaResponseBody {
	return s.Body
}

func (s *BatchUpdateFileMetaResponse) SetHeaders(v map[string]*string) *BatchUpdateFileMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateFileMetaResponse) SetStatusCode(v int32) *BatchUpdateFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateFileMetaResponse) SetBody(v *BatchUpdateFileMetaResponseBody) *BatchUpdateFileMetaResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateFileMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
