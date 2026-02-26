// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSelectionSkuSaleInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSelectionSkuSaleInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSelectionSkuSaleInfosResponse
	GetStatusCode() *int32
	SetBody(v *SkuSaleInfoListResult) *ListSelectionSkuSaleInfosResponse
	GetBody() *SkuSaleInfoListResult
}

type ListSelectionSkuSaleInfosResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkuSaleInfoListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionSkuSaleInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSelectionSkuSaleInfosResponse) GoString() string {
	return s.String()
}

func (s *ListSelectionSkuSaleInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSelectionSkuSaleInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSelectionSkuSaleInfosResponse) GetBody() *SkuSaleInfoListResult {
	return s.Body
}

func (s *ListSelectionSkuSaleInfosResponse) SetHeaders(v map[string]*string) *ListSelectionSkuSaleInfosResponse {
	s.Headers = v
	return s
}

func (s *ListSelectionSkuSaleInfosResponse) SetStatusCode(v int32) *ListSelectionSkuSaleInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSelectionSkuSaleInfosResponse) SetBody(v *SkuSaleInfoListResult) *ListSelectionSkuSaleInfosResponse {
	s.Body = v
	return s
}

func (s *ListSelectionSkuSaleInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
