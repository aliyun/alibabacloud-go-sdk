// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubs700Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSubs700Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSubs700Response
	GetStatusCode() *int32
	SetBody(v *UpdateSubs700ResponseBody) *UpdateSubs700Response
	GetBody() *UpdateSubs700ResponseBody
}

type UpdateSubs700Response struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubs700ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubs700Response) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubs700Response) GoString() string {
	return s.String()
}

func (s *UpdateSubs700Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSubs700Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSubs700Response) GetBody() *UpdateSubs700ResponseBody {
	return s.Body
}

func (s *UpdateSubs700Response) SetHeaders(v map[string]*string) *UpdateSubs700Response {
	s.Headers = v
	return s
}

func (s *UpdateSubs700Response) SetStatusCode(v int32) *UpdateSubs700Response {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubs700Response) SetBody(v *UpdateSubs700ResponseBody) *UpdateSubs700Response {
	s.Body = v
	return s
}

func (s *UpdateSubs700Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
