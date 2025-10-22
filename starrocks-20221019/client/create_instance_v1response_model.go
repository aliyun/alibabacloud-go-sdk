// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceV1Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceV1Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceV1Response
	GetStatusCode() *int32
	SetBody(v *CreateInstanceV1ResponseBody) *CreateInstanceV1Response
	GetBody() *CreateInstanceV1ResponseBody
}

type CreateInstanceV1Response struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceV1ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceV1Response) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1Response) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceV1Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceV1Response) GetBody() *CreateInstanceV1ResponseBody {
	return s.Body
}

func (s *CreateInstanceV1Response) SetHeaders(v map[string]*string) *CreateInstanceV1Response {
	s.Headers = v
	return s
}

func (s *CreateInstanceV1Response) SetStatusCode(v int32) *CreateInstanceV1Response {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceV1Response) SetBody(v *CreateInstanceV1ResponseBody) *CreateInstanceV1Response {
	s.Body = v
	return s
}

func (s *CreateInstanceV1Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
