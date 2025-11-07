// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaPeriodVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id2MetaPeriodVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id2MetaPeriodVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Id2MetaPeriodVerifyResponseBody) *Id2MetaPeriodVerifyResponse
	GetBody() *Id2MetaPeriodVerifyResponseBody
}

type Id2MetaPeriodVerifyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaPeriodVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaPeriodVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaPeriodVerifyResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id2MetaPeriodVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id2MetaPeriodVerifyResponse) GetBody() *Id2MetaPeriodVerifyResponseBody {
	return s.Body
}

func (s *Id2MetaPeriodVerifyResponse) SetHeaders(v map[string]*string) *Id2MetaPeriodVerifyResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaPeriodVerifyResponse) SetStatusCode(v int32) *Id2MetaPeriodVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyResponse) SetBody(v *Id2MetaPeriodVerifyResponseBody) *Id2MetaPeriodVerifyResponse {
	s.Body = v
	return s
}

func (s *Id2MetaPeriodVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
