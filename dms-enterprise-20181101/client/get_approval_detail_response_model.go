// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApprovalDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApprovalDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetApprovalDetailResponseBody) *GetApprovalDetailResponse
	GetBody() *GetApprovalDetailResponseBody
}

type GetApprovalDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApprovalDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApprovalDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponse) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApprovalDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApprovalDetailResponse) GetBody() *GetApprovalDetailResponseBody {
	return s.Body
}

func (s *GetApprovalDetailResponse) SetHeaders(v map[string]*string) *GetApprovalDetailResponse {
	s.Headers = v
	return s
}

func (s *GetApprovalDetailResponse) SetStatusCode(v int32) *GetApprovalDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApprovalDetailResponse) SetBody(v *GetApprovalDetailResponseBody) *GetApprovalDetailResponse {
	s.Body = v
	return s
}

func (s *GetApprovalDetailResponse) Validate() error {
	return dara.Validate(s)
}
