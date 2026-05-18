// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsQuotaListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarFsQuotaListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarFsQuotaListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarFsQuotaListResponseBody) *DescribePolarFsQuotaListResponse
	GetBody() *DescribePolarFsQuotaListResponseBody
}

type DescribePolarFsQuotaListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarFsQuotaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarFsQuotaListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaListResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarFsQuotaListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarFsQuotaListResponse) GetBody() *DescribePolarFsQuotaListResponseBody {
	return s.Body
}

func (s *DescribePolarFsQuotaListResponse) SetHeaders(v map[string]*string) *DescribePolarFsQuotaListResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarFsQuotaListResponse) SetStatusCode(v int32) *DescribePolarFsQuotaListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarFsQuotaListResponse) SetBody(v *DescribePolarFsQuotaListResponseBody) *DescribePolarFsQuotaListResponse {
	s.Body = v
	return s
}

func (s *DescribePolarFsQuotaListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
