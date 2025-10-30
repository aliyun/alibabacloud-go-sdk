// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxb700Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAxb700Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAxb700Response
	GetStatusCode() *int32
	SetBody(v *BindAxb700ResponseBody) *BindAxb700Response
	GetBody() *BindAxb700ResponseBody
}

type BindAxb700Response struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxb700ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxb700Response) String() string {
	return dara.Prettify(s)
}

func (s BindAxb700Response) GoString() string {
	return s.String()
}

func (s *BindAxb700Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAxb700Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAxb700Response) GetBody() *BindAxb700ResponseBody {
	return s.Body
}

func (s *BindAxb700Response) SetHeaders(v map[string]*string) *BindAxb700Response {
	s.Headers = v
	return s
}

func (s *BindAxb700Response) SetStatusCode(v int32) *BindAxb700Response {
	s.StatusCode = &v
	return s
}

func (s *BindAxb700Response) SetBody(v *BindAxb700ResponseBody) *BindAxb700Response {
	s.Body = v
	return s
}

func (s *BindAxb700Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
