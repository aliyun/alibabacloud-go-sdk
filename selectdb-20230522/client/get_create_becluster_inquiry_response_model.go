// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateBEClusterInquiryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCreateBEClusterInquiryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCreateBEClusterInquiryResponse
	GetStatusCode() *int32
	SetBody(v *GetCreateBEClusterInquiryResponseBody) *GetCreateBEClusterInquiryResponse
	GetBody() *GetCreateBEClusterInquiryResponseBody
}

type GetCreateBEClusterInquiryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateBEClusterInquiryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreateBEClusterInquiryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponse) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCreateBEClusterInquiryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCreateBEClusterInquiryResponse) GetBody() *GetCreateBEClusterInquiryResponseBody {
	return s.Body
}

func (s *GetCreateBEClusterInquiryResponse) SetHeaders(v map[string]*string) *GetCreateBEClusterInquiryResponse {
	s.Headers = v
	return s
}

func (s *GetCreateBEClusterInquiryResponse) SetStatusCode(v int32) *GetCreateBEClusterInquiryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponse) SetBody(v *GetCreateBEClusterInquiryResponseBody) *GetCreateBEClusterInquiryResponse {
	s.Body = v
	return s
}

func (s *GetCreateBEClusterInquiryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
