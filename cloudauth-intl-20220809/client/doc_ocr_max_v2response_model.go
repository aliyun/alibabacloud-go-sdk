// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrMaxV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocOcrMaxV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocOcrMaxV2Response
	GetStatusCode() *int32
	SetBody(v *DocOcrMaxV2ResponseBody) *DocOcrMaxV2Response
	GetBody() *DocOcrMaxV2ResponseBody
}

type DocOcrMaxV2Response struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocOcrMaxV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocOcrMaxV2Response) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxV2Response) GoString() string {
	return s.String()
}

func (s *DocOcrMaxV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocOcrMaxV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocOcrMaxV2Response) GetBody() *DocOcrMaxV2ResponseBody {
	return s.Body
}

func (s *DocOcrMaxV2Response) SetHeaders(v map[string]*string) *DocOcrMaxV2Response {
	s.Headers = v
	return s
}

func (s *DocOcrMaxV2Response) SetStatusCode(v int32) *DocOcrMaxV2Response {
	s.StatusCode = &v
	return s
}

func (s *DocOcrMaxV2Response) SetBody(v *DocOcrMaxV2ResponseBody) *DocOcrMaxV2Response {
	s.Body = v
	return s
}

func (s *DocOcrMaxV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
