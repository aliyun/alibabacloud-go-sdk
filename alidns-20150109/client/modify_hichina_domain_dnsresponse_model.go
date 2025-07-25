// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHichinaDomainDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHichinaDomainDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHichinaDomainDNSResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHichinaDomainDNSResponseBody) *ModifyHichinaDomainDNSResponse
	GetBody() *ModifyHichinaDomainDNSResponseBody
}

type ModifyHichinaDomainDNSResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHichinaDomainDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHichinaDomainDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHichinaDomainDNSResponse) GoString() string {
	return s.String()
}

func (s *ModifyHichinaDomainDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHichinaDomainDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHichinaDomainDNSResponse) GetBody() *ModifyHichinaDomainDNSResponseBody {
	return s.Body
}

func (s *ModifyHichinaDomainDNSResponse) SetHeaders(v map[string]*string) *ModifyHichinaDomainDNSResponse {
	s.Headers = v
	return s
}

func (s *ModifyHichinaDomainDNSResponse) SetStatusCode(v int32) *ModifyHichinaDomainDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHichinaDomainDNSResponse) SetBody(v *ModifyHichinaDomainDNSResponseBody) *ModifyHichinaDomainDNSResponse {
	s.Body = v
	return s
}

func (s *ModifyHichinaDomainDNSResponse) Validate() error {
	return dara.Validate(s)
}
