// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeTask2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeTask2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeTask2Response
	GetStatusCode() *int32
	SetBody(v *ListComputeTask2ResponseBody) *ListComputeTask2Response
	GetBody() *ListComputeTask2ResponseBody
}

type ListComputeTask2Response struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeTask2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeTask2Response) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTask2Response) GoString() string {
	return s.String()
}

func (s *ListComputeTask2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeTask2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeTask2Response) GetBody() *ListComputeTask2ResponseBody {
	return s.Body
}

func (s *ListComputeTask2Response) SetHeaders(v map[string]*string) *ListComputeTask2Response {
	s.Headers = v
	return s
}

func (s *ListComputeTask2Response) SetStatusCode(v int32) *ListComputeTask2Response {
	s.StatusCode = &v
	return s
}

func (s *ListComputeTask2Response) SetBody(v *ListComputeTask2ResponseBody) *ListComputeTask2Response {
	s.Body = v
	return s
}

func (s *ListComputeTask2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
