// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskContentV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskContentV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskContentV2Response
	GetStatusCode() *int32
	SetBody(v *UpdateTaskContentV2ResponseBody) *UpdateTaskContentV2Response
	GetBody() *UpdateTaskContentV2ResponseBody
}

type UpdateTaskContentV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskContentV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskContentV2Response) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskContentV2Response) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskContentV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskContentV2Response) GetBody() *UpdateTaskContentV2ResponseBody {
	return s.Body
}

func (s *UpdateTaskContentV2Response) SetHeaders(v map[string]*string) *UpdateTaskContentV2Response {
	s.Headers = v
	return s
}

func (s *UpdateTaskContentV2Response) SetStatusCode(v int32) *UpdateTaskContentV2Response {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskContentV2Response) SetBody(v *UpdateTaskContentV2ResponseBody) *UpdateTaskContentV2Response {
	s.Body = v
	return s
}

func (s *UpdateTaskContentV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
