// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkitemV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkitemV2Response
	GetStatusCode() *int32
	SetBody(v *CreateWorkitemV2ResponseBody) *CreateWorkitemV2Response
	GetBody() *CreateWorkitemV2ResponseBody
}

type CreateWorkitemV2Response struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkitemV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkitemV2Response) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemV2Response) GoString() string {
	return s.String()
}

func (s *CreateWorkitemV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkitemV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkitemV2Response) GetBody() *CreateWorkitemV2ResponseBody {
	return s.Body
}

func (s *CreateWorkitemV2Response) SetHeaders(v map[string]*string) *CreateWorkitemV2Response {
	s.Headers = v
	return s
}

func (s *CreateWorkitemV2Response) SetStatusCode(v int32) *CreateWorkitemV2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemV2Response) SetBody(v *CreateWorkitemV2ResponseBody) *CreateWorkitemV2Response {
	s.Body = v
	return s
}

func (s *CreateWorkitemV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
