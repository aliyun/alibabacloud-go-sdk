// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOwnerApplyOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOwnerApplyOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOwnerApplyOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetOwnerApplyOrderDetailResponseBody) *GetOwnerApplyOrderDetailResponse
	GetBody() *GetOwnerApplyOrderDetailResponseBody
}

type GetOwnerApplyOrderDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOwnerApplyOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOwnerApplyOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOwnerApplyOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOwnerApplyOrderDetailResponse) GetBody() *GetOwnerApplyOrderDetailResponseBody {
	return s.Body
}

func (s *GetOwnerApplyOrderDetailResponse) SetHeaders(v map[string]*string) *GetOwnerApplyOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponse) SetStatusCode(v int32) *GetOwnerApplyOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponse) SetBody(v *GetOwnerApplyOrderDetailResponseBody) *GetOwnerApplyOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponse) Validate() error {
	return dara.Validate(s)
}
