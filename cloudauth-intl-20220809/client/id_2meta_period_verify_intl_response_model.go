// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaPeriodVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id2MetaPeriodVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id2MetaPeriodVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *Id2MetaPeriodVerifyIntlResponseBody) *Id2MetaPeriodVerifyIntlResponse
	GetBody() *Id2MetaPeriodVerifyIntlResponseBody
}

type Id2MetaPeriodVerifyIntlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaPeriodVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaPeriodVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaPeriodVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id2MetaPeriodVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id2MetaPeriodVerifyIntlResponse) GetBody() *Id2MetaPeriodVerifyIntlResponseBody {
	return s.Body
}

func (s *Id2MetaPeriodVerifyIntlResponse) SetHeaders(v map[string]*string) *Id2MetaPeriodVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponse) SetStatusCode(v int32) *Id2MetaPeriodVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponse) SetBody(v *Id2MetaPeriodVerifyIntlResponseBody) *Id2MetaPeriodVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
