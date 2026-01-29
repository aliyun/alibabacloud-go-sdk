// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePackagePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourcePackagePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourcePackagePriceResponse
	GetStatusCode() *int32
	SetBody(v *GetResourcePackagePriceResponseBody) *GetResourcePackagePriceResponse
	GetBody() *GetResourcePackagePriceResponseBody
}

type GetResourcePackagePriceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcePackagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcePackagePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePackagePriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourcePackagePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourcePackagePriceResponse) GetBody() *GetResourcePackagePriceResponseBody {
	return s.Body
}

func (s *GetResourcePackagePriceResponse) SetHeaders(v map[string]*string) *GetResourcePackagePriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePackagePriceResponse) SetStatusCode(v int32) *GetResourcePackagePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePackagePriceResponse) SetBody(v *GetResourcePackagePriceResponseBody) *GetResourcePackagePriceResponse {
	s.Body = v
	return s
}

func (s *GetResourcePackagePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
