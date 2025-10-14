// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIPv6Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIPv6Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIPv6Response
	GetStatusCode() *int32
	SetBody(v *UpdateIPv6ResponseBody) *UpdateIPv6Response
	GetBody() *UpdateIPv6ResponseBody
}

type UpdateIPv6Response struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIPv6ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIPv6Response) String() string {
	return dara.Prettify(s)
}

func (s UpdateIPv6Response) GoString() string {
	return s.String()
}

func (s *UpdateIPv6Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIPv6Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIPv6Response) GetBody() *UpdateIPv6ResponseBody {
	return s.Body
}

func (s *UpdateIPv6Response) SetHeaders(v map[string]*string) *UpdateIPv6Response {
	s.Headers = v
	return s
}

func (s *UpdateIPv6Response) SetStatusCode(v int32) *UpdateIPv6Response {
	s.StatusCode = &v
	return s
}

func (s *UpdateIPv6Response) SetBody(v *UpdateIPv6ResponseBody) *UpdateIPv6Response {
	s.Body = v
	return s
}

func (s *UpdateIPv6Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
