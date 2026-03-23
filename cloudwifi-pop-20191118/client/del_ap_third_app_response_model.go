// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelApThirdAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelApThirdAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelApThirdAppResponse
	GetStatusCode() *int32
	SetBody(v *DelApThirdAppResponseBody) *DelApThirdAppResponse
	GetBody() *DelApThirdAppResponseBody
}

type DelApThirdAppResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelApThirdAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelApThirdAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DelApThirdAppResponse) GoString() string {
	return s.String()
}

func (s *DelApThirdAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelApThirdAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelApThirdAppResponse) GetBody() *DelApThirdAppResponseBody {
	return s.Body
}

func (s *DelApThirdAppResponse) SetHeaders(v map[string]*string) *DelApThirdAppResponse {
	s.Headers = v
	return s
}

func (s *DelApThirdAppResponse) SetStatusCode(v int32) *DelApThirdAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DelApThirdAppResponse) SetBody(v *DelApThirdAppResponseBody) *DelApThirdAppResponse {
	s.Body = v
	return s
}

func (s *DelApThirdAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
