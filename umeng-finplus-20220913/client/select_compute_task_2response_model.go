// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectComputeTask2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectComputeTask2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectComputeTask2Response
	GetStatusCode() *int32
	SetBody(v *SelectComputeTask2ResponseBody) *SelectComputeTask2Response
	GetBody() *SelectComputeTask2ResponseBody
}

type SelectComputeTask2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectComputeTask2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectComputeTask2Response) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTask2Response) GoString() string {
	return s.String()
}

func (s *SelectComputeTask2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectComputeTask2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectComputeTask2Response) GetBody() *SelectComputeTask2ResponseBody {
	return s.Body
}

func (s *SelectComputeTask2Response) SetHeaders(v map[string]*string) *SelectComputeTask2Response {
	s.Headers = v
	return s
}

func (s *SelectComputeTask2Response) SetStatusCode(v int32) *SelectComputeTask2Response {
	s.StatusCode = &v
	return s
}

func (s *SelectComputeTask2Response) SetBody(v *SelectComputeTask2ResponseBody) *SelectComputeTask2Response {
	s.Body = v
	return s
}

func (s *SelectComputeTask2Response) Validate() error {
	return dara.Validate(s)
}
