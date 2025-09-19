// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckSaleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckSaleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckSaleResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckSaleResponseBody) *GetCheckSaleResponse
	GetBody() *GetCheckSaleResponseBody
}

type GetCheckSaleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckSaleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckSaleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSaleResponse) GoString() string {
	return s.String()
}

func (s *GetCheckSaleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckSaleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckSaleResponse) GetBody() *GetCheckSaleResponseBody {
	return s.Body
}

func (s *GetCheckSaleResponse) SetHeaders(v map[string]*string) *GetCheckSaleResponse {
	s.Headers = v
	return s
}

func (s *GetCheckSaleResponse) SetStatusCode(v int32) *GetCheckSaleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckSaleResponse) SetBody(v *GetCheckSaleResponseBody) *GetCheckSaleResponse {
	s.Body = v
	return s
}

func (s *GetCheckSaleResponse) Validate() error {
	return dara.Validate(s)
}
