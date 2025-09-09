// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDrdsIpWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDrdsIpWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDrdsIpWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDrdsIpWhiteListResponseBody) *ModifyDrdsIpWhiteListResponse
	GetBody() *ModifyDrdsIpWhiteListResponseBody
}

type ModifyDrdsIpWhiteListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDrdsIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDrdsIpWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDrdsIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDrdsIpWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDrdsIpWhiteListResponse) GetBody() *ModifyDrdsIpWhiteListResponseBody {
	return s.Body
}

func (s *ModifyDrdsIpWhiteListResponse) SetHeaders(v map[string]*string) *ModifyDrdsIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyDrdsIpWhiteListResponse) SetStatusCode(v int32) *ModifyDrdsIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDrdsIpWhiteListResponse) SetBody(v *ModifyDrdsIpWhiteListResponseBody) *ModifyDrdsIpWhiteListResponse {
	s.Body = v
	return s
}

func (s *ModifyDrdsIpWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
