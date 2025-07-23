// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppServiceForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateAppServiceForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateAppServiceForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *OperateAppServiceForPartnerResponseBody) *OperateAppServiceForPartnerResponse
	GetBody() *OperateAppServiceForPartnerResponseBody
}

type OperateAppServiceForPartnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAppServiceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAppServiceForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateAppServiceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *OperateAppServiceForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateAppServiceForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateAppServiceForPartnerResponse) GetBody() *OperateAppServiceForPartnerResponseBody {
	return s.Body
}

func (s *OperateAppServiceForPartnerResponse) SetHeaders(v map[string]*string) *OperateAppServiceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *OperateAppServiceForPartnerResponse) SetStatusCode(v int32) *OperateAppServiceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAppServiceForPartnerResponse) SetBody(v *OperateAppServiceForPartnerResponseBody) *OperateAppServiceForPartnerResponse {
	s.Body = v
	return s
}

func (s *OperateAppServiceForPartnerResponse) Validate() error {
	return dara.Validate(s)
}
