// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateFinanceTaxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateFinanceTaxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateFinanceTaxResponse
	GetStatusCode() *int32
	SetBody(v *OperateFinanceTaxResponseBody) *OperateFinanceTaxResponse
	GetBody() *OperateFinanceTaxResponseBody
}

type OperateFinanceTaxResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateFinanceTaxResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateFinanceTaxResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateFinanceTaxResponse) GoString() string {
	return s.String()
}

func (s *OperateFinanceTaxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateFinanceTaxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateFinanceTaxResponse) GetBody() *OperateFinanceTaxResponseBody {
	return s.Body
}

func (s *OperateFinanceTaxResponse) SetHeaders(v map[string]*string) *OperateFinanceTaxResponse {
	s.Headers = v
	return s
}

func (s *OperateFinanceTaxResponse) SetStatusCode(v int32) *OperateFinanceTaxResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateFinanceTaxResponse) SetBody(v *OperateFinanceTaxResponseBody) *OperateFinanceTaxResponse {
	s.Body = v
	return s
}

func (s *OperateFinanceTaxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
