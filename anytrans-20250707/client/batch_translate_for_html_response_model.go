// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateForHtmlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchTranslateForHtmlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchTranslateForHtmlResponse
	GetStatusCode() *int32
	SetBody(v *BatchTranslateForHtmlResponseBody) *BatchTranslateForHtmlResponse
	GetBody() *BatchTranslateForHtmlResponseBody
}

type BatchTranslateForHtmlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchTranslateForHtmlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchTranslateForHtmlResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlResponse) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchTranslateForHtmlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchTranslateForHtmlResponse) GetBody() *BatchTranslateForHtmlResponseBody {
	return s.Body
}

func (s *BatchTranslateForHtmlResponse) SetHeaders(v map[string]*string) *BatchTranslateForHtmlResponse {
	s.Headers = v
	return s
}

func (s *BatchTranslateForHtmlResponse) SetStatusCode(v int32) *BatchTranslateForHtmlResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchTranslateForHtmlResponse) SetBody(v *BatchTranslateForHtmlResponseBody) *BatchTranslateForHtmlResponse {
	s.Body = v
	return s
}

func (s *BatchTranslateForHtmlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
