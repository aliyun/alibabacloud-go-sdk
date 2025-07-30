// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModifyBEClusterInquiryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModifyBEClusterInquiryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModifyBEClusterInquiryResponse
	GetStatusCode() *int32
	SetBody(v *GetModifyBEClusterInquiryResponseBody) *GetModifyBEClusterInquiryResponse
	GetBody() *GetModifyBEClusterInquiryResponseBody
}

type GetModifyBEClusterInquiryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModifyBEClusterInquiryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModifyBEClusterInquiryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponse) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModifyBEClusterInquiryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModifyBEClusterInquiryResponse) GetBody() *GetModifyBEClusterInquiryResponseBody {
	return s.Body
}

func (s *GetModifyBEClusterInquiryResponse) SetHeaders(v map[string]*string) *GetModifyBEClusterInquiryResponse {
	s.Headers = v
	return s
}

func (s *GetModifyBEClusterInquiryResponse) SetStatusCode(v int32) *GetModifyBEClusterInquiryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponse) SetBody(v *GetModifyBEClusterInquiryResponseBody) *GetModifyBEClusterInquiryResponse {
	s.Body = v
	return s
}

func (s *GetModifyBEClusterInquiryResponse) Validate() error {
	return dara.Validate(s)
}
