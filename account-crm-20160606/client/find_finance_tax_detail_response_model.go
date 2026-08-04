// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindFinanceTaxDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindFinanceTaxDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindFinanceTaxDetailResponse
	GetStatusCode() *int32
	SetBody(v *FindFinanceTaxDetailResponseBody) *FindFinanceTaxDetailResponse
	GetBody() *FindFinanceTaxDetailResponseBody
}

type FindFinanceTaxDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindFinanceTaxDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindFinanceTaxDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s FindFinanceTaxDetailResponse) GoString() string {
	return s.String()
}

func (s *FindFinanceTaxDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindFinanceTaxDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindFinanceTaxDetailResponse) GetBody() *FindFinanceTaxDetailResponseBody {
	return s.Body
}

func (s *FindFinanceTaxDetailResponse) SetHeaders(v map[string]*string) *FindFinanceTaxDetailResponse {
	s.Headers = v
	return s
}

func (s *FindFinanceTaxDetailResponse) SetStatusCode(v int32) *FindFinanceTaxDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *FindFinanceTaxDetailResponse) SetBody(v *FindFinanceTaxDetailResponseBody) *FindFinanceTaxDetailResponse {
	s.Body = v
	return s
}

func (s *FindFinanceTaxDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
