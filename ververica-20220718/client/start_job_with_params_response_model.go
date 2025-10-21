// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobWithParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartJobWithParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartJobWithParamsResponse
	GetStatusCode() *int32
	SetBody(v *StartJobWithParamsResponseBody) *StartJobWithParamsResponse
	GetBody() *StartJobWithParamsResponseBody
}

type StartJobWithParamsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartJobWithParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobWithParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s StartJobWithParamsResponse) GoString() string {
	return s.String()
}

func (s *StartJobWithParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartJobWithParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartJobWithParamsResponse) GetBody() *StartJobWithParamsResponseBody {
	return s.Body
}

func (s *StartJobWithParamsResponse) SetHeaders(v map[string]*string) *StartJobWithParamsResponse {
	s.Headers = v
	return s
}

func (s *StartJobWithParamsResponse) SetStatusCode(v int32) *StartJobWithParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *StartJobWithParamsResponse) SetBody(v *StartJobWithParamsResponseBody) *StartJobWithParamsResponse {
	s.Body = v
	return s
}

func (s *StartJobWithParamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
