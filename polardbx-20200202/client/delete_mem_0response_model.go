// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMem0Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMem0Response
	GetStatusCode() *int32
	SetBody(v *DeleteMem0ResponseBody) *DeleteMem0Response
	GetBody() *DeleteMem0ResponseBody
}

type DeleteMem0Response struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMem0ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMem0Response) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0Response) GoString() string {
	return s.String()
}

func (s *DeleteMem0Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMem0Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMem0Response) GetBody() *DeleteMem0ResponseBody {
	return s.Body
}

func (s *DeleteMem0Response) SetHeaders(v map[string]*string) *DeleteMem0Response {
	s.Headers = v
	return s
}

func (s *DeleteMem0Response) SetStatusCode(v int32) *DeleteMem0Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteMem0Response) SetBody(v *DeleteMem0ResponseBody) *DeleteMem0Response {
	s.Body = v
	return s
}

func (s *DeleteMem0Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
