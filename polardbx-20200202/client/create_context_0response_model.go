// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContext0Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContext0Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContext0Response
	GetStatusCode() *int32
	SetBody(v *CreateContext0ResponseBody) *CreateContext0Response
	GetBody() *CreateContext0ResponseBody
}

type CreateContext0Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContext0ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContext0Response) String() string {
	return dara.Prettify(s)
}

func (s CreateContext0Response) GoString() string {
	return s.String()
}

func (s *CreateContext0Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContext0Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContext0Response) GetBody() *CreateContext0ResponseBody {
	return s.Body
}

func (s *CreateContext0Response) SetHeaders(v map[string]*string) *CreateContext0Response {
	s.Headers = v
	return s
}

func (s *CreateContext0Response) SetStatusCode(v int32) *CreateContext0Response {
	s.StatusCode = &v
	return s
}

func (s *CreateContext0Response) SetBody(v *CreateContext0ResponseBody) *CreateContext0Response {
	s.Body = v
	return s
}

func (s *CreateContext0Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
