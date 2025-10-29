// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarFsQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarFsQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarFsQuotaResponseBody) *DescribePolarFsQuotaResponse
	GetBody() *DescribePolarFsQuotaResponseBody
}

type DescribePolarFsQuotaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarFsQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarFsQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarFsQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarFsQuotaResponse) GetBody() *DescribePolarFsQuotaResponseBody {
	return s.Body
}

func (s *DescribePolarFsQuotaResponse) SetHeaders(v map[string]*string) *DescribePolarFsQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarFsQuotaResponse) SetStatusCode(v int32) *DescribePolarFsQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarFsQuotaResponse) SetBody(v *DescribePolarFsQuotaResponseBody) *DescribePolarFsQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribePolarFsQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
