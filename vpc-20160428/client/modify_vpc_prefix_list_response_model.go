// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcPrefixListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcPrefixListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcPrefixListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcPrefixListResponseBody) *ModifyVpcPrefixListResponse
	GetBody() *ModifyVpcPrefixListResponseBody
}

type ModifyVpcPrefixListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcPrefixListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcPrefixListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcPrefixListResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcPrefixListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcPrefixListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcPrefixListResponse) GetBody() *ModifyVpcPrefixListResponseBody {
	return s.Body
}

func (s *ModifyVpcPrefixListResponse) SetHeaders(v map[string]*string) *ModifyVpcPrefixListResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcPrefixListResponse) SetStatusCode(v int32) *ModifyVpcPrefixListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcPrefixListResponse) SetBody(v *ModifyVpcPrefixListResponseBody) *ModifyVpcPrefixListResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcPrefixListResponse) Validate() error {
	return dara.Validate(s)
}
