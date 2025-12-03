// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityGroupsResponseBody) *ModifySecurityGroupsResponse
	GetBody() *ModifySecurityGroupsResponseBody
}

type ModifySecurityGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityGroupsResponse) GetBody() *ModifySecurityGroupsResponseBody {
	return s.Body
}

func (s *ModifySecurityGroupsResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupsResponse) SetStatusCode(v int32) *ModifySecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityGroupsResponse) SetBody(v *ModifySecurityGroupsResponseBody) *ModifySecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
