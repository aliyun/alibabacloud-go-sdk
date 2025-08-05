// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolicyV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolicyV2Response
	GetStatusCode() *int32
	SetBody(v *CreatePolicyV2ResponseBody) *CreatePolicyV2Response
	GetBody() *CreatePolicyV2ResponseBody
}

type CreatePolicyV2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyV2Response) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2Response) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolicyV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolicyV2Response) GetBody() *CreatePolicyV2ResponseBody {
	return s.Body
}

func (s *CreatePolicyV2Response) SetHeaders(v map[string]*string) *CreatePolicyV2Response {
	s.Headers = v
	return s
}

func (s *CreatePolicyV2Response) SetStatusCode(v int32) *CreatePolicyV2Response {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyV2Response) SetBody(v *CreatePolicyV2ResponseBody) *CreatePolicyV2Response {
	s.Body = v
	return s
}

func (s *CreatePolicyV2Response) Validate() error {
	return dara.Validate(s)
}
