// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AiSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AiSearchResponse
	GetStatusCode() *int32
	SetBody(v *AiSearchResponseBody) *AiSearchResponse
	GetBody() *AiSearchResponseBody
}

type AiSearchResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AiSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AiSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s AiSearchResponse) GoString() string {
	return s.String()
}

func (s *AiSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AiSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AiSearchResponse) GetBody() *AiSearchResponseBody {
	return s.Body
}

func (s *AiSearchResponse) SetHeaders(v map[string]*string) *AiSearchResponse {
	s.Headers = v
	return s
}

func (s *AiSearchResponse) SetStatusCode(v int32) *AiSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *AiSearchResponse) SetBody(v *AiSearchResponseBody) *AiSearchResponse {
	s.Body = v
	return s
}

func (s *AiSearchResponse) Validate() error {
	return dara.Validate(s)
}
