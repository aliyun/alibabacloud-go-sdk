// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansDiscountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySavingsPlansDiscountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySavingsPlansDiscountResponse
	GetStatusCode() *int32
	SetBody(v *QuerySavingsPlansDiscountResponseBody) *QuerySavingsPlansDiscountResponse
	GetBody() *QuerySavingsPlansDiscountResponseBody
}

type QuerySavingsPlansDiscountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySavingsPlansDiscountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySavingsPlansDiscountResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDiscountResponse) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDiscountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySavingsPlansDiscountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySavingsPlansDiscountResponse) GetBody() *QuerySavingsPlansDiscountResponseBody {
	return s.Body
}

func (s *QuerySavingsPlansDiscountResponse) SetHeaders(v map[string]*string) *QuerySavingsPlansDiscountResponse {
	s.Headers = v
	return s
}

func (s *QuerySavingsPlansDiscountResponse) SetStatusCode(v int32) *QuerySavingsPlansDiscountResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponse) SetBody(v *QuerySavingsPlansDiscountResponseBody) *QuerySavingsPlansDiscountResponse {
	s.Body = v
	return s
}

func (s *QuerySavingsPlansDiscountResponse) Validate() error {
	return dara.Validate(s)
}
