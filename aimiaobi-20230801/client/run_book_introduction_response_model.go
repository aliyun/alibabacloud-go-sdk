// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookIntroductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunBookIntroductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunBookIntroductionResponse
	GetStatusCode() *int32
	SetBody(v *RunBookIntroductionResponseBody) *RunBookIntroductionResponse
	GetBody() *RunBookIntroductionResponseBody
}

type RunBookIntroductionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunBookIntroductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunBookIntroductionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunBookIntroductionResponse) GoString() string {
	return s.String()
}

func (s *RunBookIntroductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunBookIntroductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunBookIntroductionResponse) GetBody() *RunBookIntroductionResponseBody {
	return s.Body
}

func (s *RunBookIntroductionResponse) SetHeaders(v map[string]*string) *RunBookIntroductionResponse {
	s.Headers = v
	return s
}

func (s *RunBookIntroductionResponse) SetStatusCode(v int32) *RunBookIntroductionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunBookIntroductionResponse) SetBody(v *RunBookIntroductionResponseBody) *RunBookIntroductionResponse {
	s.Body = v
	return s
}

func (s *RunBookIntroductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
