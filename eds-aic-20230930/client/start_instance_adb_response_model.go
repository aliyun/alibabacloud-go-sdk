// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceAdbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartInstanceAdbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartInstanceAdbResponse
	GetStatusCode() *int32
	SetBody(v *StartInstanceAdbResponseBody) *StartInstanceAdbResponse
	GetBody() *StartInstanceAdbResponseBody
}

type StartInstanceAdbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstanceAdbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartInstanceAdbResponse) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceAdbResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceAdbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartInstanceAdbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartInstanceAdbResponse) GetBody() *StartInstanceAdbResponseBody {
	return s.Body
}

func (s *StartInstanceAdbResponse) SetHeaders(v map[string]*string) *StartInstanceAdbResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceAdbResponse) SetStatusCode(v int32) *StartInstanceAdbResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceAdbResponse) SetBody(v *StartInstanceAdbResponseBody) *StartInstanceAdbResponse {
	s.Body = v
	return s
}

func (s *StartInstanceAdbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
