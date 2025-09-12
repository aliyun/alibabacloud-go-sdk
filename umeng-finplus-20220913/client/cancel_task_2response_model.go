// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTask2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelTask2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelTask2Response
	GetStatusCode() *int32
	SetBody(v *CancelTask2ResponseBody) *CancelTask2Response
	GetBody() *CancelTask2ResponseBody
}

type CancelTask2Response struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelTask2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelTask2Response) String() string {
	return dara.Prettify(s)
}

func (s CancelTask2Response) GoString() string {
	return s.String()
}

func (s *CancelTask2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelTask2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelTask2Response) GetBody() *CancelTask2ResponseBody {
	return s.Body
}

func (s *CancelTask2Response) SetHeaders(v map[string]*string) *CancelTask2Response {
	s.Headers = v
	return s
}

func (s *CancelTask2Response) SetStatusCode(v int32) *CancelTask2Response {
	s.StatusCode = &v
	return s
}

func (s *CancelTask2Response) SetBody(v *CancelTask2ResponseBody) *CancelTask2Response {
	s.Body = v
	return s
}

func (s *CancelTask2Response) Validate() error {
	return dara.Validate(s)
}
