// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMem0Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMem0Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMem0Response
	GetStatusCode() *int32
	SetBody(v *CreateMem0ResponseBody) *CreateMem0Response
	GetBody() *CreateMem0ResponseBody
}

type CreateMem0Response struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMem0ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMem0Response) String() string {
	return dara.Prettify(s)
}

func (s CreateMem0Response) GoString() string {
	return s.String()
}

func (s *CreateMem0Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMem0Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMem0Response) GetBody() *CreateMem0ResponseBody {
	return s.Body
}

func (s *CreateMem0Response) SetHeaders(v map[string]*string) *CreateMem0Response {
	s.Headers = v
	return s
}

func (s *CreateMem0Response) SetStatusCode(v int32) *CreateMem0Response {
	s.StatusCode = &v
	return s
}

func (s *CreateMem0Response) SetBody(v *CreateMem0ResponseBody) *CreateMem0Response {
	s.Body = v
	return s
}

func (s *CreateMem0Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
