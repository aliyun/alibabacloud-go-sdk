// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSubs700Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindSubs700Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindSubs700Response
	GetStatusCode() *int32
	SetBody(v *UnbindSubs700ResponseBody) *UnbindSubs700Response
	GetBody() *UnbindSubs700ResponseBody
}

type UnbindSubs700Response struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindSubs700ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindSubs700Response) String() string {
	return dara.Prettify(s)
}

func (s UnbindSubs700Response) GoString() string {
	return s.String()
}

func (s *UnbindSubs700Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindSubs700Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindSubs700Response) GetBody() *UnbindSubs700ResponseBody {
	return s.Body
}

func (s *UnbindSubs700Response) SetHeaders(v map[string]*string) *UnbindSubs700Response {
	s.Headers = v
	return s
}

func (s *UnbindSubs700Response) SetStatusCode(v int32) *UnbindSubs700Response {
	s.StatusCode = &v
	return s
}

func (s *UnbindSubs700Response) SetBody(v *UnbindSubs700ResponseBody) *UnbindSubs700Response {
	s.Body = v
	return s
}

func (s *UnbindSubs700Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
