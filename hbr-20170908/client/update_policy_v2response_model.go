// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolicyV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolicyV2Response
	GetStatusCode() *int32
	SetBody(v *UpdatePolicyV2ResponseBody) *UpdatePolicyV2Response
	GetBody() *UpdatePolicyV2ResponseBody
}

type UpdatePolicyV2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolicyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolicyV2Response) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2Response) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolicyV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolicyV2Response) GetBody() *UpdatePolicyV2ResponseBody {
	return s.Body
}

func (s *UpdatePolicyV2Response) SetHeaders(v map[string]*string) *UpdatePolicyV2Response {
	s.Headers = v
	return s
}

func (s *UpdatePolicyV2Response) SetStatusCode(v int32) *UpdatePolicyV2Response {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolicyV2Response) SetBody(v *UpdatePolicyV2ResponseBody) *UpdatePolicyV2Response {
	s.Body = v
	return s
}

func (s *UpdatePolicyV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
