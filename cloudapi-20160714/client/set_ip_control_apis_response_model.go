// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIpControlApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetIpControlApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetIpControlApisResponse
	GetStatusCode() *int32
	SetBody(v *SetIpControlApisResponseBody) *SetIpControlApisResponse
	GetBody() *SetIpControlApisResponseBody
}

type SetIpControlApisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetIpControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetIpControlApisResponse) String() string {
	return dara.Prettify(s)
}

func (s SetIpControlApisResponse) GoString() string {
	return s.String()
}

func (s *SetIpControlApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetIpControlApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetIpControlApisResponse) GetBody() *SetIpControlApisResponseBody {
	return s.Body
}

func (s *SetIpControlApisResponse) SetHeaders(v map[string]*string) *SetIpControlApisResponse {
	s.Headers = v
	return s
}

func (s *SetIpControlApisResponse) SetStatusCode(v int32) *SetIpControlApisResponse {
	s.StatusCode = &v
	return s
}

func (s *SetIpControlApisResponse) SetBody(v *SetIpControlApisResponseBody) *SetIpControlApisResponse {
	s.Body = v
	return s
}

func (s *SetIpControlApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
