// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBakDataSourceStorageAccessInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBakDataSourceStorageAccessInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBakDataSourceStorageAccessInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBakDataSourceStorageAccessInfoResponseBody) *DescribeBakDataSourceStorageAccessInfoResponse
	GetBody() *DescribeBakDataSourceStorageAccessInfoResponseBody
}

type DescribeBakDataSourceStorageAccessInfoResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBakDataSourceStorageAccessInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBakDataSourceStorageAccessInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBakDataSourceStorageAccessInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeBakDataSourceStorageAccessInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBakDataSourceStorageAccessInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBakDataSourceStorageAccessInfoResponse) GetBody() *DescribeBakDataSourceStorageAccessInfoResponseBody {
	return s.Body
}

func (s *DescribeBakDataSourceStorageAccessInfoResponse) SetHeaders(v map[string]*string) *DescribeBakDataSourceStorageAccessInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponse) SetStatusCode(v int32) *DescribeBakDataSourceStorageAccessInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponse) SetBody(v *DescribeBakDataSourceStorageAccessInfoResponseBody) *DescribeBakDataSourceStorageAccessInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
