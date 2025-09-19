// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProtectVpcListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddProtectVpcListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddProtectVpcListResponse
	GetStatusCode() *int32
	SetBody(v *AddProtectVpcListResponseBody) *AddProtectVpcListResponse
	GetBody() *AddProtectVpcListResponseBody
}

type AddProtectVpcListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddProtectVpcListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddProtectVpcListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddProtectVpcListResponse) GoString() string {
	return s.String()
}

func (s *AddProtectVpcListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddProtectVpcListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddProtectVpcListResponse) GetBody() *AddProtectVpcListResponseBody {
	return s.Body
}

func (s *AddProtectVpcListResponse) SetHeaders(v map[string]*string) *AddProtectVpcListResponse {
	s.Headers = v
	return s
}

func (s *AddProtectVpcListResponse) SetStatusCode(v int32) *AddProtectVpcListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProtectVpcListResponse) SetBody(v *AddProtectVpcListResponseBody) *AddProtectVpcListResponse {
	s.Body = v
	return s
}

func (s *AddProtectVpcListResponse) Validate() error {
	return dara.Validate(s)
}
