// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIPv6Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIPv6Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIPv6Response
	GetStatusCode() *int32
	SetBody(v *GetIPv6ResponseBody) *GetIPv6Response
	GetBody() *GetIPv6ResponseBody
}

type GetIPv6Response struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIPv6ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIPv6Response) String() string {
	return dara.Prettify(s)
}

func (s GetIPv6Response) GoString() string {
	return s.String()
}

func (s *GetIPv6Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIPv6Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIPv6Response) GetBody() *GetIPv6ResponseBody {
	return s.Body
}

func (s *GetIPv6Response) SetHeaders(v map[string]*string) *GetIPv6Response {
	s.Headers = v
	return s
}

func (s *GetIPv6Response) SetStatusCode(v int32) *GetIPv6Response {
	s.StatusCode = &v
	return s
}

func (s *GetIPv6Response) SetBody(v *GetIPv6ResponseBody) *GetIPv6Response {
	s.Body = v
	return s
}

func (s *GetIPv6Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
