// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupCorpTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GroupCorpTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GroupCorpTokenResponse
	GetStatusCode() *int32
	SetBody(v *GroupCorpTokenResponseBody) *GroupCorpTokenResponse
	GetBody() *GroupCorpTokenResponseBody
}

type GroupCorpTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupCorpTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupCorpTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GroupCorpTokenResponse) GoString() string {
	return s.String()
}

func (s *GroupCorpTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GroupCorpTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GroupCorpTokenResponse) GetBody() *GroupCorpTokenResponseBody {
	return s.Body
}

func (s *GroupCorpTokenResponse) SetHeaders(v map[string]*string) *GroupCorpTokenResponse {
	s.Headers = v
	return s
}

func (s *GroupCorpTokenResponse) SetStatusCode(v int32) *GroupCorpTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupCorpTokenResponse) SetBody(v *GroupCorpTokenResponseBody) *GroupCorpTokenResponse {
	s.Body = v
	return s
}

func (s *GroupCorpTokenResponse) Validate() error {
	return dara.Validate(s)
}
