// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourcePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourcePriceResponse
	GetStatusCode() *int32
	SetBody(v *GetResourcePriceResponseBody) *GetResourcePriceResponse
	GetBody() *GetResourcePriceResponseBody
}

type GetResourcePriceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourcePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourcePriceResponse) GetBody() *GetResourcePriceResponseBody {
	return s.Body
}

func (s *GetResourcePriceResponse) SetHeaders(v map[string]*string) *GetResourcePriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePriceResponse) SetStatusCode(v int32) *GetResourcePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePriceResponse) SetBody(v *GetResourcePriceResponseBody) *GetResourcePriceResponse {
	s.Body = v
	return s
}

func (s *GetResourcePriceResponse) Validate() error {
	return dara.Validate(s)
}
