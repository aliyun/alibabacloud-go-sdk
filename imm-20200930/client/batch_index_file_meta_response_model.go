// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchIndexFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchIndexFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchIndexFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *BatchIndexFileMetaResponseBody) *BatchIndexFileMetaResponse
	GetBody() *BatchIndexFileMetaResponseBody
}

type BatchIndexFileMetaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchIndexFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchIndexFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchIndexFileMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchIndexFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchIndexFileMetaResponse) GetBody() *BatchIndexFileMetaResponseBody {
	return s.Body
}

func (s *BatchIndexFileMetaResponse) SetHeaders(v map[string]*string) *BatchIndexFileMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchIndexFileMetaResponse) SetStatusCode(v int32) *BatchIndexFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchIndexFileMetaResponse) SetBody(v *BatchIndexFileMetaResponseBody) *BatchIndexFileMetaResponse {
	s.Body = v
	return s
}

func (s *BatchIndexFileMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
