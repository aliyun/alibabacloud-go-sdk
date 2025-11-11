// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContinueContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunContinueContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunContinueContentResponse
	GetStatusCode() *int32
	SetBody(v *RunContinueContentResponseBody) *RunContinueContentResponse
	GetBody() *RunContinueContentResponseBody
}

type RunContinueContentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunContinueContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunContinueContentResponse) String() string {
	return dara.Prettify(s)
}

func (s RunContinueContentResponse) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunContinueContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunContinueContentResponse) GetBody() *RunContinueContentResponseBody {
	return s.Body
}

func (s *RunContinueContentResponse) SetHeaders(v map[string]*string) *RunContinueContentResponse {
	s.Headers = v
	return s
}

func (s *RunContinueContentResponse) SetStatusCode(v int32) *RunContinueContentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunContinueContentResponse) SetBody(v *RunContinueContentResponseBody) *RunContinueContentResponse {
	s.Body = v
	return s
}

func (s *RunContinueContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
