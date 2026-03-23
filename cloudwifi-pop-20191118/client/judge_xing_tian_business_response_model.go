// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJudgeXingTianBusinessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JudgeXingTianBusinessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JudgeXingTianBusinessResponse
	GetStatusCode() *int32
	SetBody(v *JudgeXingTianBusinessResponseBody) *JudgeXingTianBusinessResponse
	GetBody() *JudgeXingTianBusinessResponseBody
}

type JudgeXingTianBusinessResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JudgeXingTianBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JudgeXingTianBusinessResponse) String() string {
	return dara.Prettify(s)
}

func (s JudgeXingTianBusinessResponse) GoString() string {
	return s.String()
}

func (s *JudgeXingTianBusinessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JudgeXingTianBusinessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JudgeXingTianBusinessResponse) GetBody() *JudgeXingTianBusinessResponseBody {
	return s.Body
}

func (s *JudgeXingTianBusinessResponse) SetHeaders(v map[string]*string) *JudgeXingTianBusinessResponse {
	s.Headers = v
	return s
}

func (s *JudgeXingTianBusinessResponse) SetStatusCode(v int32) *JudgeXingTianBusinessResponse {
	s.StatusCode = &v
	return s
}

func (s *JudgeXingTianBusinessResponse) SetBody(v *JudgeXingTianBusinessResponseBody) *JudgeXingTianBusinessResponse {
	s.Body = v
	return s
}

func (s *JudgeXingTianBusinessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
