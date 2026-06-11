// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenewalRateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRenewalRateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRenewalRateListResponse
	GetStatusCode() *int32
	SetBody(v *GetRenewalRateListResponseBody) *GetRenewalRateListResponse
	GetBody() *GetRenewalRateListResponseBody
}

type GetRenewalRateListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRenewalRateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRenewalRateListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRenewalRateListResponse) GoString() string {
	return s.String()
}

func (s *GetRenewalRateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRenewalRateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRenewalRateListResponse) GetBody() *GetRenewalRateListResponseBody {
	return s.Body
}

func (s *GetRenewalRateListResponse) SetHeaders(v map[string]*string) *GetRenewalRateListResponse {
	s.Headers = v
	return s
}

func (s *GetRenewalRateListResponse) SetStatusCode(v int32) *GetRenewalRateListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRenewalRateListResponse) SetBody(v *GetRenewalRateListResponseBody) *GetRenewalRateListResponse {
	s.Body = v
	return s
}

func (s *GetRenewalRateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
