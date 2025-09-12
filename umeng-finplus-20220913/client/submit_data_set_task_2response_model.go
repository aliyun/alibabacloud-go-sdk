// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataSetTask2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDataSetTask2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDataSetTask2Response
	GetStatusCode() *int32
	SetBody(v *SubmitDataSetTask2ResponseBody) *SubmitDataSetTask2Response
	GetBody() *SubmitDataSetTask2ResponseBody
}

type SubmitDataSetTask2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDataSetTask2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDataSetTask2Response) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataSetTask2Response) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTask2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDataSetTask2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDataSetTask2Response) GetBody() *SubmitDataSetTask2ResponseBody {
	return s.Body
}

func (s *SubmitDataSetTask2Response) SetHeaders(v map[string]*string) *SubmitDataSetTask2Response {
	s.Headers = v
	return s
}

func (s *SubmitDataSetTask2Response) SetStatusCode(v int32) *SubmitDataSetTask2Response {
	s.StatusCode = &v
	return s
}

func (s *SubmitDataSetTask2Response) SetBody(v *SubmitDataSetTask2ResponseBody) *SubmitDataSetTask2Response {
	s.Body = v
	return s
}

func (s *SubmitDataSetTask2Response) Validate() error {
	return dara.Validate(s)
}
