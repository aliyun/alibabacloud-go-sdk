// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartApplicationResponse
	GetStatusCode() *int32
	SetBody(v *RestartApplicationResponseBody) *RestartApplicationResponse
	GetBody() *RestartApplicationResponseBody
}

type RestartApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartApplicationResponse) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartApplicationResponse) GetBody() *RestartApplicationResponseBody {
	return s.Body
}

func (s *RestartApplicationResponse) SetHeaders(v map[string]*string) *RestartApplicationResponse {
	s.Headers = v
	return s
}

func (s *RestartApplicationResponse) SetStatusCode(v int32) *RestartApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartApplicationResponse) SetBody(v *RestartApplicationResponseBody) *RestartApplicationResponse {
	s.Body = v
	return s
}

func (s *RestartApplicationResponse) Validate() error {
	return dara.Validate(s)
}
