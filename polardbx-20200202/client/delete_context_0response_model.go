// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContext0Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContext0Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContext0Response
	GetStatusCode() *int32
	SetBody(v *DeleteContext0ResponseBody) *DeleteContext0Response
	GetBody() *DeleteContext0ResponseBody
}

type DeleteContext0Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContext0ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContext0Response) String() string {
	return dara.Prettify(s)
}

func (s DeleteContext0Response) GoString() string {
	return s.String()
}

func (s *DeleteContext0Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContext0Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContext0Response) GetBody() *DeleteContext0ResponseBody {
	return s.Body
}

func (s *DeleteContext0Response) SetHeaders(v map[string]*string) *DeleteContext0Response {
	s.Headers = v
	return s
}

func (s *DeleteContext0Response) SetStatusCode(v int32) *DeleteContext0Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteContext0Response) SetBody(v *DeleteContext0ResponseBody) *DeleteContext0Response {
	s.Body = v
	return s
}

func (s *DeleteContext0Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
