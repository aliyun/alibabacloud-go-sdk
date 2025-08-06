// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeSlsRoleV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssumeSlsRoleV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssumeSlsRoleV2Response
	GetStatusCode() *int32
	SetBody(v *AssumeSlsRoleV2ResponseBody) *AssumeSlsRoleV2Response
	GetBody() *AssumeSlsRoleV2ResponseBody
}

type AssumeSlsRoleV2Response struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeSlsRoleV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeSlsRoleV2Response) String() string {
	return dara.Prettify(s)
}

func (s AssumeSlsRoleV2Response) GoString() string {
	return s.String()
}

func (s *AssumeSlsRoleV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssumeSlsRoleV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssumeSlsRoleV2Response) GetBody() *AssumeSlsRoleV2ResponseBody {
	return s.Body
}

func (s *AssumeSlsRoleV2Response) SetHeaders(v map[string]*string) *AssumeSlsRoleV2Response {
	s.Headers = v
	return s
}

func (s *AssumeSlsRoleV2Response) SetStatusCode(v int32) *AssumeSlsRoleV2Response {
	s.StatusCode = &v
	return s
}

func (s *AssumeSlsRoleV2Response) SetBody(v *AssumeSlsRoleV2ResponseBody) *AssumeSlsRoleV2Response {
	s.Body = v
	return s
}

func (s *AssumeSlsRoleV2Response) Validate() error {
	return dara.Validate(s)
}
