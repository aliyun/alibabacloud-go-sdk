// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocOcrV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocOcrV2Response
	GetStatusCode() *int32
	SetBody(v *DocOcrV2ResponseBody) *DocOcrV2Response
	GetBody() *DocOcrV2ResponseBody
}

type DocOcrV2Response struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocOcrV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocOcrV2Response) String() string {
	return dara.Prettify(s)
}

func (s DocOcrV2Response) GoString() string {
	return s.String()
}

func (s *DocOcrV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocOcrV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocOcrV2Response) GetBody() *DocOcrV2ResponseBody {
	return s.Body
}

func (s *DocOcrV2Response) SetHeaders(v map[string]*string) *DocOcrV2Response {
	s.Headers = v
	return s
}

func (s *DocOcrV2Response) SetStatusCode(v int32) *DocOcrV2Response {
	s.StatusCode = &v
	return s
}

func (s *DocOcrV2Response) SetBody(v *DocOcrV2ResponseBody) *DocOcrV2Response {
	s.Body = v
	return s
}

func (s *DocOcrV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
