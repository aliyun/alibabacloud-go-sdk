// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldGroupSaleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopOversoldGroupSaleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopOversoldGroupSaleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopOversoldGroupSaleResponseBody) *ModifyDesktopOversoldGroupSaleResponse
	GetBody() *ModifyDesktopOversoldGroupSaleResponseBody
}

type ModifyDesktopOversoldGroupSaleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopOversoldGroupSaleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopOversoldGroupSaleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldGroupSaleResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldGroupSaleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopOversoldGroupSaleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopOversoldGroupSaleResponse) GetBody() *ModifyDesktopOversoldGroupSaleResponseBody {
	return s.Body
}

func (s *ModifyDesktopOversoldGroupSaleResponse) SetHeaders(v map[string]*string) *ModifyDesktopOversoldGroupSaleResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleResponse) SetStatusCode(v int32) *ModifyDesktopOversoldGroupSaleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleResponse) SetBody(v *ModifyDesktopOversoldGroupSaleResponseBody) *ModifyDesktopOversoldGroupSaleResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleResponse) Validate() error {
	return dara.Validate(s)
}
