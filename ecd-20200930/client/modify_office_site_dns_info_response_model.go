// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteDnsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOfficeSiteDnsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOfficeSiteDnsInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOfficeSiteDnsInfoResponseBody) *ModifyOfficeSiteDnsInfoResponse
	GetBody() *ModifyOfficeSiteDnsInfoResponseBody
}

type ModifyOfficeSiteDnsInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOfficeSiteDnsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOfficeSiteDnsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteDnsInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteDnsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOfficeSiteDnsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOfficeSiteDnsInfoResponse) GetBody() *ModifyOfficeSiteDnsInfoResponseBody {
	return s.Body
}

func (s *ModifyOfficeSiteDnsInfoResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteDnsInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteDnsInfoResponse) SetStatusCode(v int32) *ModifyOfficeSiteDnsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOfficeSiteDnsInfoResponse) SetBody(v *ModifyOfficeSiteDnsInfoResponseBody) *ModifyOfficeSiteDnsInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyOfficeSiteDnsInfoResponse) Validate() error {
	return dara.Validate(s)
}
