// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceModifyPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceModifyPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceModifyPriceResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceModifyPriceResponseBody) *GetResourceModifyPriceResponse
	GetBody() *GetResourceModifyPriceResponseBody
}

type GetResourceModifyPriceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceModifyPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceModifyPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceModifyPriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceModifyPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceModifyPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceModifyPriceResponse) GetBody() *GetResourceModifyPriceResponseBody {
	return s.Body
}

func (s *GetResourceModifyPriceResponse) SetHeaders(v map[string]*string) *GetResourceModifyPriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceModifyPriceResponse) SetStatusCode(v int32) *GetResourceModifyPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceModifyPriceResponse) SetBody(v *GetResourceModifyPriceResponseBody) *GetResourceModifyPriceResponse {
	s.Body = v
	return s
}

func (s *GetResourceModifyPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
