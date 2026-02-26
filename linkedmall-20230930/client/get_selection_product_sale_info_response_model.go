// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSelectionProductSaleInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSelectionProductSaleInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSelectionProductSaleInfoResponse
	GetStatusCode() *int32
	SetBody(v *ProductSaleInfo) *GetSelectionProductSaleInfoResponse
	GetBody() *ProductSaleInfo
}

type GetSelectionProductSaleInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductSaleInfo   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSelectionProductSaleInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSelectionProductSaleInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSelectionProductSaleInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSelectionProductSaleInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSelectionProductSaleInfoResponse) GetBody() *ProductSaleInfo {
	return s.Body
}

func (s *GetSelectionProductSaleInfoResponse) SetHeaders(v map[string]*string) *GetSelectionProductSaleInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSelectionProductSaleInfoResponse) SetStatusCode(v int32) *GetSelectionProductSaleInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSelectionProductSaleInfoResponse) SetBody(v *ProductSaleInfo) *GetSelectionProductSaleInfoResponse {
	s.Body = v
	return s
}

func (s *GetSelectionProductSaleInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
