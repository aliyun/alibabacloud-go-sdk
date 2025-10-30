// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindGxb700Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindGxb700Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindGxb700Response
	GetStatusCode() *int32
	SetBody(v *BindGxb700ResponseBody) *BindGxb700Response
	GetBody() *BindGxb700ResponseBody
}

type BindGxb700Response struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindGxb700ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindGxb700Response) String() string {
	return dara.Prettify(s)
}

func (s BindGxb700Response) GoString() string {
	return s.String()
}

func (s *BindGxb700Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindGxb700Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindGxb700Response) GetBody() *BindGxb700ResponseBody {
	return s.Body
}

func (s *BindGxb700Response) SetHeaders(v map[string]*string) *BindGxb700Response {
	s.Headers = v
	return s
}

func (s *BindGxb700Response) SetStatusCode(v int32) *BindGxb700Response {
	s.StatusCode = &v
	return s
}

func (s *BindGxb700Response) SetBody(v *BindGxb700ResponseBody) *BindGxb700Response {
	s.Body = v
	return s
}

func (s *BindGxb700Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
