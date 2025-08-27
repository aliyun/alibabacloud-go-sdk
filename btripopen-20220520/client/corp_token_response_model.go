// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCorpTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CorpTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CorpTokenResponse
	GetStatusCode() *int32
	SetBody(v *CorpTokenResponseBody) *CorpTokenResponse
	GetBody() *CorpTokenResponseBody
}

type CorpTokenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CorpTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CorpTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CorpTokenResponse) GoString() string {
	return s.String()
}

func (s *CorpTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CorpTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CorpTokenResponse) GetBody() *CorpTokenResponseBody {
	return s.Body
}

func (s *CorpTokenResponse) SetHeaders(v map[string]*string) *CorpTokenResponse {
	s.Headers = v
	return s
}

func (s *CorpTokenResponse) SetStatusCode(v int32) *CorpTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CorpTokenResponse) SetBody(v *CorpTokenResponseBody) *CorpTokenResponse {
	s.Body = v
	return s
}

func (s *CorpTokenResponse) Validate() error {
	return dara.Validate(s)
}
