// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcClassicLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableVpcClassicLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableVpcClassicLinkResponse
	GetStatusCode() *int32
	SetBody(v *DisableVpcClassicLinkResponseBody) *DisableVpcClassicLinkResponse
	GetBody() *DisableVpcClassicLinkResponseBody
}

type DisableVpcClassicLinkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableVpcClassicLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableVpcClassicLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcClassicLinkResponse) GoString() string {
	return s.String()
}

func (s *DisableVpcClassicLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableVpcClassicLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableVpcClassicLinkResponse) GetBody() *DisableVpcClassicLinkResponseBody {
	return s.Body
}

func (s *DisableVpcClassicLinkResponse) SetHeaders(v map[string]*string) *DisableVpcClassicLinkResponse {
	s.Headers = v
	return s
}

func (s *DisableVpcClassicLinkResponse) SetStatusCode(v int32) *DisableVpcClassicLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableVpcClassicLinkResponse) SetBody(v *DisableVpcClassicLinkResponseBody) *DisableVpcClassicLinkResponse {
	s.Body = v
	return s
}

func (s *DisableVpcClassicLinkResponse) Validate() error {
	return dara.Validate(s)
}
