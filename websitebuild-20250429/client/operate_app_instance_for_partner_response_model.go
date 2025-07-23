// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppInstanceForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateAppInstanceForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateAppInstanceForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *OperateAppInstanceForPartnerResponseBody) *OperateAppInstanceForPartnerResponse
	GetBody() *OperateAppInstanceForPartnerResponseBody
}

type OperateAppInstanceForPartnerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAppInstanceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAppInstanceForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateAppInstanceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *OperateAppInstanceForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateAppInstanceForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateAppInstanceForPartnerResponse) GetBody() *OperateAppInstanceForPartnerResponseBody {
	return s.Body
}

func (s *OperateAppInstanceForPartnerResponse) SetHeaders(v map[string]*string) *OperateAppInstanceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *OperateAppInstanceForPartnerResponse) SetStatusCode(v int32) *OperateAppInstanceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponse) SetBody(v *OperateAppInstanceForPartnerResponseBody) *OperateAppInstanceForPartnerResponse {
	s.Body = v
	return s
}

func (s *OperateAppInstanceForPartnerResponse) Validate() error {
	return dara.Validate(s)
}
