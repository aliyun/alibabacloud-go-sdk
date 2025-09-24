// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPayAsYouGoPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPayAsYouGoPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPayAsYouGoPriceResponse
	GetStatusCode() *int32
	SetBody(v *GetPayAsYouGoPriceResponseBody) *GetPayAsYouGoPriceResponse
	GetBody() *GetPayAsYouGoPriceResponseBody
}

type GetPayAsYouGoPriceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPayAsYouGoPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPayAsYouGoPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceResponse) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPayAsYouGoPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPayAsYouGoPriceResponse) GetBody() *GetPayAsYouGoPriceResponseBody {
	return s.Body
}

func (s *GetPayAsYouGoPriceResponse) SetHeaders(v map[string]*string) *GetPayAsYouGoPriceResponse {
	s.Headers = v
	return s
}

func (s *GetPayAsYouGoPriceResponse) SetStatusCode(v int32) *GetPayAsYouGoPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPayAsYouGoPriceResponse) SetBody(v *GetPayAsYouGoPriceResponseBody) *GetPayAsYouGoPriceResponse {
	s.Body = v
	return s
}

func (s *GetPayAsYouGoPriceResponse) Validate() error {
	return dara.Validate(s)
}
