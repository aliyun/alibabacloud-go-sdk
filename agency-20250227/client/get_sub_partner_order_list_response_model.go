// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubPartnerOrderListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubPartnerOrderListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubPartnerOrderListResponse
	GetStatusCode() *int32
	SetBody(v *GetSubPartnerOrderListResponseBody) *GetSubPartnerOrderListResponse
	GetBody() *GetSubPartnerOrderListResponseBody
}

type GetSubPartnerOrderListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubPartnerOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubPartnerOrderListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerOrderListResponse) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubPartnerOrderListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubPartnerOrderListResponse) GetBody() *GetSubPartnerOrderListResponseBody {
	return s.Body
}

func (s *GetSubPartnerOrderListResponse) SetHeaders(v map[string]*string) *GetSubPartnerOrderListResponse {
	s.Headers = v
	return s
}

func (s *GetSubPartnerOrderListResponse) SetStatusCode(v int32) *GetSubPartnerOrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubPartnerOrderListResponse) SetBody(v *GetSubPartnerOrderListResponseBody) *GetSubPartnerOrderListResponse {
	s.Body = v
	return s
}

func (s *GetSubPartnerOrderListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
