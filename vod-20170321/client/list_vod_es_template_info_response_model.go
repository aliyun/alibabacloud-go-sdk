// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodEsTemplateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodEsTemplateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodEsTemplateInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListVodEsTemplateInfoResponseBody) *ListVodEsTemplateInfoResponse
	GetBody() *ListVodEsTemplateInfoResponseBody
}

type ListVodEsTemplateInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodEsTemplateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodEsTemplateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodEsTemplateInfoResponse) GoString() string {
	return s.String()
}

func (s *ListVodEsTemplateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodEsTemplateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodEsTemplateInfoResponse) GetBody() *ListVodEsTemplateInfoResponseBody {
	return s.Body
}

func (s *ListVodEsTemplateInfoResponse) SetHeaders(v map[string]*string) *ListVodEsTemplateInfoResponse {
	s.Headers = v
	return s
}

func (s *ListVodEsTemplateInfoResponse) SetStatusCode(v int32) *ListVodEsTemplateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodEsTemplateInfoResponse) SetBody(v *ListVodEsTemplateInfoResponseBody) *ListVodEsTemplateInfoResponse {
	s.Body = v
	return s
}

func (s *ListVodEsTemplateInfoResponse) Validate() error {
	return dara.Validate(s)
}
