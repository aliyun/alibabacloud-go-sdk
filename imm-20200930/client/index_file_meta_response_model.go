// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IndexFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IndexFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *IndexFileMetaResponseBody) *IndexFileMetaResponse
	GetBody() *IndexFileMetaResponseBody
}

type IndexFileMetaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IndexFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IndexFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s IndexFileMetaResponse) GoString() string {
	return s.String()
}

func (s *IndexFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IndexFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IndexFileMetaResponse) GetBody() *IndexFileMetaResponseBody {
	return s.Body
}

func (s *IndexFileMetaResponse) SetHeaders(v map[string]*string) *IndexFileMetaResponse {
	s.Headers = v
	return s
}

func (s *IndexFileMetaResponse) SetStatusCode(v int32) *IndexFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *IndexFileMetaResponse) SetBody(v *IndexFileMetaResponseBody) *IndexFileMetaResponse {
	s.Body = v
	return s
}

func (s *IndexFileMetaResponse) Validate() error {
	return dara.Validate(s)
}
