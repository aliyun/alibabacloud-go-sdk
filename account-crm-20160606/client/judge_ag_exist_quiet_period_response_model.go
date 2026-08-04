// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJudgeAgExistQuietPeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JudgeAgExistQuietPeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JudgeAgExistQuietPeriodResponse
	GetStatusCode() *int32
	SetBody(v *JudgeAgExistQuietPeriodResponseBody) *JudgeAgExistQuietPeriodResponse
	GetBody() *JudgeAgExistQuietPeriodResponseBody
}

type JudgeAgExistQuietPeriodResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JudgeAgExistQuietPeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JudgeAgExistQuietPeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s JudgeAgExistQuietPeriodResponse) GoString() string {
	return s.String()
}

func (s *JudgeAgExistQuietPeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JudgeAgExistQuietPeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JudgeAgExistQuietPeriodResponse) GetBody() *JudgeAgExistQuietPeriodResponseBody {
	return s.Body
}

func (s *JudgeAgExistQuietPeriodResponse) SetHeaders(v map[string]*string) *JudgeAgExistQuietPeriodResponse {
	s.Headers = v
	return s
}

func (s *JudgeAgExistQuietPeriodResponse) SetStatusCode(v int32) *JudgeAgExistQuietPeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *JudgeAgExistQuietPeriodResponse) SetBody(v *JudgeAgExistQuietPeriodResponseBody) *JudgeAgExistQuietPeriodResponse {
	s.Body = v
	return s
}

func (s *JudgeAgExistQuietPeriodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
