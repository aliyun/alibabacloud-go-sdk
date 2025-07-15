// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWritingV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunWritingV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunWritingV2Response
	GetStatusCode() *int32
	SetBody(v *RunWritingV2ResponseBody) *RunWritingV2Response
	GetBody() *RunWritingV2ResponseBody
}

type RunWritingV2Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunWritingV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunWritingV2Response) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2Response) GoString() string {
	return s.String()
}

func (s *RunWritingV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunWritingV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunWritingV2Response) GetBody() *RunWritingV2ResponseBody {
	return s.Body
}

func (s *RunWritingV2Response) SetHeaders(v map[string]*string) *RunWritingV2Response {
	s.Headers = v
	return s
}

func (s *RunWritingV2Response) SetStatusCode(v int32) *RunWritingV2Response {
	s.StatusCode = &v
	return s
}

func (s *RunWritingV2Response) SetBody(v *RunWritingV2ResponseBody) *RunWritingV2Response {
	s.Body = v
	return s
}

func (s *RunWritingV2Response) Validate() error {
	return dara.Validate(s)
}
