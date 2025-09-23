// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResourcePackageResponseBody) *ModifyResourcePackageResponse
	GetBody() *ModifyResourcePackageResponseBody
}

type ModifyResourcePackageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResourcePackageResponse) GetBody() *ModifyResourcePackageResponseBody {
	return s.Body
}

func (s *ModifyResourcePackageResponse) SetHeaders(v map[string]*string) *ModifyResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourcePackageResponse) SetStatusCode(v int32) *ModifyResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourcePackageResponse) SetBody(v *ModifyResourcePackageResponseBody) *ModifyResourcePackageResponse {
	s.Body = v
	return s
}

func (s *ModifyResourcePackageResponse) Validate() error {
	return dara.Validate(s)
}
