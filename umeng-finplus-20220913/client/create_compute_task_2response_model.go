// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeTask2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComputeTask2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComputeTask2Response
	GetStatusCode() *int32
	SetBody(v *CreateComputeTask2ResponseBody) *CreateComputeTask2Response
	GetBody() *CreateComputeTask2ResponseBody
}

type CreateComputeTask2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeTask2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeTask2Response) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeTask2Response) GoString() string {
	return s.String()
}

func (s *CreateComputeTask2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComputeTask2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComputeTask2Response) GetBody() *CreateComputeTask2ResponseBody {
	return s.Body
}

func (s *CreateComputeTask2Response) SetHeaders(v map[string]*string) *CreateComputeTask2Response {
	s.Headers = v
	return s
}

func (s *CreateComputeTask2Response) SetStatusCode(v int32) *CreateComputeTask2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeTask2Response) SetBody(v *CreateComputeTask2ResponseBody) *CreateComputeTask2Response {
	s.Body = v
	return s
}

func (s *CreateComputeTask2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
