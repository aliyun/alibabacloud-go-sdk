// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAppCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAppCodeResponse
	GetStatusCode() *int32
	SetBody(v *ResetAppCodeResponseBody) *ResetAppCodeResponse
	GetBody() *ResetAppCodeResponseBody
}

type ResetAppCodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAppCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAppCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAppCodeResponse) GoString() string {
	return s.String()
}

func (s *ResetAppCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAppCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAppCodeResponse) GetBody() *ResetAppCodeResponseBody {
	return s.Body
}

func (s *ResetAppCodeResponse) SetHeaders(v map[string]*string) *ResetAppCodeResponse {
	s.Headers = v
	return s
}

func (s *ResetAppCodeResponse) SetStatusCode(v int32) *ResetAppCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAppCodeResponse) SetBody(v *ResetAppCodeResponseBody) *ResetAppCodeResponse {
	s.Body = v
	return s
}

func (s *ResetAppCodeResponse) Validate() error {
	return dara.Validate(s)
}
