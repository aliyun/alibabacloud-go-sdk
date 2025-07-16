// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateImageBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateImageBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateImageBatchResponse
	GetStatusCode() *int32
	SetBody(v *TranslateImageBatchResponseBody) *TranslateImageBatchResponse
	GetBody() *TranslateImageBatchResponseBody
}

type TranslateImageBatchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateImageBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateImageBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateImageBatchResponse) GoString() string {
	return s.String()
}

func (s *TranslateImageBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateImageBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateImageBatchResponse) GetBody() *TranslateImageBatchResponseBody {
	return s.Body
}

func (s *TranslateImageBatchResponse) SetHeaders(v map[string]*string) *TranslateImageBatchResponse {
	s.Headers = v
	return s
}

func (s *TranslateImageBatchResponse) SetStatusCode(v int32) *TranslateImageBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateImageBatchResponse) SetBody(v *TranslateImageBatchResponseBody) *TranslateImageBatchResponse {
	s.Body = v
	return s
}

func (s *TranslateImageBatchResponse) Validate() error {
	return dara.Validate(s)
}
