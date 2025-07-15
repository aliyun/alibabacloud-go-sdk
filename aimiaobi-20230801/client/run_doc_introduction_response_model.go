// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocIntroductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocIntroductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocIntroductionResponse
	GetStatusCode() *int32
	SetBody(v *RunDocIntroductionResponseBody) *RunDocIntroductionResponse
	GetBody() *RunDocIntroductionResponseBody
}

type RunDocIntroductionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunDocIntroductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocIntroductionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponse) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocIntroductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocIntroductionResponse) GetBody() *RunDocIntroductionResponseBody {
	return s.Body
}

func (s *RunDocIntroductionResponse) SetHeaders(v map[string]*string) *RunDocIntroductionResponse {
	s.Headers = v
	return s
}

func (s *RunDocIntroductionResponse) SetStatusCode(v int32) *RunDocIntroductionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocIntroductionResponse) SetBody(v *RunDocIntroductionResponseBody) *RunDocIntroductionResponse {
	s.Body = v
	return s
}

func (s *RunDocIntroductionResponse) Validate() error {
	return dara.Validate(s)
}
