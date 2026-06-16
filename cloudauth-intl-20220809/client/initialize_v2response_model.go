// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeV2Response
	GetStatusCode() *int32
	SetBody(v *InitializeV2ResponseBody) *InitializeV2Response
	GetBody() *InitializeV2ResponseBody
}

type InitializeV2Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeV2Response) String() string {
	return dara.Prettify(s)
}

func (s InitializeV2Response) GoString() string {
	return s.String()
}

func (s *InitializeV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeV2Response) GetBody() *InitializeV2ResponseBody {
	return s.Body
}

func (s *InitializeV2Response) SetHeaders(v map[string]*string) *InitializeV2Response {
	s.Headers = v
	return s
}

func (s *InitializeV2Response) SetStatusCode(v int32) *InitializeV2Response {
	s.StatusCode = &v
	return s
}

func (s *InitializeV2Response) SetBody(v *InitializeV2ResponseBody) *InitializeV2Response {
	s.Body = v
	return s
}

func (s *InitializeV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
