// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPublicUrlIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPublicUrlIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPublicUrlIpListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPublicUrlIpListResponseBody) *ModifyPublicUrlIpListResponse
	GetBody() *ModifyPublicUrlIpListResponseBody
}

type ModifyPublicUrlIpListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPublicUrlIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPublicUrlIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPublicUrlIpListResponse) GoString() string {
	return s.String()
}

func (s *ModifyPublicUrlIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPublicUrlIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPublicUrlIpListResponse) GetBody() *ModifyPublicUrlIpListResponseBody {
	return s.Body
}

func (s *ModifyPublicUrlIpListResponse) SetHeaders(v map[string]*string) *ModifyPublicUrlIpListResponse {
	s.Headers = v
	return s
}

func (s *ModifyPublicUrlIpListResponse) SetStatusCode(v int32) *ModifyPublicUrlIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPublicUrlIpListResponse) SetBody(v *ModifyPublicUrlIpListResponseBody) *ModifyPublicUrlIpListResponse {
	s.Body = v
	return s
}

func (s *ModifyPublicUrlIpListResponse) Validate() error {
	return dara.Validate(s)
}
