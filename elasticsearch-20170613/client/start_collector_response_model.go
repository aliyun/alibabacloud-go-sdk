// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCollectorResponse
	GetStatusCode() *int32
	SetBody(v *StartCollectorResponseBody) *StartCollectorResponse
	GetBody() *StartCollectorResponseBody
}

type StartCollectorResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCollectorResponse) GoString() string {
	return s.String()
}

func (s *StartCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCollectorResponse) GetBody() *StartCollectorResponseBody {
	return s.Body
}

func (s *StartCollectorResponse) SetHeaders(v map[string]*string) *StartCollectorResponse {
	s.Headers = v
	return s
}

func (s *StartCollectorResponse) SetStatusCode(v int32) *StartCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCollectorResponse) SetBody(v *StartCollectorResponseBody) *StartCollectorResponse {
	s.Body = v
	return s
}

func (s *StartCollectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
